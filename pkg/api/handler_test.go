package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestRouter() *gin.Engine {
	dsn := "user=postgres password=1234 dbname=command host=localhost port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	SetupRoutes(r)
	return r
}

func TestCreateCommand(t *testing.T) {
	r := setupTestRouter()

	jsonData := []byte(`{"Name": "test_command", "Script": "echo Hello, World!"}`)
	req, err := http.NewRequest("POST", "/create-command", bytes.NewBuffer(jsonData))
	assert.NoError(t, err)

	resp := httptest.NewRecorder()

	r.POST("/create-command", CreateCommand)
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var result map[string]interface{}
	err = json.Unmarshal(resp.Body.Bytes(), &result)
	assert.NoError(t, err)
	assert.Equal(t, "test_command", result["Name"])
	assert.Equal(t, "Hello, World!\n", result["Output"])
}

func TestGetCommands(t *testing.T) {
	r := setupTestRouter()
	req, _ := http.NewRequest("GET", "/commands", nil)

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected status OK, got %v", resp.Code)
	}
}

func TestGetCommand(t *testing.T) {
	r := setupTestRouter()
	req, _ := http.NewRequest("GET", "/commands/1", nil)

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected status OK, got %v", resp.Code)
	}
}
