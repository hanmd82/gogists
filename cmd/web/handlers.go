package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

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
	w.Write([]byte("Create a new gist..."))
}

func (app *application) createGist(w http.ResponseWriter, r *http.Request) {
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expiresInDays := "7"

	id, err := app.gists.Insert(title, content, expiresInDays)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/gist/%d", id), http.StatusSeeOther) // 303
}
