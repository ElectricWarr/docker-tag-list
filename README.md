# tag-list

Docker image (or command, your choice) to query Docker Hub/Registry for information that should _really_ be provided via the Docker CLI :(

Created this repo after realising that no-one sees things on Dockerhub.

# Usage

Example:

```
$ my-machine:tag-list michael$ tag-list parity/parity
TAG             ARCHITECTURE  OS
nightly         amd64         linux
stable-release  amd64         linux
stable          amd64         linux
beta-release    amd64         linux
beta            amd64         linux
v1.9.4          amd64         linux
v1.8.11         amd64         linux
v1.9.3          amd64         linux
v1.8.10         amd64         linux
v1.8.9          amd64         linux
```

Ideally there would be the _option_ of filtering and of displaying more columns, but I wrote this to be a quick reference when setting up projects with [dockerprep](https://github.com/ElectricWarr/dockerprep)

# Setup

1. Copy `project.cfg.template` to `project.cfg` (don't worry, it's in your `.gitignore`)

If you'd like to push to your own Docker Hub Repo:

2. Edit `project.cfg` to include your dockerhub ID/namespace

# Development

- `./build.sh` to build the image
- `./debug.sh` to drop into a shell in a running instance of the container
- `./run.sh REPO_NAME` to run the tool from your local build

# ToDo

## Fix OS Spam

Sometimes output looks like this:

```
TAG          ARCHITECTURE  OS
latest       amd64         amd64    s390x  ppc64le  386  arm64  arm    arm    amd64  windows  windows  linux  linux  linux  linux  linux  linux  linux
3            amd64         amd64    s390x  ppc64le  386  arm64  arm    arm    amd64  windows  windows  linux  linux  linux  linux  linux  linux  linux
3.6          amd64         amd64    s390x  ppc64le  386  arm64  arm    arm    amd64  windows  windows  linux  linux  linux  linux  linux  linux  linux
3.6.4        amd64         amd64    s390x  ppc64le  386  arm64  arm    arm    amd64  windows  windows  linux  linux  linux  linux  linux  linux  linux
rc           amd64         amd64    s390x  ppc64le  386  arm64  arm    arm    amd64  windows  windows  linux  linux  linux  linux  linux  linux  linux
3.7-rc       amd64         amd64    s390x  ppc64le  386  arm64  arm    arm    amd64  windows  windows  linux  linux  linux  linux  linux  linux  linux
3.7.0b2      amd64         amd64    s390x  ppc64le  386  arm64  arm    arm    amd64  windows  windows  linux  linux  linux  linux  linux  linux  linux
onbuild      s390x         ppc64le  386    arm64    arm  arm    amd64  linux  linux  linux    linux    linux  linux  linux
3-onbuild    s390x         ppc64le  386    arm64    arm  arm    amd64  linux  linux  linux    linux    linux  linux  linux
3.6-onbuild  s390x         ppc64le  386    arm64    arm  arm    amd64  linux  linux  linux    linux    linux  linux  linux
```

No idea why!

# Ideas

## `:latest`

Which tag matches `latest`? (Which tags are the same?)

Requires:

- Getting and comparing manifest data from dockerhub

## Base Image

What is the base image of each tag?

Requirements:

- No idea

