-- +migrate Up
-- +migrate StatementBegin
CREATE OR REPLACE VIEW v_laporan_pembelian_konversi AS
SELECT 
    a.tanggal,
    a.nama_pelanggan,
    a.nama_barang AS produk,
    to_char(a.jumlah_barang, 'FM999,999,999,990') AS jumlah,
    CASE 
        WHEN a.angkaKurs * a.total_harga = 0 
            THEN '' 
        ELSE to_char((a.angkaKurs * a.total_harga), 'FM999,999,999,990.00')
    END AS total
FROM (
    SELECT 
        t.*,
        p.nama_pelanggan,
        pr.nama_barang,
        k.tanggal AS kurs_tanggal, 
        COALESCE(k.kurs, 0) AS angkaKurs
    FROM transaksi t
    LEFT JOIN pelanggan p 
        ON t.kode_pelanggan = p.kode_pelanggan
    LEFT JOIN barang pr
        ON t.kode_barang = pr.kode_barang
    LEFT JOIN kurs k 
      ON k.kode_mata_uang = t.kode_mata_uang
     AND k.tanggal = (
          SELECT MAX(k2.tanggal)
          FROM kurs k2
          WHERE k2.kode_mata_uang = t.kode_mata_uang
            AND k2.tanggal <= t.tanggal
      )
) a;
-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin
DROP VIEW IF EXISTS v_laporan_pembelian_konversi;
-- +migrate StatementEnd
