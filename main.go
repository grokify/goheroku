package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/grokify/goheroku/templates"
	"github.com/grokify/gotilla/os/osutil"
)

type File struct {
	ProductSlug  string
	TemplateFunc func(string) string
}

func main() {
	projectSlug := ""

	argsWithoutProg := os.Args
	if len(argsWithoutProg) > 1 {
		projectSlug = strings.TrimSpace(argsWithoutProg[1])
	}
	if len(projectSlug) == 0 {
		fmt.Println("Please specify a project name: goheroku <projectName>")
		os.Exit(1)
	} else if regexp.MustCompile(`\s`).MatchString(projectSlug) {
		fmt.Println("Please do not spaces in project names.")
		os.Exit(1)
	}

	exists, err := osutil.Exists(projectSlug)
	if err != nil {
		log.Fatal(err)
	} else if exists {
		fmt.Printf("Projects [%v] exists, exiting...\n", projectSlug)
		os.Exit(1)
	}

	err = os.Mkdir(projectSlug, 0755)
	if err != nil {
		panic(err)
	}

	files := []File{{
		ProductSlug:  "app.json",
		TemplateFunc: templates.AppJsonTemplate}, {
		ProductSlug:  "Dockerfile",
		TemplateFunc: templates.DockerfileTemplate}, {
		ProductSlug:  "heroku.yml",
		TemplateFunc: templates.HerokuYmlTemplate}, {
		ProductSlug:  "Makefile",
		TemplateFunc: templates.MakefileTemplate}, {
		ProductSlug:  "Procfile",
		TemplateFunc: templates.ProcfileTemplate}}

	for i, file := range files {
		str := file.TemplateFunc(projectSlug)
		filename := filepath.Join(projectSlug, file.ProductSlug)
		err := ioutil.WriteFile(filename, []byte(str), 0644)
		fmt.Printf("Writing file %v: %v", i+1, filename)
		if err != nil {
			panic(err)
		}
		fmt.Printf("\n")
	}

	fmt.Println("DONE")
}
