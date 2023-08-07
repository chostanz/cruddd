package service

import (
	"projectcrud/database"
	"projectcrud/models"
	"strconv"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB = database.Connection()

func Getusers() ([]models.Karyawan, error) {
	userGet := []models.Karyawan{}

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
		userGet = append(userGet, place) //menambah elemen baru menggunakan append ke users
	}
	return userGet, nil
}

func GetSpesificUser(user models.Karyawan, id int) (models.Karyawan, error) {
	idStr := strconv.Itoa(id)

	err := db.Get(&user, "SELECT * FROM users WHERE id = $1", idStr)
	if err != nil {
		return models.Karyawan{}, err
	}
	return user, nil
}

func UpdateUser(reqBody models.Karyawan, id int) (models.Karyawan, error) {

	//menggunakan parameter untuk menghapus data dgn nilai dari parameter :id, bertipe string
	//mengembalikan nilai id ke reqBody.Id (sturct ReqBody Id)

	// updateuser.Id = id

	// // query update value pakek dari reqbody
	// _, err := db.NamedExec("update users SET name= :name, phone= :phone, address= :address WHERE id= :id", updateuser)
	// if err != nil {
	// 	return models.Karyawan{}, err // Mengembalikan error jika terjadi kesalahan dalam query atau eksekusi.
	// }

	// return updateuser, nil
	idStr := strconv.Itoa(id)

	// Menggunakan id dari parameter untuk meng-update data berdasarkan id tertentu
	_, err := db.NamedExec("UPDATE users SET name = :name, phone = :phone, address = :address WHERE id = :id", map[string]interface{}{
		"name":    reqBody.Nama,
		"phone":   reqBody.No_hp,
		"address": reqBody.Alamat,
		"id":      idStr,
	})
	if err != nil {
		return models.Karyawan{}, err // Mengembalikan error jika terjadi kesalahan dalam query atau eksekusi.
	}

	return reqBody, nil
}

func InsertUser(userInsert models.Karyawan) (models.Karyawan, error) {
	// reqBody := models.Karyawan{}

	// query insert pakek insert into tapi value pakek dari reqbody
	_, err := db.NamedExec("insert into users(name, phone, address) values (:name, :phone, :address)", userInsert)

	if err != nil {
		return models.Karyawan{}, err
	}
	return userInsert, nil
}

func DeleteUser(user models.Karyawan, id int) (models.Karyawan, error) {

	idStr := strconv.Itoa(id)

	// Lakukan DELETE pada database menggunakan ID yang bertipe int
	// hapus data dengan id tertentu

	_, err := db.Exec("DELETE FROM users WHERE id = $1", idStr)
	if err != nil {
		return user, err
	}
	return user, nil
}
