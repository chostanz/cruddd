package models

type Karyawan struct {
	Id     int    `json:"id" db:"id"`
	Nama   string `json:"name" db:"name" validate:"required"`
	No_hp  string `json:"phone" db:"phone" validate:"required"`
	Alamat string `json:"address" db:"address" validate:"required"`
}

type Respons struct {
	Message string
	Status  bool
	//Data    []Karyawan
}

func ResponMessage() {
	// respon := Respons
	// 	Message: "Sukses menjalankan query",
	// 	Status:  true,

}
