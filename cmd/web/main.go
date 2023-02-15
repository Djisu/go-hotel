package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Djisu/go-hotel/internal/config"
	"github.com/Djisu/go-hotel/internal/handlers"
	"github.com/Djisu/go-hotel/internal/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

// Declare a variable for the config.AppConfig file.
var app config.AppConfig
var session *scs.SessionManager

// main is the main application function
func main() {

	//change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	//
	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// start the server
	fmt.Printf("Starting the application on port %s", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
