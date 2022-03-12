package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/guillermogrillo/comments-api/internal/transport/http"
)

type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting up the app")

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
