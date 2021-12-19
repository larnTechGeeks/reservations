package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/larnTechGeeks/reservations/pkg/config"
	"github.com/larnTechGeeks/reservations/pkg/handlers"
	"github.com/larnTechGeeks/reservations/pkg/helpers"
)

const defaultPort = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main the netry point of the application
func main() {

	//initialize all configurations and global varibales
	// change to true in production
	app.InProduction = false

	// initializing session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true // set to false to clear the session on browser close
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction // development mode
	// set the session in AppConfig
	app.Session = session
	//build the template Cache
	ts, err := helpers.BuildTC()
	if err != nil {
		log.Println(err)
	}
	app.TCache = ts
	app.DEBUG = true

	// add app config for helper
	helpers.AddAppConfig(&app)

	// create the Repository
	repo := handlers.NewHandler(&app)

	// set the Repo
	handlers.NewRepo(repo)

	log.Println("Server started on PORT ", defaultPort)
	// http.ListenAndServe(defaultPort, nil)

	serve := &http.Server{
		Addr:    defaultPort,
		Handler: routes(&app),
	}

	err = serve.ListenAndServe()

	log.Fatal(err)
}
