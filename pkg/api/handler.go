package api

import (
	"PGStart/pkg/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateCommand(c *gin.Context) {
	var cmd models.Command
	if err := c.ShouldBindJSON(&cmd); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Create(&cmd).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, cmd)
}

func GetCommands(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var cmds []models.Command
	if err := db.Find(&cmds).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, cmds)
}

func GetCommand(c *gin.Context) {
	id := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)
	var cmd models.Command
	if err := db.First(&cmd, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Command not found"})
		return
	}

	c.JSON(200, cmd)
}
