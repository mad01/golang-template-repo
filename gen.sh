#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

if [[ -z "${GOPATH}" ]]; then
    echo "missing env GOPATH in env"
    exit 1
fi

read -p 'Set github username: ' GITHUB_USERNAME
read -p 'Set github repo name: ' GITHUB_REPO

GITHUB_USERNAME_DIR=$GOPATH/src/github.com/$GITHUB_USERNAME
mkdir -p $GITHUB_USERNAME_DIR

cp -r template $GITHUB_REPO

sed -e "s/{{GITHUB_REPO}}/${GITHUB_REPO}/g; s/{{GITHUB_USERNAME}}/${GITHUB_USERNAME}/g;" template/Makefile > $GITHUB_REPO/Makefile
sed -e "s/{{GITHUB_REPO}}/${GITHUB_REPO}/g; s/{{GITHUB_USERNAME}}/${GITHUB_USERNAME}/g;" template/Dockerfile > $GITHUB_REPO/Dockerfile
sed -e "s/{{GITHUB_REPO}}/${GITHUB_REPO}/g;" template/template/deployment.yaml > $GITHUB_REPO/template/deployment.yaml
sed -e "s/{{GITHUB_REPO}}/${GITHUB_REPO}/g;" template/README.md > $GITHUB_REPO/README.md
sed -e "s/{{GITHUB_REPO}}/${GITHUB_REPO}/g;" template/.gitignore > $GITHUB_REPO/.gitignore
sed -e "s/{{GITHUB_REPO}}/${GITHUB_REPO}/g;" template/cmd.go > $GITHUB_REPO/cmd.go

cp -r $GITHUB_REPO $GITHUB_USERNAME_DIR
rm -r $GITHUB_REPO

pushd $GITHUB_USERNAME_DIR/$GITHUB_REPO
git init 
git add .
git commit -m 'init'
popd
