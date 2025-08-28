package repository

import (
	"FinalProject/structs"
	"database/sql"
)

func GetAllLaporanPembelianKonversi(db *sql.DB) ([]structs.LaporanPembelianKonversi, error) {
	rows, err := db.Query(`SELECT tanggal, nama_pelanggan, produk, jumlah, total FROM v_laporan_pembelian_konversi`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var laporan []structs.LaporanPembelianKonversi

	for rows.Next() {
		var l structs.LaporanPembelianKonversi
		err := rows.Scan(&l.Tanggal, &l.NamaPelanggan, &l.Produk, &l.Jumlah, &l.Total)
		if err != nil {
			return nil, err
		}
		laporan = append(laporan, l)
	}

	return laporan, nil
}
