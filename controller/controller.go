package controller

import (
	"net/http"
	"projectcrud/models"
	"projectcrud/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllUser(c echo.Context) error {
	users, err := service.Getusers()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}

func GetSpecUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var user models.Karyawan

	user, err := service.GetSpesificUser(user, id)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

func UpdateUsers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var reqBody models.Karyawan
	if err := c.Bind(&reqBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if errValidate := c.Validate(&reqBody); errValidate != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Data tidak boleh kosong")
	}

	_, errorUpdate := service.UpdateUser(reqBody, id)
	if errorUpdate != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Nomor telepon tidak boleh sama!")
	}

	return c.JSON(http.StatusOK, &models.Respons{
		Message: "Data sukses diupdate",
		Status:  true,
	})
}

func AddUser(c echo.Context) error {
	userInsert := new(models.Karyawan)
	c.Bind(userInsert)

	err := c.Validate(userInsert)

	if err == nil {
		_, insertErr := service.InsertUser(*userInsert)
		if insertErr != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Nomor telepon tidak boleh sama!")
		}
		return c.JSON(http.StatusCreated, &models.Respons{
			Message: "Sukses menambah user",
			Status:  true,
		})
	}

	return echo.NewHTTPError(http.StatusBadRequest, err.Error())

}

func UserDelete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	var user models.Karyawan

	_, err = service.DeleteUser(user, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())

	}

	return c.JSON(http.StatusOK, &models.Respons{
		Message: "Sukses menghapus data",
		Status:  true,
	})

}
