package services
import (
	"PJApp/app/models"
	"github.com/revel/revel"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

)

var DB gorm.DB

func InitDB() {
	var err error
	dbStr := "pjuser:pjpass@tcp(localhost:3306)/pjapp?charset=utf8&parseTime=True"
	if DB, err = gorm.Open("mysql", dbStr); err!= nil {
		revel.ERROR.Println("FATAL", err)
		panic(err)
	}
	// Tableの作成
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Report{})
}