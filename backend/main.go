package main

import (
	"github.com/gin-gonic/gin"
	"github.com/liki31/db-visualizer/backend/handlers"
	"github.com/gin-contrib/cors"
	"time"
)

func main() {
    r := gin.Default()

    // Allow frontend requests
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"}, // Vite frontend
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
    }))

    r.POST("/connect", handlers.ConnectHandler)
    r.GET("/schema", handlers.SchemaHandler)
    r.GET("/schema/:table", handlers.TableSchemaHandler)
    r.GET("/table/:table", handlers.TableDataHandler)

    r.Run(":8080")
}