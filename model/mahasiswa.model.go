package model

type Mahasiswa struct {
	Id                int    `json:"id"`
	Nama              string `json:"nama"`
	Usia              int    `json:"usia"`
	Gender            int    `json:"gender"`
	TanggalRegistrasi string `json:"tanggalRegistrasi"`
	Jurusan           int    `json:"jurusan"`
	NamaJurusan       string `json:"namaJurusan"`
	NamaHobi          string `json:"namaHobi"`
}
