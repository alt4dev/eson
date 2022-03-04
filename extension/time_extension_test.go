package extension

import (
	"testing"
	"time"
)

func TestDateTimeExtension_ShouldEncode(t *testing.T) {
	ext := DateTimeExtension{}
	// Test that should encode is only valid for time.
	if valid, _ := ext.ShouldEncode(time.Now()); !valid {
		t.Error("A Time object should be encoded")
	}

	if ok, _ := ext.ShouldEncode(nil); ok {
		t.Error("A nil value should not be encoded")
	}

	strOk, _ := ext.ShouldEncode("")
	intOk, _ := ext.ShouldEncode(10)
	if strOk || intOk {
		t.Error("Integers and strings should not be encoded")
	}
}

func TestDateTimeExtension_Encode(t *testing.T) {
	ext := DateTimeExtension{}
	now := time.Now()
	if int64(now.UnixMilli()) != ext.Encode(now).(int64) {
		t.Error("Expected time in microseconds doesn't match")
	}
}

func TestDateTimeExtension_Decode(t *testing.T) {
	ext := DateTimeExtension{}
	// Timestamp
	expectedTime := time.Unix(1598887850, 0)
	decodedTime := ext.Decode(float64(1598887850000), false).(time.Time)
	if expectedTime != decodedTime {
		t.Error("Timestamp not decoded to the expected time")
		t.Error(expectedTime, expectedTime.UnixNano())
		t.Error(decodedTime)
	}
}
