-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE IF NOT EXISTS transaksi (
    id SERIAL PRIMARY KEY,
    kode_pelanggan VARCHAR(20),
    kode_barang VARCHAR(20),
    kode_mata_uang VARCHAR(10),
    jumlah_barang INT NOT NULL,
    total_harga NUMERIC(18,2) NOT NULL,
    tanggal DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(50),
    modified_at TIMESTAMP,
    modified_by VARCHAR(50)
);

-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin
DROP TABLE IF EXISTS transaksi;
-- +migrate StatementEnd