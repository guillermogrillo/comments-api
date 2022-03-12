package main

import (
	"net/http"

	"github.com/guillermogrillo/comments-api/internal/comment"

	"github.com/guillermogrillo/comments-api/internal/database"

	transportHTTP "github.com/guillermogrillo/comments-api/internal/transport/http"

	log "github.com/sirupsen/logrus"
)

type App struct {
	Name    string
	Version string
}

func (app *App) Run() error {
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(
		log.Fields{
			"AppName":    app.Name,
			"AppVersion": app.Version,
		}).Info("Starting application")
	var err error
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}

	err = database.MigrateDB(db)
	if err != nil {
		log.Error("Error running the migration")
		return err
	}

	commentService := comment.NewService(db)

	handler := transportHTTP.NewHandler(commentService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		return err
	}
	return nil
}

func main() {
	app := App{
		Name:    "comments-api",
		Version: "1.0.0",
	}
	if err := app.Run(); err != nil {
		log.Error("Error starting app")
		log.Fatal(err)
	}
}
