package entities

type Model_vietnamenight struct {
	Vietnamnight_id          int    `json:"vietnamnight_id"`
	Vietnamnight_date        string `json:"vietnamnight_date"`
	Vietnamnight_prize1_1300 string `json:"vietnamnight_prize1_1300"`
	Vietnamnight_prize1_1700 string `json:"vietnamnight_prize1_1700"`
	Vietnamnight_prize1_2000 string `json:"vietnamnight_prize1_2000"`
	Vietnamnight_prize1_2200 string `json:"vietnamnight_prize1_2200"`
	Vietnamnight_create      string `json:"vietnamnight_create"`
	Vietnamnight_update      string `json:"vietnamnight_update"`
}
type Controller_vietnamenightsave struct {
	Sdata    string `json:"sdata" validate:"required"`
	Page     string `json:"page" validate:"required"`
	Idrecord int    `json:"idrecord"`
	Tanggal  string `json:"tanggal" validate:"required"`
}
type Controller_vietnamenightprizesave struct {
	Sdata    string `json:"sdata" validate:"required"`
	Page     string `json:"page" validate:"required"`
	Idrecord int    `json:"idrecord"`
	Tipe     string `json:"tipe" validate:"required"`
	Prize    string `json:"prize" validate:"required"`
}
type Responseredis_vietnamenight struct {
	Vietnamnight_id          int    `json:"vietnamnight_id"`
	Vietnamnight_date        string `json:"vietnamnight_date"`
	Vietnamnight_prize1_1300 string `json:"vietnamnight_prize1_1300"`
	Vietnamnight_prize1_1700 string `json:"vietnamnight_prize1_1700"`
	Vietnamnight_prize1_2000 string `json:"vietnamnight_prize1_2000"`
	Vietnamnight_prize1_2200 string `json:"vietnamnight_prize1_2200"`
	Vietnamnight_create      string `json:"vietnamnight_create"`
	Vietnamnight_update      string `json:"vietnamnight_update"`
}
