package trust

import (
	"fmt"
	"path"
	"sort"
	"strings"

	"github.com/docker/cli/cli"
	"github.com/docker/cli/cli/command"
	"github.com/docker/cli/cli/command/image"
	"github.com/docker/cli/cli/trust"
	"github.com/docker/notary/client"
	"github.com/docker/notary/tuf/data"
	"github.com/spf13/cobra"
)

func newSignCommand(dockerCli command.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sign [OPTIONS] IMAGE:TAG",
		Short: "Sign an image",
		Args:  cli.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return signImage(dockerCli, args[0])
		},
	}
	return cmd
}

func signImage(cli command.Cli, imageName string) error {
	ctx, ref, repoInfo, authConfig, err := getImageReferencesAndAuth(cli, imageName)
	if err != nil {
		return err
	}

	notaryRepo, err := trust.GetNotaryRepository(cli, repoInfo, *authConfig, "push", "pull")
	if err != nil {
		return trust.NotaryError(ref.Name(), err)
	}
	if err = clearChangeList(notaryRepo); err != nil {
		return err
	}
	defer clearChangeList(notaryRepo)
	tag, err := getTag(ref)
	if err != nil {
		return err
	}
	if tag == "" {
		return fmt.Errorf("No tag specified for %s", imageName)
	}

	// get the latest repository metadata so we can figure out which roles to sign
	if err = notaryRepo.Update(false); err != nil {
		switch err.(type) {
		case client.ErrRepoNotInitialized, client.ErrRepositoryNotExist:
			// before initializing a new repo, check that the image exists locally:
			if err := checkLocalImageExistence(ctx, cli, imageName); err != nil {
				return err
			}

			userRole := data.RoleName(path.Join(data.CanonicalTargetsRole.String(), authConfig.Username))
			if err := initNotaryRepoWithSigners(notaryRepo, userRole); err != nil {
				return trust.NotaryError(ref.Name(), err)
			}

			fmt.Fprintf(cli.Out(), "Created signer: %s\n", authConfig.Username)
			fmt.Fprintf(cli.Out(), "Finished initializing %q\n", notaryRepo.GetGUN().String())
		default:
			return trust.NotaryError(repoInfo.Name.Name(), err)
		}
	}
	requestPrivilege := command.RegistryAuthenticationPrivilegedFunc(cli, repoInfo.Index, "push")
	target, err := createTarget(notaryRepo, tag)
	if err != nil {
		switch err := err.(type) {
		case client.ErrNoSuchTarget, client.ErrRepositoryNotExist:
			// Fail fast if the image doesn't exist locally
			if err := checkLocalImageExistence(ctx, cli, imageName); err != nil {
				return err
			}
			return image.TrustedPush(ctx, cli, repoInfo, ref, *authConfig, requestPrivilege)
		default:
			return err
		}
	}

	fmt.Fprintf(cli.Out(), "Signing and pushing trust metadata for %s\n", imageName)
	existingSigInfo, err := getExistingSignatureInfoForReleasedTag(notaryRepo, tag)
	if err != nil {
		return err
	}
	err = image.AddTargetToAllSignableRoles(notaryRepo, &target)
	if err == nil {
		prettyPrintExistingSignatureInfo(cli, existingSigInfo)
		err = notaryRepo.Publish()
	}
	if err != nil {
		return fmt.Errorf("failed to sign %q:%s - %s", repoInfo.Name.Name(), tag, err.Error())
	}
	fmt.Fprintf(cli.Out(), "Successfully signed %q:%s\n", repoInfo.Name.Name(), tag)
	return nil
}

func createTarget(notaryRepo *client.NotaryRepository, tag string) (client.Target, error) {
	target := &client.Target{}
	var err error
	if tag == "" {
		return *target, fmt.Errorf("No tag specified")
	}
	target.Name = tag
	target.Hashes, target.Length, err = getSignedManifestHashAndSize(notaryRepo, tag)
	return *target, err
}

