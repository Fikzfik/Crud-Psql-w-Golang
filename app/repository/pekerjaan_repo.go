package repository

import (
	"crud-alumni/app/models"
	"crud-alumni/database"
	"fmt"
	"time"
)

func GetAllPekerjaan() ([]models.PekerjaanAlumni, error) {
	rows, err := database.DB.Query(`SELECT * FROM pekerjaan_alumni ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.PekerjaanAlumni
	for rows.Next() {
		var p models.PekerjaanAlumni
		rows.Scan(&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan,
			&p.BidangIndustri, &p.LokasiKerja, &p.GajiRange,
			&p.TanggalMulaiKerja, &p.TanggalSelesaiKerja,
			&p.StatusPekerjaan, &p.DeskripsiPekerjaan,
			&p.CreatedAt, &p.UpdatedAt)
		list = append(list, p)
	}
	return list, nil
}

func GetPekerjaanByID(id int) (models.PekerjaanAlumni, error) {
	var p models.PekerjaanAlumni
	err := database.DB.QueryRow(`SELECT * FROM pekerjaan_alumni WHERE id=$1`, id).
		Scan(&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan,
			&p.BidangIndustri, &p.LokasiKerja, &p.GajiRange,
			&p.TanggalMulaiKerja, &p.TanggalSelesaiKerja,
			&p.StatusPekerjaan, &p.DeskripsiPekerjaan,
			&p.CreatedAt, &p.UpdatedAt)
	return p, err
}

func GetPekerjaanByAlumni(alumniID int) ([]models.PekerjaanAlumni, error) {
	rows, err := database.DB.Query(`SELECT * FROM pekerjaan_alumni WHERE alumni_id=$1 ORDER BY id`, alumniID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.PekerjaanAlumni
	for rows.Next() {
		var p models.PekerjaanAlumni
		rows.Scan(&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan,
			&p.BidangIndustri, &p.LokasiKerja, &p.GajiRange,
			&p.TanggalMulaiKerja, &p.TanggalSelesaiKerja,
			&p.StatusPekerjaan, &p.DeskripsiPekerjaan,
			&p.CreatedAt, &p.UpdatedAt)
		list = append(list, p)
	}
	return list, nil
}

func InsertPekerjaan(p models.PekerjaanAlumni) error {
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

func UpdatePekerjaan(id int, p models.PekerjaanAlumni) error {
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

func GetPekerjaanWithPagination(search, sortBy, order string, limit, offset int) ([]models.PekerjaanAlumni, error) {
	query := fmt.Sprintf(`
		SELECT id, alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, 
		       lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja, 
		       status_pekerjaan, deskripsi_pekerjaan, created_at, updated_at
		FROM pekerjaan_alumni
		WHERE nama_perusahaan ILIKE $1 
		   OR posisi_jabatan ILIKE $1 
		   OR bidang_industri ILIKE $1 
		   OR lokasi_kerja ILIKE $1
		ORDER BY %s %s
		LIMIT $2 OFFSET $3
	`, sortBy, order)

	rows, err := database.DB.Query(query, "%"+search+"%", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.PekerjaanAlumni
	for rows.Next() {
		var p models.PekerjaanAlumni
		rows.Scan(
			&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan,
			&p.BidangIndustri, &p.LokasiKerja, &p.GajiRange,
			&p.TanggalMulaiKerja, &p.TanggalSelesaiKerja, &p.StatusPekerjaan,
			&p.DeskripsiPekerjaan, &p.CreatedAt, &p.UpdatedAt,
		)
		list = append(list, p)
	}
	return list, nil
}

func CountPekerjaan(search string) (int, error) {
	var total int
	err := database.DB.QueryRow(`
		SELECT COUNT(*) FROM pekerjaan_alumni 
		WHERE nama_perusahaan ILIKE $1 
		   OR posisi_jabatan ILIKE $1 
		   OR bidang_industri ILIKE $1 
		   OR lokasi_kerja ILIKE $1`,
		"%"+search+"%",
	).Scan(&total)
	return total, err
}

func IsDeleted(role string, id int) error {
	if role == "admin" {
		// kalau admin → reset semua pekerjaan jadi isdelete=false
		fmt.Println("alumni hapus semua")
		_, err := database.DB.Exec(`
            UPDATE pekerjaan_alumni 
            SET updated_at = NOW(), isdelete = false`)
		if err != nil {
			return fmt.Errorf("gagal reset pekerjaan_alumni: %w", err)
		}
		return nil
	}
	var alumniID int
	err := database.DB.QueryRow(`SELECT alumni_id FROM users WHERE id=$1`, id).Scan(&alumniID)
	fmt.Println(err)
	if err != nil {
		return fmt.Errorf("gagal ambil alumni_id: %w", err)
	}
	fmt.Println(alumniID)
	_, err = database.DB.Exec(`
        UPDATE pekerjaan_alumni 
        SET updated_at = NOW(), isdelete = false
        WHERE alumni_id = $1`, alumniID)
	if err != nil {
		return fmt.Errorf("gagal update pekerjaan_alumni: %w", err)
	}

	return nil
}
