package migration

import(
	"fmt"
	"github.com/adyfp24/golang-jwt-auth/database"
	"github.com/adyfp24/golang-jwt-auth/models"
	"log"
)

func RunMigration(){
	db, err := database.InitDB()
	err = db.AutoMigrate(&models.User{})
	if err != nil{
		log.Print(err)
	}
	fmt.Println("migrasi berhasil dijalankan")
}