package api

import (
	"PGStart/pkg/models"

	"bytes"
	"os/exec"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateCommand(c *gin.Context) {
	var cmd models.Command
	if err := c.ShouldBindJSON(&cmd); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmdExecution := exec.Command("bash", "-c", cmd.Script)
	cmdExecution.Stdout = &out
	cmdExecution.Stderr = &stderr
	err := cmdExecution.Run()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to execute command", "stderr": stderr.String()})
		return
	}

	cmd.Output = out.String()

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
