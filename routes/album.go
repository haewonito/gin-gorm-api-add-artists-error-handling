package routes

import( "github.com/gin-gonic/gin"
		"github.com/haewonito/gin-gorm-api-add-artists-many-many/controllers"
)

func AlbumRoute(router *gin.Engine) {
	router.GET("/albums", controller.GetAlbums)
}