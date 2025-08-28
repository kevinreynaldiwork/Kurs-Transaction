package repository

import (
	"FinalProject/structs"
	"database/sql"
	"errors"
)

// GetAllBarang mengambil semua data barang
func GetAllBarang(db *sql.DB) (result []structs.Barang, err error) {
	query := `SELECT kode_barang, nama_barang, jumlah_barang, created_at, created_by, modified_at, modified_by 
			  FROM barang`

	rows, err := db.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var b structs.Barang
		err = rows.Scan(&b.KodeBarang, &b.NamaBarang, &b.JumlahBarang, &b.CreatedAt, &b.CreatedBy, &b.ModifiedAt, &b.ModifiedBy)
		if err != nil {
			return nil, err
		}
		result = append(result, b)
	}

	return
}

// GetOneBarang mengambil data barang berdasarkan kode
func GetOneBarang(db *sql.DB, kode string) (b structs.Barang, err error) {
	query := `SELECT kode_barang, nama_barang, jumlah_barang, created_at, created_by, modified_at, modified_by
	          FROM barang 
			  WHERE kode_barang = $1`

	err = db.QueryRow(query, kode).Scan(
		&b.KodeBarang,
		&b.NamaBarang,
		&b.JumlahBarang,
		&b.CreatedAt,
		&b.CreatedBy,
		&b.ModifiedAt,
		&b.ModifiedBy,
	)

	return
}

// InsertBarang menambahkan barang baru
func InsertBarang(db *sql.DB, b structs.Barang) error {
	_, err := db.Exec(`INSERT INTO barang 
		(kode_barang, nama_barang, jumlah_barang, created_at, created_by,modified_at, modified_by ) 
		VALUES ($1, $2, $3, NOW(), $4, NOW(), $5)`,
		b.KodeBarang, b.NamaBarang, b.JumlahBarang, b.CreatedBy, b.ModifiedBy,
	)
	return err
}

// UpdateBarang mengubah data barang
func UpdateBarang(db *sql.DB, b structs.Barang) error {
	res, err := db.Exec(`UPDATE barang 
		SET nama_barang=$1, jumlah_barang=$2, modified_at=NOW(), modified_by=$3 
		WHERE kode_barang=$4`,
		b.NamaBarang, b.JumlahBarang, b.ModifiedBy, b.KodeBarang,
	)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("barang not found")
	}
	return nil
}

// DeleteBarang menghapus data barang
func DeleteBarang(db *sql.DB, kode string) error {
	res, err := db.Exec(`DELETE FROM barang WHERE kode_barang=$1`, kode)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("barang not found")
	}
	return nil
}
