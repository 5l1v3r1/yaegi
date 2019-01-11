package stdlib

// Code generated by 'goexports reflect'. DO NOT EDIT.

import (
	"reflect"
)

func init() {
	Value["reflect"] = map[string]reflect.Value{
		"Append":          reflect.ValueOf(reflect.Append),
		"AppendSlice":     reflect.ValueOf(reflect.AppendSlice),
		"Array":           reflect.ValueOf(reflect.Array),
		"ArrayOf":         reflect.ValueOf(reflect.ArrayOf),
		"Bool":            reflect.ValueOf(reflect.Bool),
		"BothDir":         reflect.ValueOf(reflect.BothDir),
		"Chan":            reflect.ValueOf(reflect.Chan),
		"ChanOf":          reflect.ValueOf(reflect.ChanOf),
		"Complex128":      reflect.ValueOf(reflect.Complex128),
		"Complex64":       reflect.ValueOf(reflect.Complex64),
		"Copy":            reflect.ValueOf(reflect.Copy),
		"DeepEqual":       reflect.ValueOf(reflect.DeepEqual),
		"Float32":         reflect.ValueOf(reflect.Float32),
		"Float64":         reflect.ValueOf(reflect.Float64),
		"Func":            reflect.ValueOf(reflect.Func),
		"FuncOf":          reflect.ValueOf(reflect.FuncOf),
		"Indirect":        reflect.ValueOf(reflect.Indirect),
		"Int":             reflect.ValueOf(reflect.Int),
		"Int16":           reflect.ValueOf(reflect.Int16),
		"Int32":           reflect.ValueOf(reflect.Int32),
		"Int64":           reflect.ValueOf(reflect.Int64),
		"Int8":            reflect.ValueOf(reflect.Int8),
		"Interface":       reflect.ValueOf(reflect.Interface),
		"Invalid":         reflect.ValueOf(reflect.Invalid),
		"MakeChan":        reflect.ValueOf(reflect.MakeChan),
		"MakeFunc":        reflect.ValueOf(reflect.MakeFunc),
		"MakeMap":         reflect.ValueOf(reflect.MakeMap),
		"MakeMapWithSize": reflect.ValueOf(reflect.MakeMapWithSize),
		"MakeSlice":       reflect.ValueOf(reflect.MakeSlice),
		"Map":             reflect.ValueOf(reflect.Map),
		"MapOf":           reflect.ValueOf(reflect.MapOf),
		"New":             reflect.ValueOf(reflect.New),
		"NewAt":           reflect.ValueOf(reflect.NewAt),
		"Ptr":             reflect.ValueOf(reflect.Ptr),
		"PtrTo":           reflect.ValueOf(reflect.PtrTo),
		"RecvDir":         reflect.ValueOf(reflect.RecvDir),
		"Select":          reflect.ValueOf(reflect.Select),
		"SelectDefault":   reflect.ValueOf(reflect.SelectDefault),
		"SelectRecv":      reflect.ValueOf(reflect.SelectRecv),
		"SelectSend":      reflect.ValueOf(reflect.SelectSend),
		"SendDir":         reflect.ValueOf(reflect.SendDir),
		"Slice":           reflect.ValueOf(reflect.Slice),
		"SliceOf":         reflect.ValueOf(reflect.SliceOf),
		"String":          reflect.ValueOf(reflect.String),
		"Struct":          reflect.ValueOf(reflect.Struct),
		"StructOf":        reflect.ValueOf(reflect.StructOf),
		"Swapper":         reflect.ValueOf(reflect.Swapper),
		"TypeOf":          reflect.ValueOf(reflect.TypeOf),
		"Uint":            reflect.ValueOf(reflect.Uint),
		"Uint16":          reflect.ValueOf(reflect.Uint16),
		"Uint32":          reflect.ValueOf(reflect.Uint32),
		"Uint64":          reflect.ValueOf(reflect.Uint64),
		"Uint8":           reflect.ValueOf(reflect.Uint8),
		"Uintptr":         reflect.ValueOf(reflect.Uintptr),
		"UnsafePointer":   reflect.ValueOf(reflect.UnsafePointer),
		"ValueOf":         reflect.ValueOf(reflect.ValueOf),
		"Zero":            reflect.ValueOf(reflect.Zero),
	}

	Type["reflect"] = map[string]reflect.Type{
		"ChanDir":      reflect.TypeOf((*reflect.ChanDir)(nil)).Elem(),
		"Kind":         reflect.TypeOf((*reflect.Kind)(nil)).Elem(),
		"Method":       reflect.TypeOf((*reflect.Method)(nil)).Elem(),
		"SelectCase":   reflect.TypeOf((*reflect.SelectCase)(nil)).Elem(),
		"SelectDir":    reflect.TypeOf((*reflect.SelectDir)(nil)).Elem(),
		"SliceHeader":  reflect.TypeOf((*reflect.SliceHeader)(nil)).Elem(),
		"StringHeader": reflect.TypeOf((*reflect.StringHeader)(nil)).Elem(),
		"StructField":  reflect.TypeOf((*reflect.StructField)(nil)).Elem(),
		"StructTag":    reflect.TypeOf((*reflect.StructTag)(nil)).Elem(),
		"Type":         reflect.TypeOf((*reflect.Type)(nil)).Elem(),
		"Value":        reflect.TypeOf((*reflect.Value)(nil)).Elem(),
		"ValueError":   reflect.TypeOf((*reflect.ValueError)(nil)).Elem(),
	}
}
