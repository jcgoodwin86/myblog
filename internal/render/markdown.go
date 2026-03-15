package render

import (
	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
)

var Markdown = goldmark.New(
	goldmark.WithExtensions(
		meta.Meta,
		extension.Table,
		extension.Footnote,
		extension.Strikethrough,
		highlighting.NewHighlighting(
			// highlighting.WithStyle("catppuccin-latte"),
			highlighting.WithFormatOptions(
				chromahtml.WithLineNumbers(true),
				chromahtml.WithClasses(false),
			),
		),
	),
)
