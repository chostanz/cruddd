package controller

import (
	"net/http"
	"projectcrud/models"
	"projectcrud/service"

	"github.com/labstack/echo/v4"
)

func GetAllUser(c echo.Context) error {
	users, _ := service.Getusers(c)

	return c.JSON(http.StatusOK, users)
}

func GetSpecUser(c echo.Context) error {
	user, _ := service.GetSpesificUser(c)
	return c.JSON(http.StatusOK, user)
}

func UpdateUsers(c echo.Context) error {
	reqBody, _ := service.UpdateUser(c)

	if err := c.Validate(&reqBody); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, reqBody)
}

func AddUser(c echo.Context) error {
	userInsert := new(models.Karyawan)
	c.Bind(userInsert)

	err := c.Validate(userInsert)

	if err == nil {
		service.InsertUser(*userInsert)

		return c.JSON(http.StatusCreated, &models.Respons{
			Message: "Successfully add your account",
			Status:  true,
		})
	}

	return echo.NewHTTPError(http.StatusBadRequest, err.Error())

}

func UserDelete(c echo.Context) error {
	user := service.DeleteUser(c)

	return c.JSON(http.StatusOK, user)
}
