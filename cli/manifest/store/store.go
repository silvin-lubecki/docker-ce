package store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/docker/cli/cli/manifest/types"
	"github.com/docker/distribution/reference"
)

// Store manages local storage of image distribution manifests
type Store interface {
	Remove(listRef reference.Reference) error
	Get(listRef reference.Reference, manifest reference.Reference) (types.ImageManifest, error)
	GetList(listRef reference.Reference) ([]types.ImageManifest, error)
	Save(listRef reference.Reference, manifest reference.Reference, image types.ImageManifest) error
}

// fsStore manages manifest files stored on the local filesystem
type fsStore struct {
	root string
}

// NewStore returns a new store for a local file path
func NewStore(root string) Store {
	return &fsStore{root: root}
}

// Remove a manifest list from local storage
func (s *fsStore) Remove(listRef reference.Reference) error {
	path := filepath.Join(s.root, makeFilesafeName(listRef.String()))
	return os.RemoveAll(path)
}

// Get returns the local manifest
func (s *fsStore) Get(listRef reference.Reference, manifest reference.Reference) (types.ImageManifest, error) {
	filename := manifestToFilename(s.root, listRef.String(), manifest.String())
	return s.getFromFilename(manifest, filename)
}

func (s *fsStore) getFromFilename(ref reference.Reference, filename string) (types.ImageManifest, error) {
	bytes, err := ioutil.ReadFile(filename)
	switch {
	case os.IsNotExist(err):
		return types.ImageManifest{}, newNotFoundError(ref.String())
	case err != nil:
		return types.ImageManifest{}, err
	}
	var manifestInfo types.ImageManifest
	return manifestInfo, json.Unmarshal(bytes, &manifestInfo)
}

// GetList returns all the local manifests for a transaction
func (s *fsStore) GetList(listRef reference.Reference) ([]types.ImageManifest, error) {
	filenames, err := s.listManifests(listRef.String())
	switch {
	case err != nil:
		return nil, err
	case filenames == nil:
		return nil, newNotFoundError(listRef.String())
	}

	manifests := []types.ImageManifest{}
	for _, filename := range filenames {
		filename = filepath.Join(s.root, makeFilesafeName(listRef.String()), filename)
		manifest, err := s.getFromFilename(listRef, filename)
		if err != nil {
			return nil, err
		}
		manifests = append(manifests, manifest)
	}
	return manifests, nil
}

// listManifests stored in a transaction
func (s *fsStore) listManifests(transaction string) ([]string, error) {
	transactionDir := filepath.Join(s.root, makeFilesafeName(transaction))
	fileInfos, err := ioutil.ReadDir(transactionDir)
	switch {
	case os.IsNotExist(err):
		return nil, nil
	case err != nil:
		return nil, err
	}

	filenames := []string{}
	for _, info := range fileInfos {
		filenames = append(filenames, info.Name())
	}
	return filenames, nil
}

// Save a manifest as part of a local manifest list
func (s *fsStore) Save(listRef reference.Reference, manifest reference.Reference, image types.ImageManifest) error {
	if err := s.createManifestListDirectory(listRef.String()); err != nil {
		return err
	}
	filename := manifestToFilename(s.root, listRef.String(), manifest.String())
	bytes, err := json.Marshal(image)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, bytes, 0644)
}

func (s *fsStore) createManifestListDirectory(transaction string) error {
	path := filepath.Join(s.root, makeFilesafeName(transaction))
	return os.MkdirAll(path, 0755)
}

func manifestToFilename(root, manifestList, manifest string) string {
	return filepath.Join(root, makeFilesafeName(manifestList), makeFilesafeName(manifest))
}

func makeFilesafeName(ref string) string {
	fileName := strings.Replace(ref, ":", "-", -1)
	return strings.Replace(fileName, "/", "_", -1)
}

type notFoundError struct {
	object string
}

func newNotFoundError(ref string) *notFoundError {
	return &notFoundError{object: ref}
}

func (n *notFoundError) Error() string {
	return fmt.Sprintf("No such manifest: %s", n.object)
}

// NotFound interface
func (n *notFoundError) NotFound() {}

// IsNotFound returns true if the error is a not found error
func IsNotFound(err error) bool {
	_, ok := err.(notFound)
	return ok
}

type notFound interface {
	NotFound()
}
