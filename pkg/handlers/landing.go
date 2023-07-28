package handlers

import (
	"net/http"

	"github.com/xuoxod/go-app-template/pkg/render"
)

func Index(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
