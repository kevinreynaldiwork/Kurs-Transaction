-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE IF NOT EXISTS mata_uang (
    kode_mata_uang VARCHAR(10) PRIMARY KEY,
    nama_mata_uang VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(50),
    modified_at TIMESTAMP,
    modified_by VARCHAR(50)
);
-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin
DROP TABLE IF EXISTS mata_uang;
-- +migrate StatementEnd