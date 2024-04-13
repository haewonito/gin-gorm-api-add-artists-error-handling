package config
import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/haewonito/gin-gorm-api-add-artists-many-many/models"
  )

var DB *gorm.DB
func Connect() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/postgres"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Album{}, &models.Song{})
	DB = db
}