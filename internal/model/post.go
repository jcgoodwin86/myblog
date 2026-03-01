package model

import (
	"bytes"
	"html/template"
	"os"
	"path/filepath"
	"time"

	"github.com/jcgoodwin/myblog/internal/render"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

type Post struct {
	Slug        string
	Title       string
	Date        time.Time
	Description string
	Tags        []string
	Content     template.HTML // rendered HTML from markdown
}

func LoadPost(slug string) (*Post, error) {
	path := filepath.Join("content/posts", slug+".md")
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	// parse frontmatter, render markdown body...
	var buf bytes.Buffer
	context := parser.NewContext()
	if err := render.Markdown.Convert(data, &buf, parser.WithContext(context)); err != nil {
		return nil, err
	}

	metaData := meta.Get(context)

	postData, _ := postFromMeta(metaData)
	postData.Slug = slug
	postData.Content = template.HTML(buf.String())

	return postData, nil
}

func postFromMeta(meta map[string]interface{}) (*Post, error) {
	newPost := Post{}

	title, ok := meta["title"].(string)
	if ok {
		newPost.Title = title
	}

	description, ok := meta["description"].(string)
	if ok {
		newPost.Description = description
	}

	date, ok := meta["date"].(string)
	if ok {
		parsedDate, err := time.Parse("2006-01-02", date)
		if err != nil {
			return nil, err
		}
		newPost.Date = parsedDate
	}

	tags, ok := meta["tags"].([]string)
	if ok {
		newPost.Tags = tags
	}

	return &newPost, nil
}
