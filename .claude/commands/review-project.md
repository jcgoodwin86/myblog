## Files to Read
1. All .go files
2. All .templ files
3. go.mod / go.sum
4. Sample .md content files (to understand expected input)
5. Any config or layout files

## Additional Review Areas

### Templ
- Are components structured with clear separation (layout, partials, pages)?
- Is data passed cleanly from handlers → Templ components?
- Are components reused well or is there duplication?
- Error handling on templ.Execute / rendering calls

### Goldmark Pipeline
- Where is Goldmark configured — is it centralised or scattered?
- Are extensions (syntax highlighting, footnotes, etc.) configured intentionally?
- Is the Markdown → HTML conversion happening at request time or cached?
  (For a blog, parsing on every request is a red flag)
- Is the rendered HTML safely handled to avoid XSS?

### net/http → Templ → Goldmark data flow
- Are handlers thin? (fetch data → render template — nothing more)
- Is Markdown content loaded from disk on each request or preloaded?
- Is there a clear content/service layer separating HTTP concerns from
  Goldmark rendering logic?

### Markdown Content Files
- Is there a consistent frontmatter format (title, date, slug, etc.)?
- How is frontmatter parsed — is it robust?
- Is the content directory structured sensibly for future growth?
