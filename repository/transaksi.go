package repository

import (
	"FinalProject/structs"
	"database/sql"
	"errors"
)

func GetAllTransaksi(db *sql.DB) (result []structs.Transaksi, err error) {
	query := `SELECT id, kode_pelanggan, kode_barang, kode_mata_uang, jumlah_barang, total_harga, tanggal, created_at, created_by, modified_at, modified_by 
			  FROM transaksi`

	rows, err := db.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var t structs.Transaksi
		err = rows.Scan(
			&t.ID,
			&t.KodePelanggan,
			&t.KodeBarang,
			&t.KodeMataUang,
			&t.JumlahBarang,
			&t.TotalHarga,
			&t.Tanggal,
			&t.CreatedAt,
			&t.CreatedBy,
			&t.ModifiedAt,
			&t.ModifiedBy,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, t)
	}

	return
}

func GetOneTransaksi(db *sql.DB, id int) (t structs.Transaksi, err error) {
	query := `SELECT id, kode_pelanggan, kode_barang, kode_mata_uang, jumlah_barang, total_harga, tanggal, created_at, created_by, modified_at, modified_by 
			  FROM transaksi 
			  WHERE id = $1`

	err = db.QueryRow(query, id).Scan(
		&t.ID,
		&t.KodePelanggan,
		&t.KodeBarang,
		&t.KodeMataUang,
		&t.JumlahBarang,
		&t.TotalHarga,
		&t.Tanggal,
		&t.CreatedAt,
		&t.CreatedBy,
		&t.ModifiedAt,
		&t.ModifiedBy,
	)

	return
}

func InsertTransaksi(db *sql.DB, t structs.Transaksi) error {
	var stok int
	err := db.QueryRow(`SELECT jumlah_barang FROM barang WHERE kode_barang = $1`, t.KodeBarang).Scan(&stok)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("barang tidak ditemukan")
		}
		return err
	}

	if stok < t.JumlahBarang {
		return errors.New("stok barang tidak mencukupi")
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`INSERT INTO transaksi 
		(kode_pelanggan, kode_barang, kode_mata_uang, jumlah_barang, total_harga, tanggal, created_at, created_by, modified_at, modified_by) 
		VALUES ($1, $2, $3, $4, $5, $6, NOW(), $7, NOW(), $8)`,
		t.KodePelanggan, t.KodeBarang, t.KodeMataUang, t.JumlahBarang, t.TotalHarga, t.Tanggal, t.CreatedBy, t.ModifiedBy,
	)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`UPDATE barang SET jumlah_barang = jumlah_barang - $1, modified_at = NOW(), modified_by = $2 
		WHERE kode_barang = $3`,
		t.JumlahBarang, t.ModifiedBy, t.KodeBarang,
	)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func UpdateTransaksi(db *sql.DB, t structs.Transaksi) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	var stok int
	err = db.QueryRow(`SELECT jumlah_barang FROM barang WHERE kode_barang = $1`, t.KodeBarang).Scan(&stok)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("barang tidak ditemukan")
		}
		return err
	}
	var jumlahBarangNow int
	err2 := db.QueryRow(`SELECT jumlah_barang from transaksi where id = $1`, t.ID).Scan(&jumlahBarangNow)
	if err2 != nil {
		if err2 == sql.ErrNoRows {
			return errors.New("barang tidak ditemukan")
		}
		return err2
	}
	diff := jumlahBarangNow - t.JumlahBarang
	if diff > 0 && stok < diff {
		return errors.New("stok barang tidak mencukupi untuk update transaksi")
	}

	total := stok + diff

	_, err = tx.Exec(`UPDATE transaksi 
		SET kode_pelanggan=$1, kode_barang=$2, kode_mata_uang=$3, jumlah_barang=$4, total_harga=$5, tanggal=$6, modified_at=NOW(), modified_by=$7 
		WHERE id=$8`,
		t.KodePelanggan, t.KodeBarang, t.KodeMataUang, t.JumlahBarang, t.TotalHarga, t.Tanggal, t.ModifiedBy, t.ID,
	)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`UPDATE barang 
		SET jumlah_barang = jumlah_barang - $1, modified_at=NOW(), modified_by=$2
		WHERE kode_barang=$3`,
		total, t.ModifiedBy, t.KodeBarang,
	)
	if err != nil {
		return err
	}

	return tx.Commit()

}

// DeleteTransaksi menghapus data transaksi berdasarkan ID
func DeleteTransaksi(db *sql.DB, id int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var kodeBarang string
	var jumlah int
	err = tx.QueryRow(`SELECT kode_barang, jumlah_barang FROM transaksi WHERE id=$1`, id).Scan(&kodeBarang, &jumlah)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("transaksi not found")
		}
		return err
	}

	_, err = tx.Exec(`DELETE FROM transaksi WHERE id=$1`, id)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`UPDATE barang SET jumlah_barang = jumlah_barang + $1 WHERE kode_barang=$2`, jumlah, kodeBarang)
	if err != nil {
		return err
	}

	return tx.Commit()
}
