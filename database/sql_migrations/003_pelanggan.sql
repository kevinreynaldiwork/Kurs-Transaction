-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE IF NOT EXISTS pelanggan (
    kode_pelanggan VARCHAR(20) PRIMARY KEY,
    nama_pelanggan VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(50),
    modified_at TIMESTAMP,
    modified_by VARCHAR(50)
);

-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin
DROP TABLE IF EXISTS pelanggan;
-- +migrate StatementEnd