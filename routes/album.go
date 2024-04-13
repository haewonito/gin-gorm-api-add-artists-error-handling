package routes

import( "github.com/gin-gonic/gin"
		"github.com/haewonito/gin-gorm-api-add-artists-many-many/controllers"
)

func AlbumRoute(router *gin.Engine) {
	router.GET("/albums", controller.GetAlbums)
	router.POST("/albums", controller.CreateAlbum)
	router.GET("/albums/:id", controller.GetAlbumById)
	router.DELETE("/albums/:id", controller.DeleteAlbum)
	router.PATCH("/albums/:id", controller.UpdateAlbum)
}