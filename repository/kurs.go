package repository

import (
	"FinalProject/structs"
	"database/sql"
	"errors"
)

func GetAllKurs(db *sql.DB) (result []structs.Kurs, err error) {
	query := `SELECT kode_mata_uang, tanggal, kurs, created_at, created_by, modified_at, modified_by 
	          FROM kurs`

	rows, err := db.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var k structs.Kurs
		err = rows.Scan(
			&k.KodeMataUang,
			&k.Tanggal,
			&k.Kurs,
			&k.CreatedAt,
			&k.CreatedBy,
			&k.ModifiedAt,
			&k.ModifiedBy,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, k)
	}

	return
}

func GetOneKurs(db *sql.DB, kode string, tanggal string) (k structs.Kurs, err error) {
	query := `SELECT kode_mata_uang, tanggal, kurs, created_at, created_by, modified_at, modified_by
	          FROM kurs 
	          WHERE kode_mata_uang=$1 AND tanggal=$2`

	err = db.QueryRow(query, kode, tanggal).Scan(
		&k.KodeMataUang,
		&k.Tanggal,
		&k.Kurs,
		&k.CreatedAt,
		&k.CreatedBy,
		&k.ModifiedAt,
		&k.ModifiedBy,
	)

	if err == sql.ErrNoRows {
		return k, errors.New("kurs not found")
	}

	return
}

func InsertKurs(db *sql.DB, k structs.Kurs) error {
	_, err := db.Exec(`INSERT INTO kurs 
		(kode_mata_uang, tanggal, kurs, created_at, created_by, modified_at, modified_by) 
		VALUES ($1, $2, $3, NOW(), $4, NOW(), $5)`,
		k.KodeMataUang, k.Tanggal, k.Kurs, k.CreatedBy, k.ModifiedBy,
	)
	return err
}

func UpdateKurs(db *sql.DB, k structs.Kurs) error {
	res, err := db.Exec(`UPDATE kurs 
		SET tanggal = $1, kurs = $2, modified_at = NOW(), modified_by = $3
		WHERE kode_mata_uang = $4 AND tanggal = $5`,
		k.Tanggal,
		k.Kurs,
		k.ModifiedBy,
		k.KodeMataUang,
		k.Tanggal,
	)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("kurs not found")
	}
	return nil
}

func DeleteKurs(db *sql.DB, kode string, tanggal string) error {
	res, err := db.Exec(`DELETE FROM kurs WHERE kode_mata_uang=$1 AND tanggal=$2`, kode, tanggal)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("kurs not found")
	}
	return nil
}
