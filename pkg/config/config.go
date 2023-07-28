package config

import (
	"log"
	"text/template"
)

// App Configuration
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
