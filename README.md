vetissue
========

This repository contains code that can be used to reproduce an issue with `go
vet` on Go version 1.10.

Running `go vet vetissue_test.go` exits with a zero status on go `1.9.4`.

Running `go vet vetissue_test.go` exits with a non-zero status on go `1.10`.
