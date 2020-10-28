## Request Context

- Every `http.Request` has a `context.Context` object embedded in it, which can be used to store information during the lifetime of the request
- A common use-case for this is to pass information between middleware and other handlers
- This can be used to check if a user is authenticated-and-active once in some middleware, and if they are, then make this information available to all other middleware and handlers
- Basic code for adding information to a request's context - create a new copy of the `http.Request` object, decorated with the new context
    ```go
    // given that r is a *http.Request
    ctx := r.Context()
    ctx = context.WithValue(ctx, 'isAuthenticated', true)
    r = r.WithContext(ctx)
    ```
- Request context values are stored with the type `interface{}`. After retrieving them from the context, assert them to their original type before using them
    ```go
    isAuthenticated, ok := r.Context().Value("isAuthenticated").(bool)
    if !ok {
        return errors.New("could not convert value to bool")
    }
    ```
- Good practice to create custom types which can be used as context keys
    ```go
    type contextKey string
    const contextKeyIsAuthenticated = contextKey("isAuthenticated")
    ...

    ctx = context.WithValue(r.Context(), contextKeyIsAuthenticated, true)
    r = r.WithContext(ctx)
    ...

    isAuthenticated, ok := r.Context().Value(contextKeyIsAuthenticated).(bool)
    if !ok {
        return errors.New("could not convert value to bool")
    }
    ```
