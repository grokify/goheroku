# GoHeroku

CLI app to create boilerplate for a Heroku Go app. These files will allow the "Deploy to Heroku" button to function.

## Installation

```
$ go get github.com/grokfiy/goheroku
```

## Usage

```
$ goheroku my-project
Writing file 1: my-project/app.json
Writing file 2: my-project/Dockerfile
Writing file 3: my-project/heroku.yml
Writing file 4: my-project/Makefile
Writing file 5: my-project/Procfile
DONE
```

## More Info

Find information on deploying Go apps on Heroku here:

* Developer Guide: https://devcenter.heroku.com/articles/go-support
* Example Project: https://github.com/heroku/go-getting-started