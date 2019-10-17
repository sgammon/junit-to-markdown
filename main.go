package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"

	"log"
	"os"
	"text/template"

	"github.com/joshdk/go-junit"
)


func main() {
	dir, err := ioutil.ReadDir(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("# Test results")

	for _, file := range dir {
		if strings.HasPrefix(file.Name(), "TEST") {
			f, err := ioutil.ReadFile(path.Join(os.Args[1], file.Name()))
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

}
