package stdlib

// Code generated by 'goexports text/template'. DO NOT EDIT.

import (
	"reflect"
	"text/template"
)

func init() {
	Value["text/template"] = map[string]reflect.Value{
		"HTMLEscape":       reflect.ValueOf(template.HTMLEscape),
		"HTMLEscapeString": reflect.ValueOf(template.HTMLEscapeString),
		"HTMLEscaper":      reflect.ValueOf(template.HTMLEscaper),
		"IsTrue":           reflect.ValueOf(template.IsTrue),
		"JSEscape":         reflect.ValueOf(template.JSEscape),
		"JSEscapeString":   reflect.ValueOf(template.JSEscapeString),
		"JSEscaper":        reflect.ValueOf(template.JSEscaper),
		"Must":             reflect.ValueOf(template.Must),
		"New":              reflect.ValueOf(template.New),
		"ParseFiles":       reflect.ValueOf(template.ParseFiles),
		"ParseGlob":        reflect.ValueOf(template.ParseGlob),
		"URLQueryEscaper":  reflect.ValueOf(template.URLQueryEscaper),
	}

	Type["text/template"] = map[string]reflect.Type{
		"ExecError": reflect.TypeOf((*template.ExecError)(nil)).Elem(),
		"FuncMap":   reflect.TypeOf((*template.FuncMap)(nil)).Elem(),
		"Template":  reflect.TypeOf((*template.Template)(nil)).Elem(),
	}
}
