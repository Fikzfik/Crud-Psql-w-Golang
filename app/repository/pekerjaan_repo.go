package repository

import (
	"crud-alumni/app/models"
	"crud-alumni/database"
	"database/sql"
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

func RestoredUser(id int, p models.PekerjaanAlumni) ([]models.PekerjaanAlumni, error) {
	_, err := database.DB.Exec(`UPDATE pekerjaan_alumni 
	SET updated_at = NOW(), isdelete = true
	WHERE alumni_id = $1`, id)
	if err != nil {
		return nil, err
	}
	rows, err := database.DB.Query(`SELECT * FROM pekerjaan_alumni WHERE alumni_id=$1 ORDER BY id`, id)
	var list []models.PekerjaanAlumni
	for rows.Next() {
		var p models.PekerjaanAlumni
		rows.Scan(&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan,
			&p.BidangIndustri, &p.LokasiKerja, &p.GajiRange,
			&p.TanggalMulaiKerja, &p.TanggalSelesaiKerja,
			&p.StatusPekerjaan, &p.DeskripsiPekerjaan,
			&p.CreatedAt, &p.UpdatedAt, &p.IsDeleted)
		list = append(list, p)
	}
	return list, nil
}

func Userdeleted(id int, p models.PekerjaanAlumni) ([]models.PekerjaanAlumni, error) {
	_, err := database.DB.Exec(`UPDATE pekerjaan_alumni 
	SET updated_at = NOW(), isdelete = false
	WHERE alumni_id = $1`, id)
	if err != nil {
		return nil, err
	}
	rows, err := database.DB.Query(`SELECT * FROM pekerjaan_alumni WHERE alumni_id=$1 ORDER BY id`, id)
	var list []models.PekerjaanAlumni
	for rows.Next() {
		var p models.PekerjaanAlumni
		rows.Scan(&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan,
			&p.BidangIndustri, &p.LokasiKerja, &p.GajiRange,
			&p.TanggalMulaiKerja, &p.TanggalSelesaiKerja,
			&p.StatusPekerjaan, &p.DeskripsiPekerjaan,
			&p.CreatedAt, &p.UpdatedAt, &p.IsDeleted)
		list = append(list, p)
	}
	return list, nil
}

func IsAktif(role string, id int) ([]models.PekerjaanAlumni, error) {
	var (
		rows     *sql.Rows
		err      error
		alumniID int
	)

	fmt.Println("User ID:", id)

	// Ambil alumni_id dari tabel users kalau role bukan admin
	if role != "admin" {
		err = database.DB.QueryRow(`SELECT alumni_id FROM users WHERE id=$1`, id).Scan(&alumniID)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, fmt.Errorf("user dengan id %d tidak ditemukan", id)
			}
			return nil, fmt.Errorf("gagal ambil alumni_id: %v", err)
		}
		fmt.Println("Alumni ID:", alumniID)
	}

	queryAdmin := `
	SELECT id, alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, 
	       lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja, 
	       status_pekerjaan, deskripsi_pekerjaan, created_at, updated_at, isdelete 
	FROM pekerjaan_alumni 
	WHERE isdelete = true
`

	queryUser := queryAdmin + ` AND alumni_id = $1`

	// Jalankan query
	if role == "admin" {
		rows, err = database.DB.Query(queryAdmin)
	} else {
		rows, err = database.DB.Query(queryUser, alumniID)
	}
	if err != nil {
		return nil, fmt.Errorf("query error: %v", err)
	}
	defer rows.Close()

	var list []models.PekerjaanAlumni
	for rows.Next() {
		var p models.PekerjaanAlumni
		if err := rows.Scan(
			&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan,
			&p.BidangIndustri, &p.LokasiKerja, &p.GajiRange,
			&p.TanggalMulaiKerja, &p.TanggalSelesaiKerja,
			&p.StatusPekerjaan, &p.DeskripsiPekerjaan,
			&p.CreatedAt, &p.UpdatedAt, &p.IsDeleted,
		); err != nil {
			return nil, fmt.Errorf("scan error: %v", err)
		}
		list = append(list, p)
	}

	if len(list) == 0 {
		return []models.PekerjaanAlumni{}, nil
	}

	return list, nil
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
	var alumniID int
	//cek admin/bukan
	if role == "admin" {
		//mengisi id dengan parameter user lain
		alumniID = id
		} else {
			//mengambil id diri sendiri
			err := database.DB.QueryRow(`SELECT alumni_id FROM users WHERE id=$1`, id).Scan(&alumniID)
			if err != nil {
			return fmt.Errorf("gagal ambil alumni_id: %w", err)
		}
	}
	
	//mengganti soft delete dari id yg login/parameter
	_, err := database.DB.Exec(`
		UPDATE pekerjaan_alumni 
		SET updated_at = NOW(), isdelete = true
		WHERE id= $1`, alumniID)
	if err != nil {
		return fmt.Errorf("gagal update pekerjaan_alumni: %w", err)
	}
	return nil
}

func IsRestored(role string, id int) error {
	var alumniID int

	if role == "admin" {
		alumniID = id
	} else {
		err := database.DB.QueryRow(`SELECT alumni_id FROM users WHERE id=$1`, id).Scan(&alumniID)
		if err != nil {
			return fmt.Errorf("gagal ambil alumni_id: %w", err)
		}
	}

	_, err := database.DB.Exec(`
		UPDATE pekerjaan_alumni 
		SET updated_at = NOW(), isdelete = false
		WHERE id = $1`, alumniID)
	if err != nil {
		return fmt.Errorf("gagal update pekerjaan_alumni: %w", err)
	}
	return nil
}

