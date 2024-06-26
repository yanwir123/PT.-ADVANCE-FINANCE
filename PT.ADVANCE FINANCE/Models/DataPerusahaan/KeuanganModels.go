package dataperusahaan

type Keuangan struct {
	Id         int     `json:"id" from:"id"`
	Bulan      string  `json:"Bulan" from:"Bulan"`
	Product    string  `json:"Product" from:"Product"`
	Masuk      float64 `json:"Masuk" from:"Masuk"`
	Keluar     float64 `json:"Keluar" from:"Keluar"`
	Stok       float64 `json:"Stok" from:"Stok"`
	Total      float64 `json:"Total" from:"Total"`
	Keterangan string  `json:"Keterangan" from:"Keterangan"`
}