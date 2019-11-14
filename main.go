package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"github.com/joshdk/go-junit"
)

func main() {
	fmt.Println("# Test results")

	for _, file := range os.Args[1:] {
		f, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}

		suites, err := junit.Ingest(f)
		if err != nil {
			log.Fatal(err)
		}

		tmpl, err := template.New("").Parse(`
{{- range .}}
### {{.Name}}
|Success|Test|
|-------|----|
{{- range .Tests}}
|{{if .Error}}:x:{{else}}:white_check_mark:{{end}}|{{.Name}}|
{{- end}}
{{- end}}
`)
		if err != nil {
			log.Fatal(err)
		}

		err = tmpl.Execute(os.Stdout, suites)
		if err != nil {
			log.Fatal()
		}

	}

}
