package structs

type MataUang struct {
	KodeMataUang string `json:"kode_mata_uang"`
	NamaMataUang string `json:"nama_mata_uang"`
	CreatedAt    string `json:"created_at"`
	CreatedBy    string `json:"created_by"`
	ModifiedAt   string `json:"modified_at"`
	ModifiedBy   string `json:"modified_by"`
}

type Kurs struct {
	KodeMataUang string  `json:"kode_mata_uang"`
	Tanggal      string  `json:"tanggal"`
	Kurs         float64 `json:"kurs"`
	CreatedAt    string  `json:"created_at"`
	CreatedBy    string  `json:"created_by"`
	ModifiedAt   string  `json:"modified_at"`
	ModifiedBy   string  `json:"modified_by"`
}

type Pelanggan struct {
	KodePelanggan string `json:"kode_pelanggan"`
	NamaPelanggan string `json:"nama_pelanggan"`
	CreatedAt     string `json:"created_at"`
	CreatedBy     string `json:"created_by"`
	ModifiedAt    string `json:"modified_at"`
	ModifiedBy    string `json:"modified_by"`
}

type Barang struct {
	KodeBarang   string `json:"kode_barang"`
	NamaBarang   string `json:"nama_barang"`
	JumlahBarang int    `json:"jumlah_barang"`
	CreatedAt    string `json:"created_at"`
	CreatedBy    string `json:"created_by"`
	ModifiedAt   string `json:"modified_at"`
	ModifiedBy   string `json:"modified_by"`
}

type Transaksi struct {
	ID            int     `json:"id"`
	KodePelanggan string  `json:"kode_pelanggan"`
	KodeBarang    string  `json:"kode_barang"`
	KodeMataUang  string  `json:"kode_mata_uang"`
	JumlahBarang  int     `json:"jumlah_barang"`
	TotalHarga    float64 `json:"total_harga"`
	Tanggal       string  `json:"tanggal"`
	CreatedAt     string  `json:"created_at"`
	CreatedBy     string  `json:"created_by"`
	ModifiedAt    string  `json:"modified_at"`
	ModifiedBy    string  `json:"modified_by"`
}

type User struct {
	ID         int    `json:"id"`
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	CreatedAt  string `json:"created_at"`
	CreatedBy  string `json:"created_by"`
	ModifiedAt string `json:"modified_at"`
	ModifiedBy string `json:"modified_by"`
}

type LaporanPembelianKonversi struct {
	Tanggal       string `json:"tanggal"`
	NamaPelanggan string `json:"nama_pelanggan"`
	Produk        string `json:"produk"`
	Jumlah        string `json:"jumlah"`
	Total         string `json:"total"`
}
