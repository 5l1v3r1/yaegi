package stdlib

// Code generated by 'goexports crypto/rsa'. DO NOT EDIT.

import (
	"crypto/rsa"
	"reflect"
)

func init() {
	Value["crypto/rsa"] = map[string]reflect.Value{
		"DecryptOAEP":               reflect.ValueOf(rsa.DecryptOAEP),
		"DecryptPKCS1v15":           reflect.ValueOf(rsa.DecryptPKCS1v15),
		"DecryptPKCS1v15SessionKey": reflect.ValueOf(rsa.DecryptPKCS1v15SessionKey),
		"EncryptOAEP":               reflect.ValueOf(rsa.EncryptOAEP),
		"EncryptPKCS1v15":           reflect.ValueOf(rsa.EncryptPKCS1v15),
		"ErrDecryption":             reflect.ValueOf(rsa.ErrDecryption),
		"ErrMessageTooLong":         reflect.ValueOf(rsa.ErrMessageTooLong),
		"ErrVerification":           reflect.ValueOf(rsa.ErrVerification),
		"GenerateKey":               reflect.ValueOf(rsa.GenerateKey),
		"GenerateMultiPrimeKey":     reflect.ValueOf(rsa.GenerateMultiPrimeKey),
		"PSSSaltLengthAuto":         reflect.ValueOf(rsa.PSSSaltLengthAuto),
		"PSSSaltLengthEqualsHash":   reflect.ValueOf(rsa.PSSSaltLengthEqualsHash),
		"SignPKCS1v15":              reflect.ValueOf(rsa.SignPKCS1v15),
		"SignPSS":                   reflect.ValueOf(rsa.SignPSS),
		"VerifyPKCS1v15":            reflect.ValueOf(rsa.VerifyPKCS1v15),
		"VerifyPSS":                 reflect.ValueOf(rsa.VerifyPSS),
	}

	Type["crypto/rsa"] = map[string]reflect.Type{
		"CRTValue":               reflect.TypeOf((*rsa.CRTValue)(nil)).Elem(),
		"OAEPOptions":            reflect.TypeOf((*rsa.OAEPOptions)(nil)).Elem(),
		"PKCS1v15DecryptOptions": reflect.TypeOf((*rsa.PKCS1v15DecryptOptions)(nil)).Elem(),
		"PSSOptions":             reflect.TypeOf((*rsa.PSSOptions)(nil)).Elem(),
		"PrecomputedValues":      reflect.TypeOf((*rsa.PrecomputedValues)(nil)).Elem(),
		"PrivateKey":             reflect.TypeOf((*rsa.PrivateKey)(nil)).Elem(),
		"PublicKey":              reflect.TypeOf((*rsa.PublicKey)(nil)).Elem(),
	}
}
