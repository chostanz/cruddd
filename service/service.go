package service

import (
	"log"
	"projectcrud/database"
	"projectcrud/models"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

var db *sqlx.DB = database.Connection()

func Getusers(c echo.Context) ([]models.Karyawan, error) {
	var user []models.Karyawan
	//menjalankan query di simpan di var rows
	//parameter dan argumen tu sama, ada di dalam function
	rows, err := db.Queryx("select * from users")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		//membuat var bru dari struct Karyawan yg nilai awalnya kosong
		place := models.Karyawan{}
		rows.StructScan(&place)
		user = append(user, place) //menambah elemen baru menggunakan append ke users
	}
	return user, nil
}

func GetSpesificUser(c echo.Context) (models.Karyawan, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return models.Karyawan{}, err
	}

	user := models.Karyawan{}
	db.Get(&user, "SELECT * FROM users WHERE id = $1", id)

	return user, nil
}

func UpdateUser(c echo.Context) (models.Karyawan, error) {
	reqBody := models.Karyawan{}

	c.Bind(&reqBody)

	//menggunakan parameter untuk menghapus data dgn nilai dari parameter :id, bertipe string
	id, _ := strconv.Atoi(c.Param("id")) // Konversi id dari string ke int

	//mengembalikan nilai id ke reqBody.Id (sturct ReqBody Id)
	reqBody.Id = id

	// query update value pakek dari reqbody
	_, err := db.NamedExec("update users SET name= :name, phone= :phone, address= :address WHERE id= :id", reqBody)
	if err != nil {
		return models.Karyawan{}, err // Mengembalikan error jika terjadi kesalahan dalam query atau eksekusi.
	}

	return reqBody, nil
}

func InsertUser(userInsert models.Karyawan) models.Karyawan {
	// reqBody := models.Karyawan{}

	// query insert pakek insert into tapi value pakek dari reqbody
	_, err := db.NamedExec("insert into users(name, phone, address) values (:name, :phone, :address)", userInsert)

	if err != nil {
		log.Fatal(err)
		return models.Karyawan{}
	}
	return userInsert
}

func DeleteUser(c echo.Context) bool {
	id, _ := strconv.Atoi(c.Param("id"))

	// Lakukan DELETE pada database menggunakan ID yang bertipe int
	// hapus data dengan id tertentu

	db.Exec("DELETE FROM users WHERE id = $1", id)
	return true
}
