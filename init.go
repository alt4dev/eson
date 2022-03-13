package eson

var tagName *string

// DefaultTag is used as json Key for structs. By default, ESON uses the `json` tag
var DefaultTag string = "json"

func init() {
	tagName = &DefaultTag
}

// SetTagName allows you to use a custom tag name when encoding or decoding.
// by default the tag 'json' is used
func SetTagName(newTagName string) {
	tagName = &newTagName
}
