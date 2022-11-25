package models

import "time"

type Biodata struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Nama      string    `json:"nama"`
	Alamat    string    `json:"alamat"`
	Jurusan   string    `json:"jurusan"`
	Pekerjaan string    `json:"pekerjaan"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Barang struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	KodeBarang  string    `json:"kode_barang"`
	NamaBarang  string    `json:"nama_barang"`
	JenisBarang string    `json:"jenis_barang"`
	HargaBeli   int       `json:"harga_beli"`
	HargaJual   int       `json:"harga_jual"`
	Kategori    string    `json:"kategori"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
