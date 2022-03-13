package eson

import (
	"encoding/json"
	"github.com/alt4dev/eson/extension"
	"reflect"
)

func Decode(jsonString string, destination interface{}) error {
	var encodedData interface{}
	err := json.Unmarshal([]byte(jsonString), &encodedData)
	if err != nil {
		return err
	}
	decodedData := processEncodedData(encodedData)

	setDataToDestination(decodedData, destination)
	return nil
}

func setDataToDestination(decodedData interface{}, destination interface{}) {
	destinationValue := reflect.ValueOf(destination)
	destinationType := reflect.TypeOf(destination)
	if destinationType.Kind() == reflect.Ptr {
		destinationValue = reflect.Indirect(destinationValue)
		destinationType = reflect.TypeOf(destinationValue.Interface())
	}

	decodedValue := reflect.ValueOf(decodedData)
	decodedType := reflect.TypeOf(decodedData)

	switch destinationType.Kind() {
	case reflect.Map:
		mapType := destinationType.Elem()
		// Read decoded data as map
		for _, key := range decodedValue.MapKeys() {
			decodedMapValue := decodedValue.MapIndex(key)
			valueToSet := reflect.New(mapType)
			setDataToDestination(decodedMapValue.Interface(), valueToSet.Interface())
			destinationValue.SetMapIndex(key, valueToSet)
		}
		return
	case reflect.Struct:
		if decodedType.Kind() != reflect.Map {
			destinationValue.Set(decodedValue)
			return
		}
		itemsMap := make(map[string]interface{})
		// Read decoded data as map
		for _, key := range decodedValue.MapKeys() {
			itemsMap[key.String()] = decodedValue.MapIndex(key).Interface()
		}

		for i := 0; i < destinationType.NumField(); i++ {
			field := destinationType.Field(i)
			if !field.IsExported() {
				continue
			}

			fieldName := field.Name
			if tag, ok := field.Tag.Lookup(tagName); ok {
				if tag == "-" {
					continue
				}
				fieldName = tag
			}

			decodedStructValue, ok := itemsMap[fieldName]
			if !ok {
				// Leave default value
				continue
			}

			fieldValue := destinationValue.Field(i)
			valueToSet := fieldValue.Interface()
			setDataToDestination(decodedStructValue, &valueToSet)
			value := reflect.ValueOf(valueToSet)
			destinationValue.Field(i).Set(value)
		}
		return
	case reflect.Array, reflect.Slice:
		sliceType := destinationType.Elem()
		newSlice := reflect.MakeSlice(destinationType, 0, 0)

		for i := 0; i < decodedValue.Len(); i++ {
			sliceValue := decodedValue.Index(i).Interface()
			valueToSet := reflect.New(sliceType)
			setDataToDestination(sliceValue, valueToSet.Interface())
			newSlice = reflect.Append(newSlice, reflect.Indirect(valueToSet))
		}

		destinationValue.Set(newSlice)
		return
	default:
		destinationValue.Set(decodedValue)
		return
	}
}

func processEncodedData(encodedData interface{}) interface{} {
	value := reflect.ValueOf(encodedData)
	processedData := encodedData
	switch value.Type().Kind() {
	case reflect.Map:
		encodedMap := make(map[string]interface{})
		for _, key := range value.MapKeys() {
			encodedMap[key.String()] = value.MapIndex(key).Interface()
		}
		processedData = processEncodedMap(encodedMap)
		break
	case reflect.Array, reflect.Slice:
		encodedArray := make([]interface{}, 0)
		for i := 0; i < value.Len(); i++ {
			encodedArray = append(encodedArray, value.Index(i).Interface())
		}
		processedData = processEncodedArray(encodedArray)
		break
	}
	return processedData
}

func processEncodedMap(encodedMap map[string]interface{}) interface{} {
	decodedMap := make(map[string]interface{})

	for encodedKey, encodedValue := range encodedMap {
		key, value := extension.DecodeValue(encodedKey, encodedValue)
		if key != encodedKey && key == "" {
			// From array value
			return value
		}
		decodedMap[key] = processEncodedData(value)
	}
	return decodedMap
}

func processEncodedArray(encodedArray []interface{}) interface{} {
	decodedArray := make([]interface{}, len(encodedArray))

	for i, encodedValue := range encodedArray {
		decodedArray[i] = processEncodedData(encodedValue)
	}

	return decodedArray
}
