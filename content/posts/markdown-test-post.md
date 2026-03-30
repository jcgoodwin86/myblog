---
title: "Markdown Kitchen Sink"
date: 2026-03-14
description: "A test post covering every markdown element to validate rendering, syntax highlighting, and styles."
tags: 
  - test 
  - markdown 
  - meta
author: "Joe Goodwin"
---

# Markdown Kitchen Sink

A quick sanity-check post. If everything renders correctly, you're good to ship.

---

## Headings

# H1 — The Big One
## H2 — Section Header
### H3 — Subsection
#### H4 — Sub-subsection
##### H5 — Getting Small
###### H6 — Whisper Level

---

## Paragraphs & Inline Formatting

This is a normal paragraph. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas.

This is **bold text**, this is *italic text*, and this is ***bold and italic***. You can also use ~~strikethrough~~ for things you regret writing. Inline `code snippets` look like this.

Here's a [link to the Go docs](https://pkg.go.dev) and a [relative link to another post](/posts/learning-golang).

---

## Blockquotes

> "The best way to get a project done faster is to start sooner."
> — Someone on the internet, probably

Nested blockquotes:

> Outer quote — setting the scene.
>
> > Inner quote — things are getting philosophical.

---

## Lists

### Unordered

- Goldmark for markdown parsing
- Chroma for syntax highlighting
- Templ for type-safe templates
- HTMX for hypermedia interactions
  - Nested item one
  - Nested item two
    - Doubly nested

### Ordered

1. Install Go
2. Write some handlers
3. Parse some markdown
4. Deploy somewhere cheap
5. Tell nobody and hope they find it

### Task List

- [x] Set up project structure
- [x] Configure Goldmark
- [x] Add Chroma highlighting
- [ ] Write actual blog posts
- [ ] Convince people to read them

---

## Code Blocks

### Go

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, blog!")
	})

	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
```

### JavaScript

```javascript
document.addEventListener("DOMContentLoaded", () => {
  const posts = document.querySelectorAll(".post-card");

  posts.forEach((card) => {
    card.addEventListener("click", (e) => {
      console.log("Navigating to:", e.currentTarget.dataset.slug);
    });
  });
});
```

### HTML

```html
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <title>My Blog</title>
  </head>
  <body>
    <main>
      <h1>Hello, world</h1>
    </main>
  </body>
</html>
```

### CSS

```css
:root {
  --color-bg: #0f0f0f;
  --color-text: #e2e2e2;
  --color-accent: #7c6af7;
}

body {
  background-color: var(--color-bg);
  color: var(--color-text);
  font-family: "Inter", sans-serif;
  line-height: 1.7;
}
```

### Bash

```bash
# Build and run
go build -o ./bin/blog ./cmd/blog
./bin/blog

# Or just
go run main.go
```

### Plain text (no highlighting)

```
This is plain preformatted text.
No highlighting. Good for logs,
config files, or raw output.
```

---

## Tables

| Feature        | Supported | Notes                          |
| -------------- | --------- | ------------------------------ |
| Tables         | ✅        | Via Goldmark extension         |
| Syntax highlighting | ✅   | Via Chroma                     |
| Task lists     | ✅        | Via GFM extension              |
| Footnotes      | ✅        | If extension enabled           |
| Raw HTML       | ⚠️        | Disabled by default in Goldmark |

---

## Horizontal Rules

Three ways to write them, all should render the same:

---

***

___

---

## Images

![A placeholder image](https://placehold.co/800x400?text=Test+Image)

Image with a title attribute (hover text):

![Alt text](https://placehold.co/800x200?text=With+Title "This is the title")

---

## Footnotes

Goldmark supports footnotes if you enable the extension.[^1] They're great for asides that would interrupt the flow.[^2]

[^1]: Enable with `goldmark.WithExtensions(extension.Footnote)`.
[^2]: Like this one — completely optional, but nice to have.

---

## Hard Line Breaks

This line ends with two trailing spaces  
so this should be on a new line without a paragraph gap.

---

## HTML Passthrough

If you've enabled raw HTML in Goldmark, this should render as a styled element:

<mark>This text should be highlighted</mark> — if raw HTML is allowed.

<details>
  <summary>Click to expand</summary>
  Hidden content revealed. Useful for spoilers or long asides.
</details>

---

## Edge Cases

Empty-ish content below a heading — make sure spacing holds up.

###

A heading with no text (bad practice, but shouldn't explode).

Very long unbroken word to test overflow: `superlongidentifierthatmightcauseoverflowissuesifyouarenotcareful_v2_final_FINAL`.

Unicode: 日本語テスト • Ünïcödé • العربية • 中文 • Ελληνικά

Emoji: 🚀 🦫 🧪 ✅ 🎯

---

*If everything above looks right — headings, code blocks with highlighting, tables, lists, blockquotes, images, and footnotes — your markdown pipeline is solid.*
