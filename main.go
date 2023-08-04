package main

import (
	_ "fmt"

	"projectcrud/routes"
	"projectcrud/utils"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// e := echo.New()
	// e.GET("/users", routes.Getusers)
	// e.GET("/users/:id", routes.GetSpesificUser)
	// e.Use(middleware.CORS())
	e := routes.Route()

	e.Use(middleware.CORS())

	// Mengatur validator kustom
	// e.Validator = &utils.CustomValidator{validator: validator.New()}
	customValidator := &utils.CustomValidator{Validator: validator.New()}
	e.Validator = customValidator

	// Jalankan server di port tertentu (misalnya port 8080)
	e.Logger.Fatal(e.Start(":8080"))

}
