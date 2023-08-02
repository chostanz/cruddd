package routes

import (
	"crudkaryawan/projectcrud/database"
	"crudkaryawan/projectcrud/models"
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	//"github.com/labstack/echo/v4"
)

var db *sqlx.DB = database.Connection()

func Getusers(c echo.Context) error {
	var user []models.Karyawan
	//menjalankan query di simpan di var rows
	//parameter dan argumen tu sama, ada di dalam function
	rows, _ := db.Queryx("select * from users")

	for rows.Next() {
		//membuat var bru dari struct Karyawan yg nilai awalnya kosong
		place := models.Karyawan{}
		rows.StructScan(&place)
		user = append(user, place) //menambah elemen baru menggunakan append ke users
	}
	return c.JSON(http.StatusOK, user)
}

func GetSpesificUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user := models.Karyawan{}
	db.Get(&user, "SELECT * FROM users WHERE id = $1", id)

	return c.JSON(http.StatusOK, user)
}
