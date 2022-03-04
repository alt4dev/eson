package extension

import (
	"fmt"
	"strconv"
	"testing"
)

// An extension that converts integers to binary
type BinExtension struct{}

func (ext BinExtension) ShouldEncode(value interface{}) (bool, bool) {
	_, ok := value.(int)
	return ok, false
}

func (ext BinExtension) Encode(value interface{}) interface{} {
	val := int64(value.(int))
	return fmt.Sprint(strconv.FormatInt(val, 2))
}

func (ext BinExtension) Decode(value interface{}, asPointer bool) interface{} {
	i, _ := strconv.ParseInt(value.(string), 2, 64)
	return int(i)
}

func TestAddExtension(t *testing.T) {
	// Confirm
	if _, ok := extensions["EsonBinary"]; ok {
		t.Error("Unexpected extension found in extensions list")
		return
	}
	AddExtension("EsonBinary", BinExtension{})
	if _, ok := extensions["EsonBinary"]; !ok {
		t.Error("Extension not found in extensions list")
	}
}

func TestEncodeValue(t *testing.T) {
	// Encode a valid value
	encodedKey, encodedValue := EncodeValue("integer", 123)
	if encodedKey != "EsonBinary~integer" {
		t.Errorf("Unexpected key found. Expected `EsonBinary~integer`, but found `%s`", encodedKey)
	}
	if encodedValue.(string) != "1111011" {
		t.Errorf("Unexpected binary output, output %s", encodedValue)
	}

	// Encode an unsupported value
	encodedKey, encodedValue = EncodeValue("string", "Some string")
	if encodedKey != "string" || encodedValue.(string) != "Some string" {
		t.Errorf("Values and keys not supported by eson should not change")
	}
}

func TestDecodeValue(t *testing.T) {
	// Decode a valid value
	key, value := DecodeValue("EsonBinary~integer", "1111011")
	if key != "integer" || value.(int) != 123 {
		t.Error("Unexpected decoded value or key", key, value)
	}

	// Decoding with a missing extension returns the key and the value provided
	key, value = DecodeValue("EsonDateTime~datetime", "12242974287428")
	if key != "datetime" || value.(string) != "12242974287428" {
		t.Error("Unexpected decoded value or key", key, value)
	}

	// Decoding a non eson key should return the same value and key
	key, value = DecodeValue("some_float", 7.441)
	if key != "some_float" || value.(float64) != 7.441 {
		t.Error("Unexpected decoded value or key", key, value)
	}
}
