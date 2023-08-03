package routes

import (
	"projectcrud/controller"

	"github.com/labstack/echo/v4"
	//"github.com/labstack/echo/v4"
)

func Route() *echo.Echo {
	e := echo.New()
	e.GET("/users", controller.GetAllUser)
	e.GET("/users/:id", controller.GetSpecUser)
	e.PUT("/users/update/:id", controller.UpdateUsers)
	e.POST("/users", controller.AddUser)
	e.DELETE("/users/delete/:id", controller.UserDelete)
	//r.Validator = &utils.CustomValidator{Validator: validator.New()}

	return e

}
