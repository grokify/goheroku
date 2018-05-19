# GoHeroku

[![Build Status][build-status-svg]][build-status-link]
[![Go Report Card][goreport-svg]][goreport-link]
[![Docs][docs-godoc-svg]][docs-godoc-link]
[![License][license-svg]][license-link]

 [build-status-svg]: https://api.travis-ci.org/grokify/goheroku.svg?branch=master
 [build-status-link]: https://travis-ci.org/grokify/goheroku
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/goheroku
 [goreport-link]: https://goreportcard.com/report/github.com/grokify/goheroku
 [docs-godoc-svg]: https://img.shields.io/badge/docs-godoc-blue.svg
 [docs-godoc-link]: https://godoc.org/github.com/grokify/goheroku
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-link]: https://github.com/grokify/goheroku/blob/master/LICENSE

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

## More Info

Find information on deploying Go apps on Heroku here:

* Developer Guide: https://devcenter.heroku.com/articles/go-support
* Example Project: https://github.com/heroku/go-getting-started