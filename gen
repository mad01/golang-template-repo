#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

if [[ -z "${GOPATH}" ]]; then
	echo "missing env GOPATH in env"
	exit 1
fi

read -p 'github username: ' GITHUB_USERNAME
read -p 'github repo: ' GITHUB_REPO

GITHUB_USERNAME_DIR="${GOPATH}/src/github.com/${GITHUB_USERNAME}"
mkdir -p "${GITHUB_USERNAME_DIR}"

cp -r template "${GITHUB_REPO}"

find "${GITHUB_REPO}" -type f -print0 | xargs -0 sed -i -e \
	"s/{{GITHUB_REPO}}/${GITHUB_REPO}/g; s/{{GITHUB_USERNAME}}/${GITHUB_USERNAME}/g;"

cp -r "${GITHUB_REPO}" "${GITHUB_USERNAME_DIR}"
rm -r "${GITHUB_REPO}"

function setup-git() {
	git init
	git add .
	git commit -m 'init'
}

function setup-dep() {
	go get -u github.com/golang/dep/cmd/dep
	go get -u golang.org/x/tools/cmd/goimports
	dep ensure -v
}

pushd "${GITHUB_USERNAME_DIR}/${GITHUB_REPO}"
setup-git
setup-dep
popd