func getSignedManifestHashAndSize(notaryRepo *client.NotaryRepository, tag string) (data.Hashes, int64, error) {
	targets, err := notaryRepo.GetAllTargetMetadataByName(tag)
	if err != nil {
		return nil, 0, err
	}
	return getReleasedTargetHashAndSize(targets, tag)
}

func getReleasedTargetHashAndSize(targets []client.TargetSignedStruct, tag string) (data.Hashes, int64, error) {
	for _, tgt := range targets {
		if isReleasedTarget(tgt.Role.Name) {
			return tgt.Target.Hashes, tgt.Target.Length, nil
		}
	}
	return nil, 0, client.ErrNoSuchTarget(tag)
}

func getExistingSignatureInfoForReleasedTag(notaryRepo *client.NotaryRepository, tag string) (trustTagRow, error) {
	targets, err := notaryRepo.GetAllTargetMetadataByName(tag)
	if err != nil {
		return trustTagRow{}, err
	}
	releasedTargetInfoList := matchReleasedSignatures(targets)
	if len(releasedTargetInfoList) == 0 {
		return trustTagRow{}, nil
	}
	return releasedTargetInfoList[0], nil
}

func prettyPrintExistingSignatureInfo(cli command.Cli, existingSigInfo trustTagRow) {
	sort.Strings(existingSigInfo.Signers)
	joinedSigners := strings.Join(existingSigInfo.Signers, ", ")
	fmt.Fprintf(cli.Out(), "Existing signatures for tag %s digest %s from:\n%s\n", existingSigInfo.TagName, existingSigInfo.HashHex, joinedSigners)
}

func initNotaryRepoWithSigners(notaryRepo *client.NotaryRepository, newSigner data.RoleName) error {
	rootKey, err := getOrGenerateNotaryKey(notaryRepo, data.CanonicalRootRole)
	if err != nil {
		return err
	}
	rootKeyID := rootKey.ID()

	// Initialize the notary repository with a remotely managed snapshot key
	if err := notaryRepo.Initialize([]string{rootKeyID}, data.CanonicalSnapshotRole); err != nil {
		return err
	}

	signerKey, err := getOrGenerateNotaryKey(notaryRepo, newSigner)
	if err != nil {
		return err
	}
	addStagedSigner(notaryRepo, newSigner, []data.PublicKey{signerKey})

	return notaryRepo.Publish()
}

// generates an ECDSA key without a GUN for the specified role
func getOrGenerateNotaryKey(notaryRepo *client.NotaryRepository, role data.RoleName) (data.PublicKey, error) {
	// use the signer name in the PEM headers if this is a delegation key
	if data.IsDelegation(role) {
		role = data.RoleName(notaryRoleToSigner(role))
	}
	keys := notaryRepo.CryptoService.ListKeys(role)
	var err error
	var key data.PublicKey
	// always select the first key by ID
	if len(keys) > 0 {
		sort.Strings(keys)
		keyID := keys[0]
		privKey, _, err := notaryRepo.CryptoService.GetPrivateKey(keyID)
		if err != nil {
			return nil, err
		}
		key = data.PublicKeyFromPrivate(privKey)
	} else {
		key, err = notaryRepo.CryptoService.Create(role, "", data.ECDSAKey)
		if err != nil {
			return nil, err
		}
	}
	return key, nil
}

// stages changes to add a signer with the specified name and key(s).  Adds to targets/<name> and targets/releases
func addStagedSigner(notaryRepo *client.NotaryRepository, newSigner data.RoleName, signerKeys []data.PublicKey) {
	// create targets/<username>
	notaryRepo.AddDelegationRoleAndKeys(newSigner, signerKeys)
	notaryRepo.AddDelegationPaths(newSigner, []string{""})

	// create targets/releases
	notaryRepo.AddDelegationRoleAndKeys(trust.ReleasesRole, signerKeys)
	notaryRepo.AddDelegationPaths(trust.ReleasesRole, []string{""})
}
