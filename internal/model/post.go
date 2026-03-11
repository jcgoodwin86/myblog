package model

import (
	"bytes"
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"sync"
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

type postCache struct {
	mu    sync.RWMutex
	posts map[string]*Post
}

var cache = &postCache{posts: make(map[string]*Post)}
var post_location string = "content/posts"

func LoadPost(slug string) (*Post, error) {
	// Check if post exists
	cache.mu.RLock()
	post, ok := cache.posts[slug]
	cache.mu.RUnlock()
	if ok {
		return post, nil
	}

	path := filepath.Join(post_location, slug+".md")
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

	cache.mu.Lock()
	cache.posts[postData.Slug] = postData
	cache.mu.Unlock()

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

func LoadAllPosts() error {
	entries, err := os.ReadDir(post_location)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		fileName := entry.Name()
		slug := strings.TrimSuffix(fileName, ".md")
		_, err := LoadPost(slug)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetAllPosts() []*Post {
	cache.mu.RLock()
	defer cache.mu.RUnlock()
	// now safely read from cache.posts
	var posts []*Post
	for _, post := range cache.posts {
		posts = append(posts, post)
	}

	return posts
}
