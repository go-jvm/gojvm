# Contribution Guide

### Go

GO is written in [Go](http://golang.org).
If you don't have a Go development environment,
please [set one up](http://golang.org/doc/code.html).

The version of GO should be **1.12** or above.

After installation, you'll need `GOPATH` defined,
and `PATH` modified to access your Go binaries.

A common setup is the following but you could always google a setup for your own flavor.

```sh
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```



## Workflow

### Step 1: Fork in the cloud

1. Visit https://github.com/go-jvm/gojvm
2. Click `Fork` button (top right) to establish a cloud-based fork.

### Step 2: Clone fork to local storage


Define a local working directory:

```sh
# If your GOPATH has multiple paths, pick
# just one and use it instead of $GOPATH here.
working_dir=$GOPATH/src/github.com/gojvm
```
