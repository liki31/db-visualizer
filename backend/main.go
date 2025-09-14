package main

import (
	"github.com/gin-gonic/gin"
	"github.com/liki31/db-visualizer/backend/handlers"
)

func main() {
	r := gin.Default()

	r.POST("/connect", handlers.ConnectHandler)
	r.GET("/schema", handlers.SchemaHandler)
	r.GET("/schema/:table", handlers.TableSchemaHandler)
	r.GET("/table/:table", handlers.TableDataHandler)

	r.Run(":8080")
}