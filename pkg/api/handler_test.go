package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
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
	var jsonStr = []byte(`{"name":"test_command2","script":"echo 'Hello, World!'"}`)
	req, _ := http.NewRequest("POST", "/commands", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected status OK, got %v", resp.Code)
	}
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
