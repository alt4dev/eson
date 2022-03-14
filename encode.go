package eson

import (
	"encoding/json"
	"github.com/alt4dev/eson/extension"
	"reflect"
)

// EncodeWithTag is similar to Encode but allows an extra tag parameter instead of the default Tag
func EncodeWithTag(tag string, goObject interface{}, pretty bool, extensions ...extension.Extension) (string, error) {
	// Use default extensions if None are provided in the call.
	if len(extensions) == 0 {
		extensions = DefaultExtensions
	}
	goObject = preProcess(tag, goObject, extensions)

	var resp []byte
	var err error
	if pretty {
		resp, err = json.MarshalIndent(goObject, "", "    ")
	} else {
		resp, err = json.Marshal(goObject)
	}
	if err != nil {
		return "", err
	}
	return string(resp), nil
}

// Encode will convert the object provided to a JSON string.
// you can add extensions to use instead of the default extensions by including them as parameters to this function.
func Encode(goObject interface{}, pretty bool, extensions ...extension.Extension) (string, error) {
	return EncodeWithTag(*tagName, goObject, pretty, extensions...)
}

func preProcess(tagToUse string, goObject interface{}, extensions []extension.Extension) (processedObject interface{}) {
	objectValue := reflect.ValueOf(goObject)
	if objectValue.Type().Kind() == reflect.Ptr {
		if objectValue.IsZero() {
			return nil
		}
		objectValue = reflect.Indirect(objectValue)
		goObject = objectValue.Interface()
	}

	processedObject = goObject

	switch objectValue.Type().Kind() {
	case reflect.Map:
		processedObject = processMap(tagToUse, goObject.(map[string]interface{}), extensions)
		break
	case reflect.Struct:
		processedObject = processStruct(tagToUse, goObject, extensions)
		break
	case reflect.Slice, reflect.Array:
		goArray := make([]interface{}, 0)
		for i := 0; i < objectValue.Len(); i++ {
			goArray = append(goArray, objectValue.Index(i).Interface())
		}
		processedObject = processArray(tagToUse, goArray, extensions)
		break
	}
	return processedObject
}

func processMap(tagToUse string, theMap map[string]interface{}, extensions []extension.Extension) interface{} {
	encoded := make(map[string]interface{})
	for key, value := range theMap {
		encodedKey, encodedValue := extension.EncodeValue(key, value, extensions)
		if key == encodedKey {
			// Preprocess the value
			encodedValue = preProcess(tagToUse, value, extensions)
		}
		encoded[encodedKey] = encodedValue
	}
	return encoded
}

func processStruct(tagToUse string, theStruct interface{}, extensions []extension.Extension) interface{} {
	encoded := make(map[string]interface{})
	structType := reflect.TypeOf(theStruct)
	structValue := reflect.ValueOf(theStruct)

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		// Ignore non-exported fields
		if !field.IsExported() {
			continue
		}

		key := field.Name
		if jsonTag, ok := field.Tag.Lookup(tagToUse); ok {
			if jsonTag == "-" {
				continue
			}
			if jsonTag != "" {
				key = jsonTag
			}
		}

		value := structValue.FieldByName(field.Name)
		encoded[key] = value.Interface()
	}
	return preProcess(tagToUse, encoded, extensions)
}

func processArray(tagToUse string, theArray []interface{}, extensions []extension.Extension) interface{} {
	arrayOut := make([]interface{}, len(theArray))
	for i, value := range theArray {
		key, encodedValue := extension.EncodeValue("", value, extensions)
		if key != "" {
			encodedValue = map[string]interface{}{key: encodedValue}
		} else {
			// Preprocess the value
			encodedValue = preProcess(tagToUse, encodedValue, extensions)
		}
		arrayOut[i] = encodedValue
	}
	return arrayOut
}
