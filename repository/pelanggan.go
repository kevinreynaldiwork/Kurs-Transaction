package repository

import (
	"FinalProject/structs"
	"database/sql"
	"errors"
)

// GetAllPelanggan mengambil semua data pelanggan
func GetAllPelanggan(db *sql.DB) (result []structs.Pelanggan, err error) {
	query := `SELECT kode_pelanggan, nama_pelanggan, created_at, created_by, modified_at, modified_by 
			  FROM pelanggan`

	rows, err := db.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var p structs.Pelanggan
		err = rows.Scan(&p.KodePelanggan, &p.NamaPelanggan, &p.CreatedAt, &p.CreatedBy, &p.ModifiedAt, &p.ModifiedBy)
		if err != nil {
			return nil, err
		}
		result = append(result, p)
	}

	return
}

// GetOnePelanggan mengambil data pelanggan berdasarkan kode
func GetOnePelanggan(db *sql.DB, kode string) (p structs.Pelanggan, err error) {
	query := `SELECT kode_pelanggan, nama_pelanggan, created_at, created_by, modified_at, modified_by
	          FROM pelanggan 
			  WHERE kode_pelanggan = $1`

	err = db.QueryRow(query, kode).Scan(
		&p.KodePelanggan,
		&p.NamaPelanggan,
		&p.CreatedAt,
		&p.CreatedBy,
		&p.ModifiedAt,
		&p.ModifiedBy,
	)

	return
}

// InsertPelanggan menambahkan pelanggan baru
func InsertPelanggan(db *sql.DB, p structs.Pelanggan) error {
	_, err := db.Exec(`INSERT INTO pelanggan 
		(kode_pelanggan, nama_pelanggan, created_at, created_by, modified_at, modified_by) 
		VALUES ($1, $2, NOW(), $3, NOW(), $4)`,
		p.KodePelanggan, p.NamaPelanggan, p.CreatedBy, p.ModifiedBy,
	)
	return err
}

// UpdatePelanggan mengubah data pelanggan
func UpdatePelanggan(db *sql.DB, p structs.Pelanggan) error {
	res, err := db.Exec(`UPDATE pelanggan 
		SET nama_pelanggan=$1, modified_at=NOW(), modified_by=$2 
		WHERE kode_pelanggan=$3`,
		p.NamaPelanggan, p.ModifiedBy, p.KodePelanggan,
	)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("pelanggan not found")
	}
	return nil
}

// DeletePelanggan menghapus data pelanggan
func DeletePelanggan(db *sql.DB, kode string) error {
	res, err := db.Exec(`DELETE FROM pelanggan WHERE kode_pelanggan=$1`, kode)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("pelanggan not found")
	}
	return nil
}
