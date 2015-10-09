<!--[metadata]>
+++
title = "Apply custom metadata"
description = "Learn how to work with custom metadata in Docker, using labels."
keywords = ["Usage, user guide, labels, metadata, docker, documentation, examples,  annotating"]
[menu.main]
parent = "mn_use_docker"
+++
<![end-metadata]-->

# Apply custom metadata

You can apply metadata to your images, containers, or daemons via
labels. Labels serve a wide range of uses, such as adding notes or licensing
information to an image, or to identify a host.

A label is a `<key>` / `<value>` pair. Docker stores the label values as
*strings*. You can specify multiple labels but each `<key>` must be
unique or the value will be overwritten. If you specify the same `key` several
times but with different values, newer labels overwrite previous labels. Docker
uses the last `key=value` you supply.

>**Note:** Support for daemon-labels was added in Docker 1.4.1. Labels on
>containers and images are new in Docker 1.6.0

## Label keys (namespaces)

Docker puts no hard restrictions on the `key` used for a label. However, using
simple keys can easily lead to conflicts. For example, you have chosen to
categorize your images by CPU architecture using "architecture" labels in
your Dockerfiles:

    LABEL architecture="amd64"

    LABEL architecture="ARMv7"

Another user may apply the same label based on a building's "architecture":

    LABEL architecture="Art Nouveau"

To prevent naming conflicts, Docker recommends using namespaces to label keys
using reverse domain notation. Use the following guidelines to name your keys:

- All (third-party) tools should prefix their keys with the
  reverse DNS notation of a domain controlled by the author. For
  example, `com.example.some-label`.

- The `com.docker.*`, `io.docker.*` and `org.dockerproject.*` namespaces are
  reserved for Docker's internal use.

- Keys should only consist of lower-cased alphanumeric characters,
  dots and dashes (for example, `[a-z0-9-.]`)

- Keys should start *and* end with an alpha numeric character

- Keys may not contain consecutive dots or dashes.

- Keys *without* namespace (dots) are reserved for CLI use. This allows end-
  users to add metadata to their containers and images without having to type
  cumbersome namespaces on the command-line.


These are simply guidelines and Docker does not *enforce* them. However, for
the benefit of the community, you *should* use namespaces for your label keys.


## Store structured data in labels

Label values can contain any data type as long as it can be represented as a
string. For example, consider this JSON document:


    {
        "Description": "A containerized foobar",
        "Usage": "docker run --rm example/foobar [args]",
        "License": "GPL",
        "Version": "0.0.1-beta",
        "aBoolean": true,
        "aNumber" : 0.01234,
        "aNestedArray": ["a", "b", "c"]
    }

You can store this struct in a label by serializing it to a string first:

    LABEL com.example.image-specs="{\"Description\":\"A containerized foobar\",\"Usage\":\"docker run --rm example\\/foobar [args]\",\"License\":\"GPL\",\"Version\":\"0.0.1-beta\",\"aBoolean\":true,\"aNumber\":0.01234,\"aNestedArray\":[\"a\",\"b\",\"c\"]}"

While it is *possible* to store structured data in label values, Docker treats
this data as a 'regular' string. This means that Docker doesn't offer ways to
query (filter) based on nested properties. If your tool needs to filter on
nested properties, the tool itself needs to implement this functionality.


## Add labels to images

To add labels to an image, use the `LABEL` instruction in your Dockerfile:


    LABEL [<namespace>.]<key>[=<value>] ...

The `LABEL` instruction adds a label to your image, optionally with a value.
Use surrounding quotes or backslashes for labels that contain
white space characters in the `<value>`:

    LABEL vendor=ACME\ Incorporated
    LABEL com.example.version.is-beta
    LABEL com.example.version="0.0.1-beta"
    LABEL com.example.release-date="2015-02-12"

The `LABEL` instruction also supports setting multiple `<key>` / `<value>` pairs
in a single instruction:

    LABEL com.example.version="0.0.1-beta" com.example.release-date="2015-02-12"

Long lines can be split up by using a backslash (`\`) as continuation marker:

    LABEL vendor=ACME\ Incorporated \
          com.example.is-beta \
          com.example.version="0.0.1-beta" \
          com.example.release-date="2015-02-12"

Docker recommends you add multiple labels in a single `LABEL` instruction. Using
individual instructions for each label can result in an inefficient image. This
is because each `LABEL` instruction in a Dockerfile produces a new IMAGE layer.

You can view the labels via the `docker inspect` command:

    $ docker inspect 4fa6e0f0c678

    ...
    "Labels": {
        "vendor": "ACME Incorporated",
        "com.example.is-beta": "",
        "com.example.version": "0.0.1-beta",
        "com.example.release-date": "2015-02-12"
    }
    ...

    # Inspect labels on container
    $ docker inspect -f "{{json .Config.Labels }}" 4fa6e0f0c678

    {"Vendor":"ACME Incorporated","com.example.is-beta":"","com.example.version":"0.0.1-beta","com.example.release-date":"2015-02-12"}

    # Inspect labels on images
    $ docker inspect -f "{{json .ContainerConfig.Labels }}" myimage


## Query labels

Besides storing metadata, you can filter images and containers by label. To list all
running containers that have the `com.example.is-beta` label:

    # List all running containers that have a `com.example.is-beta` label
    $ docker ps --filter "label=com.example.is-beta"

List all running containers with the label `color` that have a value `blue`:

    $ docker ps --filter "label=color=blue"

List all images with the label `vendor` that have the value `ACME`:

    $ docker images --filter "label=vendor=ACME"


## Daemon labels


    docker daemon \
      --dns 8.8.8.8 \
      --dns 8.8.4.4 \
      -H unix:///var/run/docker.sock \
      --label com.example.environment="production" \
      --label com.example.storage="ssd"

These labels appear as part of the `docker info` output for the daemon:

    docker -D info
    Containers: 12
    Images: 672
    Storage Driver: aufs
     Root Dir: /var/lib/docker/aufs
     Backing Filesystem: extfs
     Dirs: 697
    Execution Driver: native-0.2
    Logging Driver: json-file
    Kernel Version: 3.13.0-32-generic
    Operating System: Ubuntu 14.04.1 LTS
    CPUs: 1
    Total Memory: 994.1 MiB
    Name: docker.example.com
    ID: RC3P:JTCT:32YS:XYSB:YUBG:VFED:AAJZ:W3YW:76XO:D7NN:TEVU:UCRW
    Debug mode (server): false
    Debug mode (client): true
    File Descriptors: 11
    Goroutines: 14
    EventsListeners: 0
    Init Path: /usr/bin/docker
    Docker Root Dir: /var/lib/docker
    WARNING: No swap limit support
    Labels:
     com.example.environment=production
     com.example.storage=ssd
