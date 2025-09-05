package repository

import (
	"crud-alumni/database"
	"crud-alumni/model"
)

func GetAllAlumni() ([]model.Alumni, error) {
	rows, err := database.DB.Query(`SELECT * FROM alumni ORDER BY id`)
	if err != nil {
		return nil, err
	}

	var list []model.Alumni
	for rows.Next() {
		var a model.Alumni
		rows.Scan(&a.ID, &a.NIM, &a.Nama, &a.Jurusan, &a.Angkatan,
			&a.TahunLulus, &a.Email, &a.NoTelepon, &a.Alamat, &a.CreatedAt, &a.UpdatedAt)
		list = append(list, a)
	}
	return list, nil
}

func GetAlumniByID(id int) (model.Alumni, error) {
	var a model.Alumni
	err := database.DB.QueryRow(`SELECT * FROM alumni WHERE id=$1`, id).
		Scan(&a.ID, &a.NIM, &a.Nama, &a.Jurusan, &a.Angkatan,
			&a.TahunLulus, &a.Email, &a.NoTelepon, &a.Alamat, &a.CreatedAt, &a.UpdatedAt)
	return a, err
}

func InsertAlumni(a model.Alumni) error {
	_, err := database.DB.Exec(`INSERT INTO alumni 
		(nim, nama, jurusan, angkatan, tahun_lulus, email, no_telepon, alamat) 
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`,
		a.NIM, a.Nama, a.Jurusan, a.Angkatan, a.TahunLulus,
		a.Email, a.NoTelepon, a.Alamat)
	return err
}

func UpdateAlumni(id int, a model.Alumni) error {
	_, err := database.DB.Exec(`UPDATE alumni SET nim=$1, nama=$2, jurusan=$3, 
		angkatan=$4, tahun_lulus=$5, email=$6, no_telepon=$7, alamat=$8, updated_at=NOW() 
		WHERE id=$9`,
		a.NIM, a.Nama, a.Jurusan, a.Angkatan, a.TahunLulus,
		a.Email, a.NoTelepon, a.Alamat, id)
	return err
}

func DeleteAlumni(id int) error {
	_, err := database.DB.Exec(`DELETE FROM alumni WHERE id=$1`, id)
	return err
}
