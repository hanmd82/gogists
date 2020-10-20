## Stateful HTTP

Goals
- Share data (or state) between HTTP requests for the same user by implementing user sessions
- Implement sessions in Go using session manager packages
- Customize session behavior (including timeouts and cookie settings)
- Use sessions to safely and securely share data between requests for a particular user

---

Setting Up the Session Manager
- Establish a session manager and make it available to handlers via the `application` struct
- Define a 32-byte secret key for encrypting and authenticating session cookies
- Wrap application routes with the middleware provided by the `Session.Enable()` method. This middleware loads and saves session data to and from the session cookie with every HTTP request and response
  - Create a new middleware chain containing the middleware specific to the dynamic application routes

Working with Session Data
- use `*Session.Put()` to add data to session, with a specified key
- use `*Session.Get()` to retrieve data from the session, with the given key
- use `*Session.PopString()` to remove data from session after reading
