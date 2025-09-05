package repository

import (
	"crud-alumni/database"
	"crud-alumni/model"
	"time"
)

func GetAllPekerjaan() ([]model.PekerjaanAlumni, error) {
	rows, err := database.DB.Query(`SELECT * FROM pekerjaan_alumni ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.PekerjaanAlumni
	for rows.Next() {
		var p model.PekerjaanAlumni
		rows.Scan(&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan,
			&p.BidangIndustri, &p.LokasiKerja, &p.GajiRange,
			&p.TanggalMulaiKerja, &p.TanggalSelesaiKerja,
			&p.StatusPekerjaan, &p.DeskripsiPekerjaan,
			&p.CreatedAt, &p.UpdatedAt)
		list = append(list, p)
	}
	return list, nil
}

func GetPekerjaanByID(id int) (model.PekerjaanAlumni, error) {
	var p model.PekerjaanAlumni
	err := database.DB.QueryRow(`SELECT * FROM pekerjaan_alumni WHERE id=$1`, id).
		Scan(&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan,
			&p.BidangIndustri, &p.LokasiKerja, &p.GajiRange,
			&p.TanggalMulaiKerja, &p.TanggalSelesaiKerja,
			&p.StatusPekerjaan, &p.DeskripsiPekerjaan,
			&p.CreatedAt, &p.UpdatedAt)
	return p, err
}

func GetPekerjaanByAlumni(alumniID int) ([]model.PekerjaanAlumni, error) {
	rows, err := database.DB.Query(`SELECT * FROM pekerjaan_alumni WHERE alumni_id=$1 ORDER BY id`, alumniID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.PekerjaanAlumni
	for rows.Next() {
		var p model.PekerjaanAlumni
		rows.Scan(&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan,
			&p.BidangIndustri, &p.LokasiKerja, &p.GajiRange,
			&p.TanggalMulaiKerja, &p.TanggalSelesaiKerja,
			&p.StatusPekerjaan, &p.DeskripsiPekerjaan,
			&p.CreatedAt, &p.UpdatedAt)
		list = append(list, p)
	}
	return list, nil
}

func InsertPekerjaan(p model.PekerjaanAlumni) error {
	now := time.Now()
	_, err := database.DB.Exec(`
        INSERT INTO pekerjaan_alumni 
        (alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri,
         lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja,
         status_pekerjaan, deskripsi_pekerjaan, created_at, updated_at)
        VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)`,
		p.AlumniID, p.NamaPerusahaan, p.PosisiJabatan, p.BidangIndustri,
		p.LokasiKerja, p.GajiRange, p.TanggalMulaiKerja, p.TanggalSelesaiKerja,
		p.StatusPekerjaan, p.DeskripsiPekerjaan, now, now,
	)
	return err
}

func UpdatePekerjaan(id int, p model.PekerjaanAlumni) error {
	_, err := database.DB.Exec(`UPDATE pekerjaan_alumni 
		SET nama_perusahaan=$1, posisi_jabatan=$2, bidang_industri=$3, lokasi_kerja=$4, 
		    gaji_range=$5, tanggal_mulai_kerja=$6, tanggal_selesai_kerja=$7, 
		    status_pekerjaan=$8, deskripsi_pekerjaan=$9, updated_at=NOW() 
		WHERE id=$10`,
		p.NamaPerusahaan, p.PosisiJabatan, p.BidangIndustri, p.LokasiKerja,
		p.GajiRange, p.TanggalMulaiKerja, p.TanggalSelesaiKerja,
		p.StatusPekerjaan, p.DeskripsiPekerjaan, id)
	return err
}

func DeletePekerjaan(id int) error {
	_, err := database.DB.Exec(`DELETE FROM pekerjaan_alumni WHERE id=$1`, id)
	return err
}
