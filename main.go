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
		db.Order("`order` asc").Find(&todos)
		c.JSON(200, todos)
	})

	r.POST("/api/todos", func(c *gin.Context) {
		var todo models.Todo
		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// Set the order to be the last item
		var count int64
		db.Model(&models.Todo{}).Count(&count)
		todo.Order = int(count)

		db.Create(&todo)
		c.JSON(201, todo)
	})

	r.PUT("/api/todos/reorder", func(c *gin.Context) {
		var reorderRequest struct {
			TodoIDs []uint `json:"todoIds"`
		}

		if err := c.ShouldBindJSON(&reorderRequest); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// Update the order of each todo
		for i, id := range reorderRequest.TodoIDs {
			db.Model(&models.Todo{}).Where("id = ?", id).Update("order", i)
		}

		c.JSON(200, gin.H{"message": "Todos reordered successfully"})
	})

	r.DELETE("/api/todos/:id", func(c *gin.Context) {
		id := c.Param("id")
		db.Delete(&models.Todo{}, id)
		c.JSON(200, gin.H{"message": "Todo deleted"})
	})

	r.Run(":8080")
}
