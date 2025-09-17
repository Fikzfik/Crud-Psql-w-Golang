package repository

import (
	"crud-alumni/app/models"
	"crud-alumni/database"

	"fmt"
)

func GetAllAlumniByFak(fak string, a models.Alumni) ([]models.Alumni, error) {
	rows, err := database.DB.Query(`SELECT * FROM alumni where fakultas=$1`, fak)
	if err != nil {
		return nil, err
	}

	var list []models.Alumni
	for rows.Next() {
		var a models.Alumni
		rows.Scan(&a.ID, &a.NIM, &a.Nama, &a.Jurusan, &a.Angkatan,
			&a.TahunLulus, &a.Email, &a.NoTelepon, &a.Alamat, &a.CreatedAt, &a.UpdatedAt, &a.Fakultas)
		list = append(list, a)
	}
	return list, nil
}

func GetAllAlumni() ([]models.Alumni, error) {
	rows, err := database.DB.Query(`SELECT * FROM alumni ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.Alumni
	for rows.Next() {
		var a models.Alumni
		rows.Scan(
			&a.ID, &a.NIM, &a.Nama, &a.Jurusan, &a.Angkatan,
			&a.TahunLulus, &a.Email, &a.NoTelepon, &a.Alamat, &a.CreatedAt, &a.UpdatedAt, &a.Fakultas)
		list = append(list, a)
	}
	return list, nil
}

func GetAlumniByID(id int) (models.Alumni, error) {
	var a models.Alumni
	err := database.DB.QueryRow(`SELECT * FROM alumni WHERE id=$1`, id).
		Scan(&a.ID, &a.NIM, &a.Nama, &a.Jurusan, &a.Angkatan,
			&a.TahunLulus, &a.Email, &a.NoTelepon, &a.Alamat, &a.CreatedAt, &a.UpdatedAt, &a.Fakultas)
	return a, err
}

func InsertAlumni(a models.Alumni) error {
	_, err := database.DB.Exec(`INSERT INTO alumni 
		(nim, nama, jurusan, angkatan, tahun_lulus, email, no_telepon, alamat) 
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`,
		a.NIM, a.Nama, a.Jurusan, a.Angkatan, a.TahunLulus,
		a.Email, a.NoTelepon, a.Alamat)
	return err
}

func UpdateAlumni(id int, a models.Alumni) error {
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

func GetAlumniWithPagination(search, sortBy, order string, limit, offset int) ([]models.Alumni, error) {
	query := fmt.Sprintf(`
        SELECT id, nim, nama, jurusan, angkatan, tahun_lulus, email, no_telepon, alamat, fakultas, created_at, updated_at
        FROM alumni
        WHERE nama ILIKE $1 OR jurusan ILIKE $1 OR email ILIKE $1 OR fakultas ILIKE $1
        ORDER BY %s %s
        LIMIT $2 OFFSET $3
    `, sortBy, order)

	rows, err := database.DB.Query(query, "%"+search+"%", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.Alumni
	for rows.Next() {
		var a models.Alumni
		rows.Scan(&a.ID, &a.NIM, &a.Nama, &a.Jurusan, &a.Angkatan,
			&a.TahunLulus, &a.Email, &a.NoTelepon, &a.Alamat,
			&a.Fakultas, &a.CreatedAt, &a.UpdatedAt)
		list = append(list, a)
	}
	return list, nil
}

func CountAlumni(search string) (int, error) {
	var total int
	err := database.DB.QueryRow(`SELECT COUNT(*) FROM alumni WHERE nama ILIKE $1 OR jurusan ILIKE $1 OR email ILIKE $1 OR fakultas ILIKE $1`, "%"+search+"%").Scan(&total)
	return total, err
}
