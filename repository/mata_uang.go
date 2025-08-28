package repository

import (
	"FinalProject/structs"
	"database/sql"
	"errors"
	"time"
)

func GetAllMataUang(db *sql.DB) (result []structs.MataUang, err error) {
	query := `SELECT kode_mata_uang, nama_mata_uang, created_at, created_by, modified_at, modified_by 
	FROM mata_uang`

	rows, err := db.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var mu structs.MataUang
		err = rows.Scan(&mu.KodeMataUang, &mu.NamaMataUang, &mu.CreatedAt, &mu.CreatedBy, &mu.ModifiedAt, &mu.ModifiedBy)
		if err != nil {
			return nil, err
		}
		result = append(result, mu)
	}

	return
}

func GetOneMataUang(db *sql.DB, kode string) (mu structs.MataUang, err error) {
	query := `SELECT kode_mata_uang, nama_mata_uang, created_at, created_by, modified_at, modified_by
	          FROM mata_uang 
			  WHERE kode_mata_uang = $1`

	err = db.QueryRow(query, kode).Scan(
		&mu.KodeMataUang,
		&mu.NamaMataUang,
		&mu.CreatedAt,
		&mu.CreatedBy,
		&mu.ModifiedAt,
		&mu.ModifiedBy,
	)

	return
}

func InsertMataUang(db *sql.DB, mu structs.MataUang) error {
	_, err := db.Exec(`
        INSERT INTO mata_uang (kode_mata_uang, nama_mata_uang, created_at, created_by, modified_at, modified_by) 
        VALUES ($1, $2, NOW(), $3, NOW(), $4)`,
		mu.KodeMataUang, mu.NamaMataUang, mu.CreatedBy, mu.ModifiedBy,
	)
	return err
}

func UpdateMataUang(db *sql.DB, mu structs.MataUang) error {
	res, err := db.Exec(`UPDATE mata_uang 
	                     SET nama_mata_uang=$1, modified_at=$2, modified_by=$3 
	                     WHERE kode_mata_uang=$4`,
		mu.NamaMataUang, time.Now(), mu.ModifiedBy, mu.KodeMataUang,
	)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("mata_uang not found")
	}
	return nil
}

// DeleteMataUang menghapus data mata_uang
func DeleteMataUang(db *sql.DB, kode string) error {
	res, err := db.Exec(`DELETE FROM mata_uang WHERE kode_mata_uang=$1`, kode)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("mata_uang not found")
	}
	return nil
}
