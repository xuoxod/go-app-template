package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/joho/godotenv"
	"github.com/xuoxod/go-app-template/pkg/config"
	"github.com/xuoxod/go-app-template/pkg/handlers"
	"github.com/xuoxod/go-app-template/pkg/render"
)

// Application configuration
var app config.AppConfig

var session *scs.SessionManager

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the template cache from appConfg

	// Application mode
	app.InProduction = false

	// Session middleware
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	// Set the app level session
	app.Session = session

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = true

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	// Routes
	// http.HandleFunc("/", handlers.Repo.Index)
	// http.HandleFunc("/about", handlers.Repo.About)

	port := os.Getenv("PORT")

	if port == "" {
		port = "4000"
	}

	fmt.Printf("\n\tServer listening on port %v\n\n", port)

	// _ = http.ListenAndServe(":"+port, nil)

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()

	log.Fatal(err)
}
