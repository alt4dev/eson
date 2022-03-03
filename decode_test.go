package eson

import (
	"testing"
	"time"
)

func TestDecodeStruct(t *testing.T) {
	data := testStruct{
		Name:       "",
		email:      "",
		NextOfKin:  "",
		DOB:        time.Time{},
		Roles:      nil,
		Registered: time.Time{},
	}

	encodedData := `{"EsonDatetime~Registered":1645858786601,"EsonDatetime~date_of_birth":1645858786601,"name":"Jane Doe","roles":["admin","client"]}`

	err := Decode(encodedData, &data)

	if err != nil {
		t.Error(err)
		return
	}

	if data.Name != "Jane Doe" {
		t.Errorf("Name mismatch: '%v' != 'Jane Doe'", data.Name)
	}

	if data.Roles[0] != "admin" || data.Roles[1] != "client" {
		t.Errorf("Roles mismatch: '%v' != '[admin, client]'", data.Roles)
	}

	if !data.DOB.Equal(time.UnixMilli(1645858786601)) {
		t.Errorf("DOB mismatch: '%v' != '%v'", data.DOB, time.UnixMilli(1645858786601))
	}

	if !data.Registered.Equal(time.UnixMilli(1645858786601)) {
		t.Errorf("Registered mismatch: '%v' != '%v'", data.Registered, time.UnixMilli(1645858786601))
	}
}
