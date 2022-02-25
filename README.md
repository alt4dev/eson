# Extended JSON(ESON)

This is an implementation of https://github.com/Billcountry/eson in golang.

![Python Tests](https://github.com/alt4dev/eson/workflows/Tests/badge.svg?branch=master)

JSON is great for sharing data in a human readable format but sometimes it lacks in object types support.
ESON does not re-invent the wheel, it just provides a base for you to implement extended JSON objects allowing you to
share data between services, apps and languages as objects.

ESON comes with built in extensions for date and datetime. You can write your own extensions to manage
custom data.

This is the golang version of ESON. See other languages [here](https://github.com/Billcountry/eson#languages)

## Getting Started

### Install
Run `pip install eson`

### Usage
Below is a summary of various operations using eson.

#### Encoding:
```go
import (
    "eson"
    "time"
)

type testStruct struct {
    Name       string    `json:"name"`
    DOB        time.Time `json:"date_of_birth"`
    Roles      []string  `json:"roles"`
    Registered time.Time
}

// Encoding the data
data := testStruct{
    Name:       "Jane Doe",
    DOB:        time.Now(),
    Roles:      []string{"admin", "client"},
    Registered: time.Now(),
}

eson_data = eson.Encode(user, true)

// Sample output
"""
{
    "name": "Jane Doe",
    "EsonDate~date_of_birth": 1645804498561,
    "EsonDatetime~Registered": 1645804498561,
	"roles": ["admin", "client"]
}
"""
```

