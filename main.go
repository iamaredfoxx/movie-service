package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	echo "github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal(err)
	}

	// Load Movie Data
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Small Service running smooth :)")
	})
	e.Logger.Fatal(e.Start(os.Getenv("APP_EXPOSE")))
}
