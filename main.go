package main

import (
	_ "fmt"

	"crudkaryawan/projectcrud/routes"

	_ "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.GET("/users", routes.Getusers)
	e.GET("/users/:id", routes.GetSpesificUser)
	e.Use(middleware.CORS())

	// Mengatur validator kustom
	//e.Validator = &CustomValidator{validator: validator.New()}

	// Jalankan server di port tertentu (misalnya port 8080)
	e.Logger.Fatal(e.Start(":8080"))

}
