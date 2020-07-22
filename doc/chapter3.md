- Use command-line flags to manage configuration settings.
- Leveled Logging
  - Prefix informational messages with "INFO" and output to stdout.
  - Prefix error messages with "ERROR" and output to stderr, along with the relevant file name and line number (for debugging).
    ```bash
    $ go run ./cmd/web >>/tmp/info.log 2>>/tmp/error.log
    ```
- By default, Go’s HTTP server logs errors using the standard logger.
- Custom loggers created by `log.New()` are concurrency-safe.
