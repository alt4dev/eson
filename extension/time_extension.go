package extension

import (
	"time"
)

const MilliSecMultiplier int64 = 1000000

type EsonDatetime struct{}

func (ext EsonDatetime) ShouldEncode(value interface{}) bool {
	_, ok := value.(time.Time)
	return ok
}

func (ext EsonDatetime) Encode(value interface{}) interface{} {
	val := value.(time.Time)
	return val.UnixNano() / MilliSecMultiplier
}

func (ext EsonDatetime) Decode(encodedValue interface{}) interface{} {
	val := encodedValue.(float64)
	return time.Unix(0, int64(val)*MilliSecMultiplier)
}
