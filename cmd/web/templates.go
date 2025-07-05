package main

import "snippetbox.alex.net/internal/models"

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}
