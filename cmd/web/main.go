package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/FloMatt/SimpleGoWebsite/pkg/config"
	"github.com/FloMatt/SimpleGoWebsite/pkg/handlers"
	"github.com/FloMatt/SimpleGoWebsite/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	//change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = time.Hour * 24
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

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
