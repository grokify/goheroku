# GoHeroku - Heroku Golang Helper

[![Go CI][go-ci-svg]][go-ci-url]
[![Go Lint][go-lint-svg]][go-lint-url]
[![Go SAST][go-sast-svg]][go-sast-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![Visualization][viz-svg]][viz-url]
[![License][license-svg]][license-url]

 [go-ci-svg]: https://github.com/grokify/goheroku/actions/workflows/go-ci.yaml/badge.svg?branch=main
 [go-ci-url]: https://github.com/grokify/goheroku/actions/workflows/go-ci.yaml
 [go-lint-svg]: https://github.com/grokify/goheroku/actions/workflows/go-lint.yaml/badge.svg?branch=main
 [go-lint-url]: https://github.com/grokify/goheroku/actions/workflows/go-lint.yaml
 [go-sast-svg]: https://github.com/grokify/goheroku/actions/workflows/go-sast-codeql.yaml/badge.svg?branch=main
 [go-sast-url]: https://github.com/grokify/goheroku/actions/workflows/go-sast-codeql.yaml
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/goheroku
 [goreport-url]: https://goreportcard.com/report/github.com/grokify/goheroku
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/grokify/goheroku
 [docs-godoc-url]: https://pkg.go.dev/github.com/grokify/goheroku
 [viz-svg]: https://img.shields.io/badge/visualizaton-Go-blue.svg
 [viz-url]: https://mango-dune-07a8b7110.1.azurestaticapps.net/?repo=grokify%2Fgoheroku
 [loc-svg]: https://tokei.rs/b1/github/grokify/goheroku
 [repo-url]: https://github.com/grokify/goheroku
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/grokify/goheroku/blob/master/LICENSE

CLI app to create boilerplate for a Heroku Go app. These files will also enable the "Deploy to Heroku" button to function.

## Installation

```
$ go get github.com/grokify/goheroku
```

## Usage

```
$ goheroku my-project
Writing file 1: my-project/.env
Writing file 2: my-project/app.json
Writing file 3: my-project/Dockerfile
Writing file 4: my-project/heroku.yml
Writing file 5: my-project/Makefile
Writing file 6: my-project/Procfile
DONE
```

You will need to use `go mod init` to create your `go.mod` file. Add the following Heroku directive to your `go.mod` file above the go version number line.

```
// +heroku goVersion go1.16
go 1.16
```

## Updating Project Docs

Once your repo is ready for Heroku, you can add various options to "Deploy to Heroku" to your `README.md` file. The following includes adding a shield, button, and a CLI manual mode.

### Shield

The following adds a shield that looks like ![](https://img.shields.io/badge/%E2%86%91_deploy-Heroku-7056bf.svg?style=flat) to your file.

```markdown
[![Heroku][heroku-svg]][heroku-url]
 [heroku-svg]: https://img.shields.io/badge/%E2%86%91_deploy-Heroku-7056bf.svg?style=flat
 [heroku-url]: https://heroku.com/deploy
```

### Button

The following adds a button that looks like ![](https://www.herokucdn.com/deploy/button.svg) to your file.

```markdown
[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)
```

### Manual

```sh
$ heroku create
$ git push heroku master
$ heroku open
```

## More Info

Find information on deploying Go apps on Heroku here:

* Developer Guide: https://devcenter.heroku.com/articles/go-support
* Example Project: https://github.com/heroku/go-getting-started



=====


