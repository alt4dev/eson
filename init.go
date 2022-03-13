package eson

var tagName string = "json"

// SetTagName allows you to use a custom tag name when encoding or decoding.
// by default the tag 'json' is used
func SetTagName(newTagName string) {
	tagName = newTagName
}
