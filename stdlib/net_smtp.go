package stdlib

// Code generated by 'goexports net/smtp'. DO NOT EDIT.

import (
	"net/smtp"
	"reflect"
)

func init() {
	Value["net/smtp"] = map[string]reflect.Value{
		"CRAMMD5Auth": reflect.ValueOf(smtp.CRAMMD5Auth),
		"Dial":        reflect.ValueOf(smtp.Dial),
		"NewClient":   reflect.ValueOf(smtp.NewClient),
		"PlainAuth":   reflect.ValueOf(smtp.PlainAuth),
		"SendMail":    reflect.ValueOf(smtp.SendMail),
	}

	Type["net/smtp"] = map[string]reflect.Type{
		"Auth":       reflect.TypeOf((*smtp.Auth)(nil)).Elem(),
		"Client":     reflect.TypeOf((*smtp.Client)(nil)).Elem(),
		"ServerInfo": reflect.TypeOf((*smtp.ServerInfo)(nil)).Elem(),
	}
}
