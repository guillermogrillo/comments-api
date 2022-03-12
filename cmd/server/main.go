package main

import (
	"fmt"
	"net/http"

	"github.com/guillermogrillo/comments-api/internal/database"

	transportHTTP "github.com/guillermogrillo/comments-api/internal/transport/http"
)

type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting up the app")
	var err error
	_, err = database.NewDatabase()
	if err != nil {
		return err
	}

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		return err
	}
	return nil
}

func main() {
	fmt.Println("Comments service")

	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting app")
		fmt.Println(err)
	}
}
