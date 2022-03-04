package extension

import (
	"fmt"
	"strings"
)

type Extension interface {
	ShouldEncode(value interface{}) (valid bool, asPointer bool)
	Encode(value interface{}) interface{}
	Decode(value interface{}, asPointer bool) interface{}
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
		if valid, asPointer := ext.ShouldEncode(value); valid {
			prefix := ""
			if asPointer {
				prefix = "*"
			}
			return fmt.Sprintf("%s%s~%s", prefix, name, key), ext.Encode(value)
		}
	}
	return key, value
}

func DecodeValue(key string, encodedValue interface{}) (string, interface{}) {
	keyParts := strings.Split(key, "~")
	if len(keyParts) != 2 {
		return key, encodedValue
	}
	extensionName := keyParts[0]
	isPointer := strings.HasPrefix(extensionName, "*")
	if isPointer {
		extensionName = strings.Replace(extensionName, "*", "", 1)
	}
	ext, ok := extensions[extensionName]
	if !ok {
		return keyParts[1], encodedValue
	}

	return keyParts[1], ext.Decode(encodedValue, isPointer)
}
