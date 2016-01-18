package models

import "time"

type Report struct {
	Id         int
	UserId     int
	CreateDay  string     `sql:"type:char(14);"`
	UpdateDay  string     `sql:"type:char(14);"`
	StartDay   string     `sql:"type:char(8);"`
	Kotei      string     `sql:"type:varchar(1000);"`
	Hokoku     string     `sql:"type:varchar(1000);"`
	Sagyo1     string     `sql:"type:varchar(200);"`
	Kadotime1  float32    `sql:"type:float;unsigned;"`
	Sagyo2     string     `sql:"type:varchar(200);"`
	Kadotime2  float32    `sql:"type:float;unsigned;"`
	Sagyo3     string     `sql:"type:varchar(200);"`
	Kadotime3  float32    `sql:"type:float;unsigned;"`
	Sagyo4     string     `sql:"type:varchar(200);"`
	Kadotime4  float32    `sql:"type:float;unsigned;"`
	Sagyo5     string     `sql:"type:varchar(200);"`
	Kadotime5  float32    `sql:"type:float;unsigned;"`
	Sagyo6     string     `sql:"type:varchar(200);"`
	Kadotime6  float32    `sql:"type:float;unsigned;"`
	Sagyo7     string     `sql:"type:varchar(200);"`
	Kadotime7  float32    `sql:"type:float;unsigned;"`
	Sagyo8     string     `sql:"type:varchar(200);"`
	Kadotime8  float32    `sql:"type:float;unsigned;"`
	Sagyo9     string     `sql:"type:varchar(200);"`
	Kadotime9  float32    `sql:"type:float;unsigned;"`
	Sagyo10    string     `sql:"type:varchar(200);"`
	Kadotime10 float32    `sql:"type:float;unsigned;"`
	Sagyo11    string     `sql:"type:varchar(200);"`
	Kadotime11 float32    `sql:"type:float;unsigned;"`
	Sagyo12    string     `sql:"type:varchar(200);"`
	Kadotime12 float32    `sql:"type:float;unsigned;"`
	Sagyo13    string     `sql:"type:varchar(200);"`
	Kadotime13 float32    `sql:"type:float;unsigned;"`
	Sagyo14    string     `sql:"type:varchar(200);"`
	Kadotime14 float32    `sql:"type:float;unsigned;"`
	Kadai      string     `sql:"type:varchar(1000);"`
	Other      string     `sql:"type:varchar(500);"`
	Status     string     `sql:"type:char(1);not null"`

}

func InitReport (startDay time.Time) *Report{
	report := Report{
		Kotei: "a",
		Hokoku: "b",
		Sagyo1: "sagyo1",
		Kadotime1: 8,
		Sagyo2: "sagyo2",
		Kadotime2: 8.5,
		Sagyo3: "",
		Kadotime3: 0,
		Sagyo4: "",
		Kadotime4: 0,
		Sagyo5: "",
		Kadotime5: 0,
		Sagyo6: "",
		Kadotime6: 0,
		Sagyo7: "",
		Kadotime7: 0,
		Sagyo8: "",
		Kadotime8: 0,
		Sagyo9: "",
		Kadotime9: 0,
		Sagyo10: "",
		Kadotime10: 0,
		Sagyo11: "",
		Kadotime11: 0,
		Sagyo12: "",
		Kadotime12: 0,
		Sagyo13: "",
		Kadotime13: 0,
		Sagyo14: "",
		Kadotime14: 0,
		Kadai: "kadai",
		Other: "other"}

	return &report
}
