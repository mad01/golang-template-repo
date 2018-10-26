# golang template repo

golang project template repo. it will create a project in `$GOPATH`

### ONE-STEP INSTALL

Use curl (for Mac OS X):
```shell
$ bash -c "$(curl -fsSL https://raw.githubusercontent.com/mad01/golang-template-repo/master/get)"
```

or wget (for most UNIX platforms):
```shell
$ bash -c "$(wget https://raw.githubusercontent.com/mad01/golang-template-repo/master/get -O -)"
```


###  shell function to init repo
```shell

function generate-go-repo() {
    bash -c "$(curl -fsSL https://raw.githubusercontent.com/mad01/golang-template-repo/master/get)"
}
```
