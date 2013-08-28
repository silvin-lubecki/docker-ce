## A maintainer's guide to releasing Docker

So you're in charge of a Docker release? Cool. Here's what to do.

If your experience deviates from this document, please document the changes
to keep it up-to-date.


### 1. Pull from master and create a release branch

```bash
git checkout master
git pull
git checkout -b bump_$VERSION
```

### 2. Update CHANGELOG.md

You can run this command for reference:

```bash
LAST_VERSION=$(git tag | grep -E "v[0-9\.]+$" | sort -nr | head -n 1)
git log $LAST_VERSION..HEAD
```

Each change should be formatted as ```BULLET CATEGORY: DESCRIPTION```

* BULLET is either ```-```, ```+``` or ```*```, to indicate a bugfix,
  new feature or upgrade, respectively.

* CATEGORY should describe which part of the project is affected.
  Valid categories are:
  * Builder
  * Documentation
  * Hack
  * Packaging
  * Remote API
  * Runtime

* DESCRIPTION: a concise description of the change that is relevant to the 
  end-user, using the present tense. Changes should be described in terms 
  of how they affect the user, for example "new feature X which allows Y", 
  "fixed bug which caused X", "increased performance of Y".

EXAMPLES:

```
+ Builder: 'docker build -t FOO' applies the tag FOO to the newly built
  container.
* Runtime: improve detection of kernel version
- Remote API: fix a bug in the optional unix socket transport
```

### 3. Change the contents of the VERSION file

### 4. Run all tests

```bash
go test
```

### 5. Commit and create a pull request

```bash
git add CHANGELOG.md
git commit -m "Bump version to $VERSION"
git push origin bump_$VERSION
```

### 6. Get 2 other maintainers to validate the pull request

### 7. Merge the pull request and apply tags

```bash
git checkout master
git merge bump_$VERSION
git tag -a v$VERSION # Don't forget the v!
git tag -f -a latest
git push
git push --tags
```

### 8. Publish binaries

To run this you will need access to the release credentials.
Get them from [the infrastructure maintainers](
https://github.com/dotcloud/docker/blob/master/hack/infrastructure/MAINTAINERS).

```bash
docker build -t releasedocker .
docker run  \
	-e AWS_S3_BUCKET=get-nightly.docker.io \
	-e AWS_ACCESS_KEY=$(cat ~/.aws/access_key) \
	-e AWS_SECRET_KEY=$(cat ~/.aws/secret_key) \
	-e GPG_PASSPHRASE=supersecretsesame \
	releasedocker
```

It will build and upload the binaries on the specified bucket (you should
use get-nightly.docker.io for general testing, and once everything is fine,
switch to get.docker.io).


### 9. Rejoice!

Congratulations! You're done.
