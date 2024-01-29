#! /usr/bin/env bash
set -uvx
set -e
get_latest_release () {
    gh release list --repo $1 | head -1 | sed 's/|/ /' | awk '{print $1, $8}'
}
rm -rf go.mod go.sum
go mod init example.com/m
go get github.com/lang-library/go-global@`get_latest_release lang-library/go-global`
go get github.com/lang-library/go-winlib@`get_latest_release lang-library/go-winlib`
go get github.com/lang-library/go-lib01@`get_latest_release lang-library/go-lib01`
go get github.com/lang-library/go-common-file-dialog@`get_latest_release lang-library/go-common-file-dialog`
go get golang.org/x/sys/windows/mkwinsyscall
go get github.com/holiman/uint256
go mod tidy
cat ./go.mod
