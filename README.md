## GoGists

A web application written in Go, for sharing text snippets or gists.

Using mostly Go's standard library and some third-party packages - without relying on a web framework.

Some topics covered:
- Routing
- Templating
- Working with Postgres. Database model design
- Mocking Dependencies
- Authentication/Authorization
- Using HTTPS
- Go's testing package

---

### Usage
```bash
go get github.com/hanmd82/gogists
go run ./cmd/web/
```

### Testing
```bash
go test -v ./...
```

---

### Screenshots

Homepage

![gogists_home](https://user-images.githubusercontent.com/761959/97804129-cddee900-1c88-11eb-9d17-bda46d14eb99.png)

Show Gist

![gogists_showgist](https://user-images.githubusercontent.com/761959/97804131-ce777f80-1c88-11eb-872c-fa9c14d8affd.png)

Create Gist

![gogists_creategist](https://user-images.githubusercontent.com/761959/97804128-cb7c8f00-1c88-11eb-8329-49b9fab627d5.png)
