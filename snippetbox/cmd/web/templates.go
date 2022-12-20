package main

import "snippetbox.jackmitchellfordyce.com/internal/models"

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
