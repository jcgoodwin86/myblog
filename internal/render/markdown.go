package render

import (
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
)

var Markdown = goldmark.New(
	goldmark.WithExtensions(
		meta.Meta,
	),
)
