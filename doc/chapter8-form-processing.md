## Form Processing

Techniques
- Parse and access form data sent in a `POST` request
  - use `r.ParseForm()` method to parse the request body. This checks that the request body is well-formed, then stores the form data in the request's `r.PostForm` map
  - extract form data from `r.PostForm` by using the `r.PostForm.Get()` method
- Perform common validation checks on the form data
- Alert the user to validation failures and re-populate form fields with previously submitted data
- Scale-up validation and keep handlers clean by creating a form helper in a separate reusable package

---

Notes
- The `r.PostForm` map is populated only for `POST`, `PATCH` and `PUT` requests, and contains the form data from the request body
- Alternatively, use the `r.Form` map which is populated for all requests (irrespective of HTTP method), and contains the form data from any request body and any query string parameters
- `net/http` package also provides the methods `r.FormValue()` and `r.PostFormValue()`
- For fields with multiple values, loop over the underlying map to access them
    ```go
    for i, item := range r.PostForm["items"] {
        fmt.Fprintf(w, "%d: Item %s\n", i, item)
    }
    ```
- use `FormErrors` to hold any validation errors and `FormData` to hold any previously submitted data
- when there are any validation errors, re-display the `create.page.tmpl` template, passing in the map of errors in the `FormErrors` field of the template data, and passing in the previously submitted data in the `FormData` field
