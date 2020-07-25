## Displaying Dynamic Data

- Go’s `html/template` package allows passing in one — and only one — item of dynamic data when rendering a template. But in a real-world application there are often multiple pieces of dynamic data that to display in the same page.
- Dynamic data can be wrapped in a struct which acts like a single ‘holding structure’ for the data.
