#!/bin/bash
set -e

# usage: ./generate.sh [versions]
#    ie: ./generate.sh
#        to update all Dockerfiles in this directory
#    or: ./generate.sh debian-jessie
#        to only update debian-jessie/Dockerfile
#    or: ./generate.sh debian-newversion
#        to create a new folder and a Dockerfile within it

cd "$(dirname "$(readlink -f "$BASH_SOURCE")")"

versions=( "$@" )
if [ ${#versions[@]} -eq 0 ]; then
	versions=( */ )
fi
versions=( "${versions[@]%/}" )

for version in "${versions[@]}"; do
	distro="${version%-*}"
	suite="${version##*-}"
	from="${distro}:${suite}"

	mkdir -p "$version"
	echo "$version -> FROM $from"
	cat > "$version/Dockerfile" <<-EOF
		#
		# THIS FILE IS AUTOGENERATED; SEE "contrib/builder/deb/generate.sh"!
		#

		FROM $from
	EOF

	case "$from" in
		debian:wheezy)
			# add -backports, like our users have to
			echo "RUN echo deb http://http.debian.net/debian $suite-backports main > /etc/apt/sources.list.d/$suite-backports.list" >> "$version/Dockerfile"
			;;
	esac

	echo >> "$version/Dockerfile"

	extraBuildTags=

	# this list is sorted alphabetically; please keep it that way
	packages=(
		bash-completion # for bash-completion debhelper integration
		btrfs-tools # for "btrfs/ioctl.h" (and "version.h" if possible)
		build-essential # "essential for building Debian packages"
		curl ca-certificates # for downloading Go
		debhelper # for easy ".deb" building
		dh-apparmor # for apparmor debhelper
		dh-systemd # for systemd debhelper integration
		git # for "git commit" info in "docker -v"
		libapparmor-dev # for "sys/apparmor.h"
		libdevmapper-dev # for "libdevmapper.h"
		libsqlite3-dev # for "sqlite3.h"
	)

	if [ "$suite" = 'precise' ]; then
		# precise has a few package issues

		# - dh-systemd doesn't exist at all
		packages=( "${packages[@]/dh-systemd}" )

		# - libdevmapper-dev is missing critical structs (too old)
		packages=( "${packages[@]/libdevmapper-dev}" )
		extraBuildTags+=' exclude_graphdriver_devicemapper'

		# - btrfs-tools is missing "ioctl.h" (too old), so it's useless
		#   (since kernels on precise are old too, just skip btrfs entirely)
		packages=( "${packages[@]/btrfs-tools}" )
		extraBuildTags+=' exclude_graphdriver_btrfs'
	fi

	echo "RUN apt-get update && apt-get install -y ${packages[*]} --no-install-recommends && rm -rf /var/lib/apt/lists/*" >> "$version/Dockerfile"

	echo >> "$version/Dockerfile"

	awk '$1 == "ENV" && $2 == "GO_VERSION" { print; exit }' ../../../Dockerfile >> "$version/Dockerfile"
	echo 'RUN curl -fSL "https://storage.googleapis.com/golang/go${GO_VERSION}.linux-amd64.tar.gz" | tar xzC /usr/local' >> "$version/Dockerfile"
	echo 'ENV PATH $PATH:/usr/local/go/bin' >> "$version/Dockerfile"

	echo >> "$version/Dockerfile"

	echo 'ENV AUTO_GOPATH 1' >> "$version/Dockerfile"
	awk '$1 == "ENV" && $2 == "DOCKER_BUILDTAGS" { print $0 "'"$extraBuildTags"'"; exit }' ../../../Dockerfile >> "$version/Dockerfile"
done
