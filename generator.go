package main 

import (
	"path/filepath"
	"flag"
	"os"
	"log"
    "html/template"
)

var (
	templateFolder string
	templateMask   string
	inputFile      string
	outputFile     string
)

func init() {	
	flag.StringVar(&templateFolder, "templates", "", "the folder where the extra templates reside")
	flag.StringVar(&templateMask, "mask", "_*", "the mask to use for extra templates")
	flag.StringVar(&inputFile, "input", "", "the primary input template")
	flag.StringVar(&outputFile, "output", "", "the output file")
}

func main() {
	flag.Parse()

	helpers, _ := filepath.Glob(filepath.Join(templateFolder, templateMask))
	files := append(helpers, inputFile)
	t, err := template.ParseFiles(files...)
	if err != nil {
		log.Fatal(err)
	}
		
	outFile, err := os.OpenFile(outputFile, os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	if err := t.Execute(outFile, nil); err != nil {
		log.Fatal(err)
	}
}