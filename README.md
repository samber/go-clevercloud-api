go-clevercloud-api
====

Go library for Clever-Cloud api

## About

This library has been developed for a Terraform provider. Some attributes of the official Clever-Cloud API may not be exported, because not used in Terraform.

## Available CC features (2016-11-30)

Supports following runtimes:

 - `docker`: for Docker-based applications
 - `go`: for Go applications
 - `gradle`: for applications launched with gradle
 - `jar`: for applications deployed as standalone jar files
 - `war`: for applications deployed as war files
 - `play1`: for Play1 applications
 - `play2`: for Play2 applications
 - `sbt`: for applications launched with SBT
 - `maven`: for applications launched with maven
 - `node`: for node.js applications
 - `php`: for PHP applications
 - `python`: for python27 applications
 - `ruby`: for ruby applications
 - `static`: for static (HTML only) websites
 - `haskell`: for Haskell applications
 - `rust`: for Rust applications

Supports following regions:

 - `par` (for Paris)
 - `mtl` (for Montreal)

Supports following instance size:

- `pico`: 256Mo RAM / 1 CPU
- `nano`: 512Mo RAM / 1 CPU
- `XS`: 1024Mo RAM / 1 CPU
- `S`: 2048Mo RAM / 2 CPU
- `M`: 4096Mo RAM / 4 CPU
- `L`: 8192Mo RAM / 6 CPU
- `XL`: 16384Mo RAM / 8 CPU

Supports following addons:

- `postgresql-addon`
- `mysql-addon`
- `redis-addon`
- `mongodb-addon`
- `cellar-addon`
- `trace`
- `socks-addon`
- `fs-bucket`

## Running examples

```sh
go get github.com/samber/go-clevercloud-api

cd examples/
go build -o cc-api-examples *.go
./cc-api-examples
```

## Contributions

Please contribute! I'll review your PRs ASAP.