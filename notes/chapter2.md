```go
# initialize Go module
go mod init github.com/hanmd82/gogists
```

Fixed URL Path vs Subtree Patterns
- Fixed path patterns are only matched (and corresponding handler called) when the request URL path exactly matches the fixed path, e.g. `/gist/:gist_id`, `/gists/create`.
- Subtree path patterns (ending in a trailing slash, e.g. `/static/`) are matched (and corresponding handler called) whenever the start of a request URL path matches the subtree path, e.g. acts like `/static/**`
