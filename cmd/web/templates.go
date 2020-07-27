package main

import "github.com/hanmd82/gogists/pkg/models"

type templateData struct {
	Gist  *models.Gist
	Gists []*models.Gist
}
