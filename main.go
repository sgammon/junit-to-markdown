package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"github.com/joshdk/go-junit"
)

type testSuite struct {
	name string
}

func main() {
	var noFailingBuilds = true

	for _, file := range os.Args[1:] {
		f, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}

		suites, err := junit.Ingest(f)
		if err != nil {
			log.Fatal(err)
		}

		for _, suite := range suites {
			if suite.Totals.Failed > 0 {
				noFailingBuilds = false
			}
		}

		tmpl, err := template.New("").Parse(`
{{- range .}}
{{- if gt .Totals.Failed 0 }}
### {{.Name}}
|Success|Test|
|-------|----|
{{- range .Tests}}
{{- if .Error}}
|:x:|{{.Name}}|
{{- end}}
{{- end}}
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

	if noFailingBuilds {
		fmt.Println("### :white_check_mark: All tests passed!")
	}
}
