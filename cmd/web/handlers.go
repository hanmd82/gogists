package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/hanmd82/gogists/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	gists, err := app.gists.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "home.page.tmpl", &templateData{
		Gists: gists,
	})
}

func (app *application) showGist(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	gist, err := app.gists.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.render(w, r, "show.page.tmpl", &templateData{
		Gist: gist,
	})
}

func (app *application) createGistForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.tmpl", nil)
}

func (app *application) createGist(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest) // 400
		return
	}

	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	expiresInDays := r.PostForm.Get("expires_in_days")

	errors := make(map[string]string)

	if strings.TrimSpace(title) == "" {
		errors["title"] = "This field cannot be blank"
	} else if utf8.RuneCountInString(title) > 100 {
		errors["title"] = "This field is too long (maximum is 100 characters)"
	}

	if strings.TrimSpace(content) == "" {
		errors["content"] = "This field cannot be blank"
	}

	if strings.TrimSpace(expiresInDays) == "" {
		errors["expiresInDays"] = "This field cannot be blank"
	} else if expiresInDays != "365" && expiresInDays != "7" && expiresInDays != "1" {
		errors["expiresInDays"] = "This field is invalid"
	}

	if len(errors) > 0 {
		fmt.Fprint(w, errors)
		return
	}

	id, err := app.gists.Insert(title, content, expiresInDays)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/gist/%d", id), http.StatusSeeOther) // 303
}
