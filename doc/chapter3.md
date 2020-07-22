Command-line Flags
- Use command-line flags to manage configuration settings.

Leveled Logging
- Prefix informational messages with "INFO" and output to stdout.
- Prefix error messages with "ERROR" and output to stderr, along with the relevant file name and line number (for debugging).
    ```bash
    $ go run ./cmd/web >>/tmp/info.log 2>>/tmp/error.log
    ```
- By default, Go’s HTTP server logs errors using the standard logger.
- Custom loggers created by `log.New()` are concurrency-safe.

Dependency Injection
- It is good practice to inject dependencies (e.g. database connection pool, centralized error handlers, template caches) into handlers.
- For applications where handlers are in the same package, a neat way to inject dependencies is to put them into a custom `application` struct, and then define handler functions as methods against `application`.
