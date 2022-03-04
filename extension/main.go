package extension

import (
	"fmt"
	"strings"
)

type Extension interface {
	ShouldEncode(value interface{}) bool
	Encode(value interface{}) interface{}
	Decode(value interface{}) interface{}
}

var extensions map[string]Extension

func init() {
	extensions = make(map[string]Extension)
	AddExtension("EsonDatetime", DateTimeExtension{})
}

func AddExtension(name string, extension Extension) {
	extensions[name] = extension
}

func EncodeValue(key string, value interface{}) (string, interface{}) {
	for name, ext := range extensions {
		if ext.ShouldEncode(value) {
			return fmt.Sprintf("%s~%s", name, key), ext.Encode(value)
		}
	}
	return key, value
}

func DecodeValue(key string, encodedValue interface{}) (string, interface{}) {
	keyParts := strings.Split(key, "~")
	if len(keyParts) != 2 {
		return key, encodedValue
	}
	ext, ok := extensions[keyParts[0]]
	if !ok {
		return keyParts[1], encodedValue
	}
	return keyParts[1], ext.Decode(encodedValue)
}
