package model

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"sort"
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
	Author      string
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

	postData, err := postFromMeta(metaData)
	if err != nil {
		return nil, err
	}
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
	if !ok || title == "" {
		return nil, errors.New("post title is required")
	}
	newPost.Title = title

	var date string
	switch v := meta["date"].(type) {
	case string:
		date = v
	case time.Time:
		date = v.Format("2006-01-02") // ??? how would you convert a time.Time to a string?
	default:
		return nil, errors.New("post date is required")
	}
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, fmt.Errorf("invalid date format: %w", err)
	}
	newPost.Date = parsedDate

	rawTags, ok := meta["tags"].([]interface{})
	if !ok || len(rawTags) == 0 {
		return nil, errors.New("tags are required")
	}
	for _, t := range rawTags {
		tag, ok := t.(string)
		if !ok {
			return nil, errors.New("invalid tag value")
		}
		newPost.Tags = append(newPost.Tags, tag)
	}

	author, ok := meta["author"].(string)
	if !ok || author == "" {
		return nil, errors.New("author is required")
	}
	newPost.Author = author

	description, ok := meta["description"].(string)
	if !ok || description == "" {
		return nil, errors.New("description is required")
	}
	newPost.Description = description

	return &newPost, nil
}

func LoadAllPosts() error {
	entries, err := os.ReadDir(post_location)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		fileName := entry.Name()
		if !entry.IsDir() && strings.HasSuffix(fileName, ".md") {
			slug := strings.TrimSuffix(fileName, ".md")
			_, err := LoadPost(slug)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func GetAllPosts() []*Post {
	cache.mu.RLock()
	defer cache.mu.RUnlock()

	var posts []*Post
	for _, post := range cache.posts {
		posts = append(posts, post)
	}

	// Orders the list from newest to oldest
	sort.Slice(posts, func(i, j int) bool {
		return posts[j].Date.Before(posts[i].Date)
	})

	return posts
}
