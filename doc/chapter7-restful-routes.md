## Restful Routes

- Go's `servemux` does not support method-based routing or semantic URLs with variables. This can be handled by third-party router packages such as `Pat` or `Gorilla Mux`
- `Pat` matches patterns in the order that they are registered. To ensure that the exact match takes precedence, register the exact match routes before any wildcard routes.
- Because `Pat` matches the `"/"` path exactly, the manual check of `r.URL.Path != "/"` can be removed in `home` handler.
- `Pat` doesn't strip the colon from the named capture key, so the value of `":id"` (instead of `"id"`) needs to be extracted from the query string in `showGist` handler.
