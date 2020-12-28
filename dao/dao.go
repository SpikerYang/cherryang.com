package dao
import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var db *gorm.DB

func InitDatabase(dataSource string) {
	database, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = database
}

func resultProcess(result *gorm.DB) {
	if err := result.Error; err != nil {
		panic(err)
	}
}



