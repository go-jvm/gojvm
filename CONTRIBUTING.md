# Contribution Guide

### Go

GO is written in [Go](http://golang.org).
If you don't have a Go development environment,
please [set one up](http://golang.org/doc/code.html).

The version of GO should be **1.11** or above.

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

Create your clone:

```sh
mkdir -p $working_dir
cd $working_dir
git clone https://github.com/$user/gojvm.git
# the following is recommended
# or: git clone git@github.com:$user/gojvm.git

cd $working_dir/tidb
git remote add upstream https://github.com/go-jvm/gojvm.git
# or: git remote add upstream git@github.com:go-jvm/gojvm.git

# Never push to upstream master since you do not have write access.
git remote set-url --push upstream no_push

# Confirm that your remotes make sense:
# It should look like:
# origin    git@github.com:$(user)/gojvm.git (fetch)
# origin    git@github.com:$(user)/gojvm.git (push)
# upstream  https://github.com/go-jvm/gojvm.git (fetch)
# upstream  no_push (push)
git remote -v
```


### Step 3: Branch

Get your local master up to date:

```sh
cd $working_dir/gojvm
git fetch upstream
git checkout master

git rebase upstream/master
```

Branch from master:

```sh
git checkout -b myfeature
```

### Step 4: Develop

#### Edit the code

#### Run Test

### Step 5: Keep your branch in sync

```sh
# While on your myfeature branch.
git fetch upstream
git rebase upstream/master
```

### Step 6: Commit

Commit your changes.

```sh
git commit
```
### Step 7: Push

When ready to review (or just to establish an offsite backup or your work),
push your branch to your fork on `github.com`:

```sh
git push -f origin myfeature
```

### Step 8: Create a pull request

1. Visit your fork at https://github.com/$user/gojvm (replace `$user` obviously).
2. Click the `Compare & pull request` button next to your `myfeature` branch.

### Step 9: Get a code review
