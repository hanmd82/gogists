## Testing

Goals:
- Create and run table-driven unit tests and sub-tests in Go
- Test HTTP handlers and middleware
- Perform 'end-to-end' testing of web application routes, middleware and handlers
- Create mocks of database models and use them in unit tests
- Test CSRF-protected HTML form submissions
- Use a test instance of MySQL to perform integration tests
- Calculate and profile test code coverage

---

- Unit tests are contained in a normal Go function with the signature `func(*testing.T)`
- To be a valid unit test, the name of this function must begin with the word `Test`
- Use the `t.Errorf()` function to mark a test as failed and log a descriptive message about the failure

