package entities

type Model_japan struct {
	Japan_id     int    `json:"japan_id"`
	Japan_date   string `json:"japan_date"`
	Japan_prize1 string `json:"japan_prize1"`
	Japan_prize2 string `json:"japan_prize2"`
	Japan_prize3 string `json:"japan_prize3"`
	Japan_create string `json:"japan_create"`
	Japan_update string `json:"japan_update"`
}

type Controller_japan struct {
	Japan_tipe string `json:"japan_tipe"`
	Japan_page int    `json:"japan_page"`
}
type Controller_japansave struct {
	Sdata    string `json:"sdata" validate:"required"`
	Page     string `json:"page" validate:"required"`
	Idrecord int    `json:"idrecord"`
	Tanggal  string `json:"tanggal" validate:"required"`
	Tipe     string `json:"tipe" validate:"required"`
}
type Controller_japanprizesave struct {
	Sdata     string `json:"sdata" validate:"required"`
	Page      string `json:"page" validate:"required"`
	Idrecord  int    `json:"idrecord"`
	Tipe      string `json:"tipe" validate:"required"`
	Tipejapan string `json:"tipejapan" validate:"required"`
	Prize     string `json:"prize" validate:"required"`
}
