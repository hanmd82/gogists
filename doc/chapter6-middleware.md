## Middleware

Standard pattern for creating middleware:

Accept the next handler in a chain as a parameter and return a handler which executes some logic and then calls the `next` handler.

```go
func myMiddleware(next http.Handler) http.Handler {
    fn := func(w http.ResponseWriter, r *http.Request) {
        // TODO: Execute middleware logic...
        next.ServeHTTP(w, r)
    }

    return http.HandlerFunc(fn)
}
```

Rewriting to use anonymous function:
```go
func myMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // TODO: Execute middleware logic...
        next.ServeHTTP(w, r)
    })
}
```

Position of the middleware function in the chain of handlers affects the behavior of the application:
- before servemux: middleware will be executed on every request that the application receives. E.g. useful for security, logging middleware.
    ```
    myMiddleware -> servemux -> application handler
    ```
- after servemux (by wrapping a specific application handler): middleware will be executed only for specific routes. E.g. useful for authorization middleware.
    ```
    servemux -> myMiddleware -> application handler
    ```

### Flow of Control
In any middleware handler, code which comes before `next.ServeHTTP()` will be executed on the way down the chain, and any code after `next.ServeHTTP()` - or in a deferred function - will be executed on the way back up.

### Early Returns

If `return` is called in middleware function before calling `next.ServeHTTP()`, then the chain will stop being executed and control will flow back upstream.

A common use-case is in authentication middleware.
