package main

import (
	"os"
	"strings"
	"text/template"
)

const text = `
{{/* This is a zoo template */}}
{{with .Name}}Welcome to {{.}}{{end}}
There are {{len .Animals}} animals, they are: 
{{range .Animals}}
{{- . | upper -}},
{{end}}
{{if gt (len .Zookeepers) 0}}
There are {{len .Zookeepers}} zookeepers, they are:
{{range $no, $name := .Zookeepers}}
{{printf "%03d" $no}}: {{$name -}}
{{end}}
{{end}}
{{block "Welcome" .Name}}You're welcome to visit {{.}} next time!{{end}}
`

type Zoo struct {
	Name       string
	Animals    []string
	Zookeepers map[int]string
}

func main() {
	// template
	tpl := template.Must(template.New("zoo").Funcs(template.FuncMap{
		"upper": func(s string) string { // self-defined functions
			return strings.ToUpper(s)
		},
	}).Parse(text))

	// zookeepers
	zooKeepers := map[int]string{
		0: "Alan",
		1: "Larry",
		2: "Alice",
	}

	// zoo
	zoo := &Zoo{
		"Beijing Zoo",
		[]string{"elephant", "tiger", "dolphin"},
		zooKeepers,
	}

	// execute
	tpl.Execute(os.Stdout, zoo)
}
