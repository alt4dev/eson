package extension

import (
	"time"
)

type DateTimeExtension struct {}

func (ext DateTimeExtension) ShouldEncode(value interface{}) bool {
	_, ok := value.(time.Time)
	return ok
}

func (ext DateTimeExtension) Encode(value interface{}) interface{} {
	val := value.(time.Time)
	return val.UnixNano() / 1000
}

func (ext DateTimeExtension) Decode(encodedValue interface{}) interface{}  {
	val := encodedValue.(float64)
	return time.Unix(0, int64(val) * 1000)
}
