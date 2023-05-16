package model

type Pasien struct {
	Id            int    `gorm:"primaryKey;column:id;autoIncrement" json:"id"`
	Nama_lengkap  string `gorm:"column:nama_lengkap" json:"nama_lengkap"`
	Nik           int    `gorm:"column:nik" json:"nik"`
	Jenin_kelamin string `gorm:"column:jenis_kelamin" json:"jenis_kelamin"`
	Tempat_lahir  string `gorm:"column:tempat_lahir" json:"tempat_lahir"`
	Tanggal_lahir string `gorm:"column:tanggal_lahir" json:"tanggal_lahir"`
	Alamat        string `gorm:"column:alamat" json:"alamat"`
	No_hp         string `gorm:"column:no_hp" json:"no_hp"`
}
