package routes

import( "github.com/gin-gonic/gin"
		"github.com/haewonito/gin-gorm-api-add-artists-many-many/controllers"
)

func AlbumSongRoute(router *gin.Engine) {
	router.GET("/albums/:id/songs", controller.GetSongByAlbumId)
}