package main

import (
	"github.com/seun-otosho/go-course/pkg/config"
	"github.com/seun-otosho/go-course/pkg/handlers"
	"github.com/seun-otosho/go-course/pkg/render"
	"log"
	"net/http"
)

const PortNumber = ":8080"

func main() {
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

	log.Printf("Port Number %s", PortNumber)

	srv := &http.Server{
		Addr:    PortNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
