package main

import (
	"github.com/gin-gonic/gin"
	"github.com/haewonito/gin-gorm-api-add-artists-many-many/config"
	"github.com/haewonito/gin-gorm-api-add-artists-many-many/routes"
)

func main() {
	router := gin.New()
	
	config.Connect()
	routes.SongRoute(router)
	routes.AlbumRoute(router)

	router.Run(":8080")
}