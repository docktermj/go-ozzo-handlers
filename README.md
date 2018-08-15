# go-ozzo-handlers

## Warning

This is exemplar code only.  It does not build.

Build `go-ozzo-handlers-M.m.P-I.x86_64.rpm`
and   `go-ozzo-handlers_M.m.P-I_amd64.deb`
where "M.m.P-I" is Major.minor.Patch-Iteration.

## Usage

A simple "hello world" program.
The purpose of the repository is to show how to:

1. Build go executable locally
1. Build go executable via Docker
1. Build RPM / DEB installation via Docker.

### Invocation

```console
go-ozzo-handlers
```

## Development

### Dependencies

#### Set environment variables

```console
export GOPATH="${HOME}/go"
export PATH="${PATH}:${GOPATH}/bin:/usr/local/go/bin"
export PROJECT_DIR="${GOPATH}/src/github.com/docktermj"
export REPOSITORY_DIR="${PROJECT_DIR}/go-ozzo-handlers"
```

#### Download project

```console
mkdir -p ${PROJECT_DIR}
cd ${PROJECT_DIR}
git clone git@github.com:docktermj/go-ozzo-handlers.git
```

#### Download dependencies

```console
cd ${REPOSITORY_DIR}
make dependencies
```

### Build

#### Local build

```console
cd ${REPOSITORY_DIR}
make
```

The results will be in the `${GOPATH}/bin` directory.

#### Docker build

```console
cd ${REPOSITORY_DIR}
make build
```

The results will be in the `.../target` directory.

### Test

```console
cd ${REPOSITORY_DIR}
make test-local
```

### Install

#### RPM-based

Example distributions: openSUSE, Fedora, CentOS, Mandrake

##### RPM Install

Example:

```console
sudo rpm -ivh go-ozzo-handlers-M.m.P-I.x86_64.rpm
```

##### RPM Update

Example: 

```console
sudo rpm -Uvh go-ozzo-handlers-M.m.P-I.x86_64.rpm
```

#### Debian

Example distributions: Ubuntu

##### Debian Install / Update

Example:

```console
sudo dpkg -i go-ozzo-handlers_M.m.P-I_amd64.deb
```

### Cleanup

```console
cd ${REPOSITORY_DIR}
make clean
```
