# Stack
net/http → Templ → HTMX → Goldmark + Chroma → Markdown files

# Folder Structure
myblog/
├── main.go
├── go.mod
├── go.sum
│
├── content/
│   └── posts/
│       ├── my-first-post.md
│       └── learning-golang.md
│
├── internal/
│   ├── handler/
│   │   ├── home.go
│   │   ├── post.go
│   │   └── about.go
│   ├── model/
│   │   └── post.go        # Post struct, frontmatter parsing
│   └── render/
│       └── markdown.go    # Goldmark setup, Chroma config
│
├── templates/
│   ├── layout/
│   │   └── base.templ     # Base HTML layout
│   ├── pages/
│   │   ├── home.templ
│   │   ├── post.templ
│   │   └── about.templ
│   └── components/
│       ├── nav.templ
│       └── postcard.templ # Post preview card
│
└── static/
    ├── css/
    │   └── style.css
    └── js/
        └── htmx.min.js
