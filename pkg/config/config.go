package config

import "html/template"

//AppConfig holds the application wide configuration
type AppConfig struct {
	TemplateCache map[string]*template.Template
	UseCache      bool
}
