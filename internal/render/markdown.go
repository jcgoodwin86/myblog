package render

import (
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
)

var Markdown = goldmark.New(
	goldmark.WithExtensions(
		meta.Meta,
		extension.Table,
		extension.Footnote,
		extension.Strikethrough,
	),
)
