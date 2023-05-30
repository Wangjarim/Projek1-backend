package model

type Pasien struct {
	Id           int    `gorm:"primaryKey;column:id;autoIncrement" json:"id"`
	Namalengkap  string `gorm:"column:nama_lengkap" json:"nama_lengkap"`
	Nik          int    `gorm:"column:nik" json:"nik"`
	Jeninkelamin string `gorm:"column:jenis_kelamin" json:"jenis_kelamin"`
	Tempatlahir  string `gorm:"column:tempat_lahir" json:"tempat_lahir"`
	Tanggallahir string `gorm:"column:tanggal_lahir" json:"tanggal_lahir"`
	Alamat       string `gorm:"column:alamat" json:"alamat"`
	Nohp         string `gorm:"column:no_hp" json:"no_hp"`
}

type Dokter struct {
	Id int `gorm:"primaryKey;column:id;autoIncrement" json:"id"`
	//IdDokter int    `gorm:"column:id_dokter" json:"id_dokter"`
	Nama     string `gorm:"column:nama" json:"nama"`
	Keahlian string `gorm:"column:keahlian" json:"keahlian"`
	Nohp     string `gorm:"column:no_hp" json:"no_hp"`
}

type Hari struct {
	Id   int    `gorm:"primaryKey;column:id;autoIncrement" json:"id"`
	Hari string `gorm:"column:hari" json:"hari"`
}

type Jam struct {
	Id  int    `gorm:"primaryKey;column:id;autoIncrement" json:"id"`
	Jam string `gorm:"column:jam" json:"jam"`
}

type JadwalDokter struct {
	Id       int    `gorm:"primaryKey;column:id;autoIncrement" json:"id"`
	IdDokter int    `gorm:"column:id_dokter;foreignKey:DokterId" json:"id_dokter"`
	IdHari   int    `gorm:"column:id_hari;foreignKey:HariId" json:"id_hari"`
	IdJam    int    `gorm:"column:id_jam;foreignKey:JamId" json:"id_jam"`
	Dokter   Dokter `gorm:"foreignKey:IdDokter" json:"dokter"`
	Hari     Hari   `gorm:"foreignKey:IdHari" json:"hari"`
	Jam      Jam    `gorm:"foreignKey:IdJam" json:"jam"`
}
