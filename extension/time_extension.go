package extension

import (
	"time"
)

type DateTimeExtension struct{}

func (ext DateTimeExtension) ShouldEncode(value interface{}) (bool, bool) {
	if _, ok := value.(time.Time); ok {
		return true, false
	}
	if _, ok := value.(*time.Time); ok {
		return true, true
	}
	return false, false
}

func (ext DateTimeExtension) Encode(value interface{}) interface{} {
	val, ok := value.(time.Time)
	if !ok {
		val = *(value.(*time.Time))
	}
	return val.UnixMilli()
}

func (ext DateTimeExtension) Decode(encodedValue interface{}, asPointer bool) interface{} {
	val := encodedValue.(float64)
	decodedValue := time.Unix(0, int64(val)*1000000)
	if asPointer {
		return &decodedValue
	}
	return decodedValue
}
