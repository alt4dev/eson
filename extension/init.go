package extension

import (
	"fmt"
	"reflect"
	"strings"
)

type Extension interface {
	ShouldEncode(value interface{}) bool
	Encode(value interface{}) interface{}
	Decode(value interface{}) interface{}
}

func EncodeValue(key string, value interface{}, extensions []Extension) (string, interface{}) {
	for _, ext := range extensions {
		if ext.ShouldEncode(value) {
			name := reflect.TypeOf(ext).Name()
			return fmt.Sprintf("%s~%s", name, key), ext.Encode(value)
		}
	}
	return key, value
}

func DecodeValue(key string, encodedValue interface{}, extensions []Extension) (string, interface{}) {
	keyParts := strings.Split(key, "~")
	if len(keyParts) != 2 {
		return key, encodedValue
	}

	extensionsMap := make(map[string]Extension)
	for _, extension := range extensions {
		extensionsMap[reflect.TypeOf(extension).Name()] = extension
	}
	ext, ok := extensionsMap[keyParts[0]]
	if !ok {
		return keyParts[1], encodedValue
	}
	return keyParts[1], ext.Decode(encodedValue)
}
