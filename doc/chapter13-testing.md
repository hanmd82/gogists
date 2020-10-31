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
### Unit Testing

- Unit tests are contained in a normal Go function with the signature `func(*testing.T)`
- To be a valid unit test, the name of this function must begin with the word `Test`
- Use the `t.Errorf()` function to mark a test as failed and log a descriptive message about the failure
- Table-Driven Tests: define test cases in a slice of anonymous structs
- `net/http/httptest` package contains the `httptest.ResponseRecorder` type, which is an implementation of `http.ResponseWriter` which records the response status code, headers and body instead of actually writing them to a HTTP connection
- To unit test HTTP handlers, create a new `httptest.ResponseRecorder` object, pass it to the handler function, and then examine it again after the handler returns
- One way to test HTTP middleware functions is to create a mock HTTP handler to pass to middleware function, then call `ServeHTTP()` and get the response in `ResponseRecorder`
- Run specific tests by using the `-run` flag, passing in a regular expression for matching with test names
- Tests can be configured to run in parallel by calling the `t.Parallel()` function at the start of the test code
- Enable Go's race detector by running with the `-race` flag

---
### End-To-End Testing

- The `httptest.NewTLSServer()` function spins up a `httptest.Server` instance that can accept HTTPS requests for end-to-end testing
  - it accepts a `http.Handler` as the argument, and this handler gets called each time the test server receives a HTTPS request

---
### Mocking Dependencies

- Test the behavior of handlers without needing to setup test database
- Use interfaces, which are satisfied by both the mocks and production database models
