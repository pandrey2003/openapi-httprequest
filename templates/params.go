package templates

import "text/template"

// Definition defines a type definition to be written to params.go
type Definition struct {
	Name       string
	TypeStr    string
	Tag        string
	DocComment string
	Properties DefinitionList
}

type DefinitionList []*Definition

func (s DefinitionList) Len() int           { return len(s) }
func (s DefinitionList) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s DefinitionList) Less(i, j int) bool { return s[i].Name < s[j].Name }

type HandlerList []*Handler

func (s HandlerList) Len() int           { return len(s) }
func (s HandlerList) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s HandlerList) Less(i, j int) bool { return s[i].Name < s[j].Name }

// Handler defines a httprequest handler function to be written to handlers.go
type Handler struct {
	Name       string
	DocComment string
	Request    string
	Response   string
}

var Params = template.Must(template.New("").Parse(`
// Code generated by openapi-httprequest. DO NOT EDIT.

package {{.Pkg}}

import(
    {{range .Imports}} "{{.}}" {{println ""}} {{end}}

    httprequest "gopkg.in/httprequest.v1"
)

type APIHandler interface {
	{{- range .Handlers }}
	{{ if .DocComment}}{{.DocComment}}{{end}}
	{{.Name}}(httprequest.Params, *{{.Request}}) ({{if .Response}}*{{.Response}}, {{end}}error)
	{{end}}
}

{{range .Types}}
{{- if .DocComment}}{{.DocComment}}{{end}}
type {{.Name}} {{if .TypeStr}}{{.TypeStr}}{{else}}struct {
  {{range .Properties}}
  {{- if .DocComment}}{{.DocComment}}
  {{end}}
  {{- .Name}} {{.TypeStr}} {{.Tag}}
  {{end}}
}{{end}}
{{end}}`))
