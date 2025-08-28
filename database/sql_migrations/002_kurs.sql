-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE IF NOT EXISTS kurs (
    kode_mata_uang VARCHAR(10) NOT NULL,
    tanggal DATE NOT NULL,
    kurs NUMERIC(18,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(50),
    modified_at TIMESTAMP,
    modified_by VARCHAR(50)
);

-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin
DROP TABLE IF EXISTS kurs;
-- +migrate StatementEnd