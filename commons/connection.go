package commons

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/nicolascalvo73/crud-api-rest-go/models"
)

func GetConnections() *gorm.DB { //retorna un *gorm.DB
	db, err := gorm.Open("mysql", "nicolascalvo73:Zzulrgx6@/user_db?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Migrate() {
	db := GetConnections()
	defer db.Close()

	log.Println("Migrating...")

	db.AutoMigrate(&models.User{})
}
