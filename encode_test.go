package eson

import (
	"fmt"
	"testing"
	"time"
)

type testStruct struct {
	Name       string    `json:"name"`
	DOB        time.Time `json:"date_of_birth"`
	Roles      []string  `json:"roles"`
	Registered time.Time
}

func TestEncode(t *testing.T) {
	data := testStruct{
		Name:       "Jane Doe",
		DOB:        time.Now(),
		Roles:      []string{"admin", "client"},
		Registered: time.Now(),
	}

	encodedData, err := Encode(data)
	if err != nil {
		t.Error(err)
		return
	}

	expectedOutput := fmt.Sprintf(`{"EsonDatetime~Registered":%v,"EsonDatetime~date_of_birth":%v,"name":"Jane Doe","roles":["admin","client"]}`, data.Registered.UnixMilli(), data.DOB.UnixMilli())

	if expectedOutput != encodedData {
		t.Error("UnExpected JSON output")
		t.Error("Expected:", expectedOutput)
		t.Error("Found:   ", encodedData)
		return
	}
}

func TestEncodeList(t *testing.T) {
	time1 := time.Now()
	time2 := time.Now().Add(time.Second * 10)
	time3 := time.Now().Add(time.Second * 30)

	times := []time.Time{time1, time2, time3}
	encodedData, err := Encode(times)
	if err != nil {
		t.Error(err)
		return
	}

	expectedOutput := fmt.Sprintf(`[{"EsonDatetime~":%v},{"EsonDatetime~":%v},{"EsonDatetime~":%v}]`, time1.UnixMilli(), time2.UnixMilli(), time3.UnixMilli())

	if expectedOutput != encodedData {
		t.Error("UnExpected JSON output")
		t.Error("Expected:", expectedOutput)
		t.Error("Found:   ", encodedData)
		return
	}
}
