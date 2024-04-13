package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/haewonito/gin-gorm-api-add-artists-many-many/models"
	"github.com/haewonito/gin-gorm-api-add-artists-many-many/config"
)

func GetSongs(c *gin.Context) {
	songs := []models.Song{}
	config.DB.Find(&songs)
	c.JSON(200, &songs)
}

func CreateSong(c *gin.Context) {
	var song models.Song
	c.BindJSON(&song)
	config.DB.Create(&song)
	c.JSON(200, &song)
}

func GetSongById(c *gin.Context) {
	var song models.Song
	config.DB.First(&song, c.Param("id"))
	c.JSON(200, &song)
}

func DeleteSong(c *gin.Context) {
	var song models.Song
	config.DB.Where("id = ?", c.Param("id")).Delete(&song)
	c.JSON(200, &song)
}

func UpdateSong(c *gin.Context) {
	var song models.Song
	config.DB.Where("id = ?", c.Param("id")).First(&song)
	c.BindJSON(&song)
	config.DB.Save(&song)
	c.JSON(200, &song)
}