package commons

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/silviorosas/crud-golang-svelte/models"
)

func GetConexion() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@/dbgolang1?charset=utf8")

	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Migrate() {
	db := GetConexion()
	defer db.Close()

	log.Println("Migrando.....")
	db.AutoMigrate(&models.Persona{})
}
