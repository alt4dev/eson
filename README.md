# [Extended JSON(ESON)](https://pkg.go.dev/mod/github.com/alt4dev/eson)

This is an implementation of https://github.com/Billcountry/eson in golang.

![Tests](https://github.com/alt4dev/eson/workflows/Tests/badge.svg?branch=master)

JSON is great for sharing data in a human-readable format, but sometimes it lacks in object types support.
ESON does not re-invent the wheel, it just provides a base for you to implement extended JSON objects allowing you to
share data between services, apps and languages as objects.

ESON comes with built-in extensions for date and datetime. You can write your own extensions to manage
custom data.

This is the golang version of ESON. See other languages [here](https://github.com/Billcountry/eson#languages)

## Getting Started
Requires go>=1.17

### Install
Run `go get github.com/alt4dev/eson`

### Usage
Below is a summary of various operations using eson.

#### Encoding:

```go
package main

import (
	"github.com/alt4dev/eson"
	"time"
)

type testStruct struct {
	Name       string    `json:"name"`
	DOB        time.Time `json:"date_of_birth"`
	Roles      []string  `json:"roles"`
	Registered time.Time
}

func main() {
	// Encoding the data
	data := testStruct{
		Name:       "Jane Doe",
		DOB:        time.Now(),
		Roles:      []string{"admin", "client"},
		Registered: time.Now(),
	}

	esonData, err := eson.Encode(data, true)
	if err != nil {
		panic(err)
	}
	println(esonData)
}

```

### Decoding

```go
package main

import (
	"github.com/alt4dev/eson"
	"time"
)

type testStruct struct {
	Name       string    `json:"name"`
	DOB        time.Time `json:"date_of_birth"`
	Roles      []string  `json:"roles"`
	Registered time.Time
}

func main() {
	encodedData := `{"name": "Jane Doe","EsonDate~date_of_birth": 1645804498561,"EsonDatetime~Registered": 1645804498561,"roles": ["admin", "client"]}`

	// Encoding the data
	data := testStruct{}

	// Write encoded data to the pointer provided
	err := eson.Decode(encodedData, &data)

	if err == nil {
		panic(err)
	}
}
```

### Extending ESON
To extend eson you must implement the extension interface:
```go
package main
type Extension interface {
	ShouldEncode(value interface{}) bool
	Encode(value interface{}) interface{}
	Decode(value interface{}) interface{}
}
```

The example below shows ESON's datetime extension
```go
package extension

import (
	"time"
	"github.com/alt4dev/eson/extension"
)

const MilliSecMultiplier int64 = 1000000

type DateTimeExtension struct{}

func (ext DateTimeExtension) ShouldEncode(value interface{}) bool {
	_, ok := value.(time.Time)
	return ok
}

func (ext DateTimeExtension) Encode(value interface{}) interface{} {
	val := value.(time.Time)
	return val.UnixNano() / MilliSecMultiplier
}

func (ext DateTimeExtension) Decode(encodedValue interface{}) interface{} {
	val := encodedValue.(float64)
	return time.Unix(0, int64(val)*MilliSecMultiplier)
}

// You can add the extension for use as follows
func init() {
    extension.AddExtension("EsonDatetime", DateTimeExtension{})
}
```


