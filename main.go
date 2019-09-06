package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"html/template"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	var (
		templatePath = ""
		dataPath     = ""
		outputPath   = ""
	)
	flag.StringVar(&templatePath, "template.path", "", "A YML template")
	flag.StringVar(&dataPath, "data.path", "", "A YML data path")
	flag.StringVar(&outputPath, "output.path", "", "A YML path for output")

	flag.Parse()

	tpl := template.Must(template.ParseFiles(templatePath))

	dataFile, err := os.Open(dataPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(dataFile)

	ctx := map[string]interface{}{}

	err = yaml.NewDecoder(dataFile).Decode(ctx)
	if err != nil {
		log.Fatal(err)
	}

	targetFile, err := os.Create(outputPath)
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.Execute(targetFile, ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Generated", outputPath)
}
