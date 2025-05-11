package main

import (
	"log"
	"todos-app/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("todos.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate the schema
	db.AutoMigrate(&models.Todo{})

	r := gin.Default()

	// CORS configuration
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"} // React dev server
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Type"}
	r.Use(cors.New(config))

	// Routes
	r.GET("/api/todos", func(c *gin.Context) {
		var todos []models.Todo
		db.Find(&todos)
		c.JSON(200, todos)
	})

	r.POST("/api/todos", func(c *gin.Context) {
		var todo models.Todo
		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		db.Create(&todo)
		c.JSON(201, todo)
	})

	r.DELETE("/api/todos/:id", func(c *gin.Context) {
		id := c.Param("id")
		db.Delete(&models.Todo{}, id)
		c.JSON(200, gin.H{"message": "Todo deleted"})
	})

	r.Run(":8080")
}
