package main

import (
	"encoding/xml"
)

type Protocol struct {
	XMLName   xml.Name `xml:"Protocol"`
	Copyright string   `xml:"Copyright"`
	Interface []Interface
}

type Interface struct {
	Name        string `xml:"name,attr"`
	Description string
	Request     []Request
	Event       []Event
	Enum        []Enum
}

type Request struct {
	Name        string `xml:"name,attr"`
	Description string `xml:"Description"`
	Arg         []Arg
}

type Event struct {
	Name        string `xml:"name,attr"`
	Description string `xml:"Description"`
	Arg         []Arg
}

type Enum struct {
	Name        string `xml:"name,attr"`
	Description string `xml:"Description"`
	Entry       []Entry
}

type Entry struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type Arg struct {
	Name      string `xml:"name,attr"`
	Type      string `xml:"type,attr"`
	Interface string `xml:"interface,attr"`
	AllowNull string `xml:"allow-null,attr"`
}

var (
	pkgTemplate string = `package gowl

	{{if .Description}}/*
	{{.Description}}
	*/{{end}}
	type {{.Name}} struct{}

		{{range $e := $.Enum}}
			{{if .Description}}/*
			{{.Description}}*/{{end}}
			
			const(
				{{range .Entry}}{{$e.Name}}_{{.Name}} = {{.Value}}
				{{end}}
			)
		{{end}}

		{{range $.Request}}
			{{if .Description}}/*
			{{.Description}}*/{{end}}
			func (*{{$.Name}}) {{.Name}}({{range .Arg}}{{.Name}} {{.Type}},{{end}}){
		}
		{{end}}
	
		{{range $.Event}}
			{{if .Description}}/*
			{{.Description}}*/{{end}}
				func (*{{$.Name}}) {{.Name}}({{range .Arg}}{{.Name}} {{.Type}},{{end}}){
			}
		{{end}}
	`
)
