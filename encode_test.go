package eson

import (
	"fmt"
	"testing"
	"time"
)

type testStruct struct {
	Name       string    `json:"name" bigquery:"Norm"`
	email      string    // Skipped because it's not exported
	NextOfKin  string    `json:"-" bigquery:"-"` // Skipped in JSON
	DOB        time.Time `json:"date_of_birth" bigquery:"Data"`
	Roles      []string  `json:"roles" bigquery:"roles"`
	Registered time.Time
}

func TestEncode(t *testing.T) {
	SetTagName("bigquery")
	defer SetTagName(DefaultTag)
	now := time.Now()
	data := testStruct{
		Name:       "Jane Doe",
		DOB:        now,
		Roles:      []string{"admin", "client"},
		Registered: now,
	}

	encodedData, err := Encode(data, false)
	if err != nil {
		t.Error(err)
		return
	}

	expectedOutput := fmt.Sprintf(`{"EsonDatetime~Data":%v,"EsonDatetime~Registered":%v,"Norm":"Jane Doe","roles":["admin","client"]}`, data.Registered.UnixMilli(), data.DOB.UnixMilli())

	if expectedOutput != encodedData {
		t.Error("UnExpected JSON output")
		t.Error("Expected:", expectedOutput)
		t.Error("Found:   ", encodedData)
		return
	}

	// Test encoding data using a different tag than set.
	expectedJsonTagOutput := fmt.Sprintf(`{"EsonDatetime~Registered":%v,"EsonDatetime~date_of_birth":%v,"name":"Jane Doe","roles":["admin","client"]}`, data.Registered.UnixMilli(), data.DOB.UnixMilli())
	encodedWithJsonTag, err := EncodeWithTag("json", data, false)
	if err != nil {
		t.Error(err)
		return
	}

	if expectedJsonTagOutput != encodedWithJsonTag {
		t.Error("UnExpected JSON output")
		t.Error("Expected:", expectedJsonTagOutput)
		t.Error("Found:   ", encodedWithJsonTag)
		return
	}
}

func TestEncodeList(t *testing.T) {
	time1 := time.Now()
	time2 := time.Now().Add(time.Second * 10)
	time3 := time.Now().Add(time.Second * 30)

	times := []time.Time{time1, time2, time3}
	encodedData, err := Encode(&times, false)
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

	var timesArray [2]time.Time
	timesArray[0] = time1
	timesArray[1] = time2
	encodedData, err = Encode(timesArray, true)
	if err != nil {
		t.Error(err)
		return
	}

	expectedOutput = fmt.Sprintf(`[
    {
        "EsonDatetime~": %v
    },
    {
        "EsonDatetime~": %v
    }
]`, time1.UnixMilli(), time2.UnixMilli())

	if expectedOutput != encodedData {
		t.Error("UnExpected JSON output")
		t.Error("Expected:", expectedOutput)
		t.Error("Found:   ", encodedData)
		return
	}
}
