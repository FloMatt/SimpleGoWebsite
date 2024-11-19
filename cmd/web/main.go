package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/FloMatt/SimpleGoWebsite/pkg/config"
	"github.com/FloMatt/SimpleGoWebsite/pkg/handlers"
	"github.com/FloMatt/SimpleGoWebsite/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	// Initialize the template cache
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")

	}
	// Initialize the template
	app.TemplateCache = tc
	app.UseCache = false

	// Initialize the logger
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers((repo))

	render.NewTemplates(&app)

	fmt.Printf("Starting application port %s\n", portNumber)
	//_ = http.ListenAndServe(portNumber, nil)

	// Use the routes function to set up routes and handlers
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	// Start the server and handle any errors that occur
	err = srv.ListenAndServe()
	log.Fatal(err)
}
