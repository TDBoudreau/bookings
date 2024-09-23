package main

import (
	"database/sql"
	"encoding/gob"
	"fmt"
	"github.com/TDBoudreau/bookings/internal/driver"
	"github.com/TDBoudreau/bookings/internal/helpers"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/TDBoudreau/bookings/internal/config"
	"github.com/TDBoudreau/bookings/internal/handlers"
	"github.com/TDBoudreau/bookings/internal/models"
	"github.com/TDBoudreau/bookings/internal/render"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":3020"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer func(SQL *sql.DB) {
		if err := SQL.Close(); err != nil {
			log.Fatal(err)
		}
	}(db.SQL)
	defer close(app.MailChan)

	fmt.Printf("Starting application on port %s\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() (*driver.DB, error) {
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Reservation{})
	gob.Register(models.Restriction{})

	mailChan := make(chan models.MailData)
	app.MailChan = mailChan

	app.InProduction = false // change this to true when in production
	app.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	// connect to database
	log.Println("Connecting to database")
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=bookings user=postgres password=password")
	if err != nil {
		log.Fatal("Cannot connect to database! Dying...")
	}

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return nil, err
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	render.NewRenderer(&app)
	helpers.NewHelpers(&app)

	return db, nil
}
