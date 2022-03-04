package extension

import (
	"testing"
	"time"
)

func TestDateTimeExtension_ShouldEncode(t *testing.T) {
	ext := DateTimeExtension{}
	// Test that should encode is only valid for time.
	if !ext.ShouldEncode(time.Now()) {
		t.Error("A Time object should be encoded")
	}

	if ext.ShouldEncode(nil) {
		t.Error("A nil value should not be encoded")
	}

	if ext.ShouldEncode("") || ext.ShouldEncode(10) {
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
	decodedTime := ext.Decode(float64(1598887850000)).(time.Time)
	if expectedTime != decodedTime {
		t.Error("Timestamp not decoded to the expected time")
		t.Error(expectedTime, expectedTime.UnixNano())
		t.Error(decodedTime)
	}
}
