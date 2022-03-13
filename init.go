package eson

import "github.com/alt4dev/eson/extension"

var tagName *string
var DefaultExtensions []extension.Extension

// DefaultTag is used as json Key for structs. By default, ESON uses the `json` tag
var DefaultTag string = "json"

func init() {
	SetTagName(DefaultTag)
	DefaultExtensions = make([]extension.Extension, 0)
	AddExtension(extension.EsonDatetime{})
}

// SetTagName allows you to use a custom tag name when encoding or decoding.
// by default the tag 'json' is used
func SetTagName(newTagName string) {
	tagName = &newTagName
}

func AddExtension(extension extension.Extension) {
	DefaultExtensions = append(DefaultExtensions, extension)
}
