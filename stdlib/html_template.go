package stdlib

// Code generated by 'goexports html/template'. DO NOT EDIT.

import (
	"html/template"
	"reflect"
)

func init() {
	Value["html/template"] = map[string]reflect.Value{
		"ErrAmbigContext":      reflect.ValueOf(template.ErrAmbigContext),
		"ErrBadHTML":           reflect.ValueOf(template.ErrBadHTML),
		"ErrBranchEnd":         reflect.ValueOf(template.ErrBranchEnd),
		"ErrEndContext":        reflect.ValueOf(template.ErrEndContext),
		"ErrNoSuchTemplate":    reflect.ValueOf(template.ErrNoSuchTemplate),
		"ErrOutputContext":     reflect.ValueOf(template.ErrOutputContext),
		"ErrPartialCharset":    reflect.ValueOf(template.ErrPartialCharset),
		"ErrPartialEscape":     reflect.ValueOf(template.ErrPartialEscape),
		"ErrPredefinedEscaper": reflect.ValueOf(template.ErrPredefinedEscaper),
		"ErrRangeLoopReentry":  reflect.ValueOf(template.ErrRangeLoopReentry),
		"ErrSlashAmbig":        reflect.ValueOf(template.ErrSlashAmbig),
		"HTMLEscape":           reflect.ValueOf(template.HTMLEscape),
		"HTMLEscapeString":     reflect.ValueOf(template.HTMLEscapeString),
		"HTMLEscaper":          reflect.ValueOf(template.HTMLEscaper),
		"IsTrue":               reflect.ValueOf(template.IsTrue),
		"JSEscape":             reflect.ValueOf(template.JSEscape),
		"JSEscapeString":       reflect.ValueOf(template.JSEscapeString),
		"JSEscaper":            reflect.ValueOf(template.JSEscaper),
		"Must":                 reflect.ValueOf(template.Must),
		"New":                  reflect.ValueOf(template.New),
		"OK":                   reflect.ValueOf(template.OK),
		"ParseFiles":           reflect.ValueOf(template.ParseFiles),
		"ParseGlob":            reflect.ValueOf(template.ParseGlob),
		"URLQueryEscaper":      reflect.ValueOf(template.URLQueryEscaper),
	}

	Type["html/template"] = map[string]reflect.Type{
		"CSS":       reflect.TypeOf((*template.CSS)(nil)).Elem(),
		"Error":     reflect.TypeOf((*template.Error)(nil)).Elem(),
		"ErrorCode": reflect.TypeOf((*template.ErrorCode)(nil)).Elem(),
		"FuncMap":   reflect.TypeOf((*template.FuncMap)(nil)).Elem(),
		"HTML":      reflect.TypeOf((*template.HTML)(nil)).Elem(),
		"HTMLAttr":  reflect.TypeOf((*template.HTMLAttr)(nil)).Elem(),
		"JS":        reflect.TypeOf((*template.JS)(nil)).Elem(),
		"JSStr":     reflect.TypeOf((*template.JSStr)(nil)).Elem(),
		"Srcset":    reflect.TypeOf((*template.Srcset)(nil)).Elem(),
		"Template":  reflect.TypeOf((*template.Template)(nil)).Elem(),
		"URL":       reflect.TypeOf((*template.URL)(nil)).Elem(),
	}
}
