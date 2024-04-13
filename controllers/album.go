package controller
//https://gorm.io/docs/query.html
//https://gorm.io/docs/create.html
import (
	"github.com/gin-gonic/gin"
	// "github.com/haewonito/gin-gorm-api-add-artists-many-many/models"  //CHANGE GITHUB repo NAMES!
	// "github.com/haewonito/gin-gorm-api-add-artists-many-many/config"
)

func GetAlbums(c *gin.Context) {
	c.String(200, "Hello World!")
}