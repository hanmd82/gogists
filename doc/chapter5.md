## Displaying Dynamic Data

- Go’s `html/template` package allows passing in one — and only one — item of dynamic data when rendering a template. But in a real-world application there are often multiple pieces of dynamic data that to display in the same page.
- Dynamic data can be wrapped in a struct which acts like a single ‘holding structure’ for the data.
- The `html/template` package automatically escapes any data that is yielded between `{{ }}` tags. This behavior helps in avoiding cross-site scripting (XSS) attacks. Also, `html/template` strips out any HTML comments.
- For nested templates, the dot `.` needs to be explicitly passed or pipelined to the template being invoked.

---
### Template Actions and Functions

Some Template Actions:
- `{{define}}`
- `{{template}}`
- `{{block}}`
- `{{if}}`
- `{{with}}` (changes the context of `.`)
- `{{range}}` (changes the context of `.`)

Some Template Functions - see [here](https://golang.org/pkg/text/template/#hdr-Functions):
- `{{eq}}`, `{{ne}}`
- `{{or}}`, `{{and}}`, `{{not}}`
- `{{index}}`
- `{{printf}}`
- `{{len}}`
- `{{$bar := len .Foo}}` (declaration of template variable, and assignment operator)
