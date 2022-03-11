package eson

import (
	"encoding/json"
	"github.com/alt4dev/eson/extension"
	"reflect"
)

func Encode(goObject interface{}, pretty bool) (string, error) {
	goObject = preProcess(goObject)

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

func preProcess(goObject interface{}) (processedObject interface{}) {
	objectValue := reflect.ValueOf(goObject)
	if objectValue.Type().Kind() == reflect.Ptr {
		objectValue = reflect.Indirect(objectValue)
		goObject = objectValue.Interface()
	}

	processedObject = goObject

	switch objectValue.Type().Kind() {
	case reflect.Map:
		processedObject = processMap(goObject.(map[string]interface{}))
		break
	case reflect.Struct:
		processedObject = processStruct(goObject)
		break
	case reflect.Slice, reflect.Array:
		goArray := make([]interface{}, 0)
		for i := 0; i < objectValue.Len(); i++ {
			goArray = append(goArray, objectValue.Index(i).Interface())
		}
		processedObject = processArray(goArray)
		break
	}
	return processedObject
}

func processMap(theMap map[string]interface{}) interface{} {
	encoded := make(map[string]interface{})
	for key, value := range theMap {
		encodedKey, encodedValue := extension.EncodeValue(key, value)
		if key == encodedKey {
			// Preprocess the value
			encodedValue = preProcess(value)
		}
		encoded[encodedKey] = encodedValue
	}
	return encoded
}

func processStruct(theStruct interface{}) interface{} {
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
		if jsonTag, ok := field.Tag.Lookup("json"); ok {
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
	return preProcess(encoded)
}

func processArray(theArray []interface{}) interface{} {
	arrayOut := make([]interface{}, len(theArray))
	for i, value := range theArray {
		key, encodedValue := extension.EncodeValue("", value)
		if key != "" {
			encodedValue = map[string]interface{}{key: encodedValue}
		} else {
			// Preprocess the value
			encodedValue = preProcess(encodedValue)
		}
		arrayOut[i] = encodedValue
	}
	return arrayOut
}
