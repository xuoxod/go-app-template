package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/xuoxod/go-app-template/pkg/config"
	"github.com/xuoxod/go-app-template/pkg/handlers"
	"github.com/xuoxod/go-app-template/pkg/render"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the template cache from appConfg

	// Application configuration
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// Routes
	http.HandleFunc("/", handlers.Repo.Index)
	http.HandleFunc("/about", handlers.Repo.About)

	port := os.Getenv("PORT")

	if port == "" {
		port = "4000"
	}

	fmt.Printf("\n\tServer listening on port %v\n\n", port)

	_ = http.ListenAndServe(":"+port, nil)
}
