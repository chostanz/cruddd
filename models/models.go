package models

import "github.com/go-playground/validator/v10"

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

type CustomValidator struct {
	validator *validator.Validate
}
