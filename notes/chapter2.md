Initialize Go module
```go
go mod init github.com/hanmd82/gogists
```

Fixed URL Path vs Subtree Patterns
- Fixed path patterns are only matched (and corresponding handler called) when the request URL path exactly matches the fixed path, e.g. `/gist/:gist_id`, `/gists/create`.
- Subtree path patterns (ending in a trailing slash, e.g. `/static/`) are matched (and corresponding handler called) whenever the start of a request URL path matches the subtree path, e.g. acts like `/static/**`.

DefaultServeMux
- `net/http` provides a default servemux as a global variable:
  ```
  var DefaultServeMux = NewServeMux()
  ```

Customizing HTTP Headers
- `w.WriteHeader()` can only be called once per response, and cannot be changed after status code has been written.
- The first call to `w.Write()` will automatically send a `200 OK` status code. To send a non-`200` status code, call `w.WriteHeader()` before any call to `w.Write()`.
- Use `w.Header().Set()` method to add new headers to the response header map.
- Changing the response header map after a call to `w.WriteHeader()` or `w.Write()` will have no effect.
