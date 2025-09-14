package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SchemaHandler(c *gin.Context) {
	if activeConn == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not connected"})
		return
	}

	tables, err := activeConn.GetTables()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tables)
}

func TableSchemaHandler(c *gin.Context) {
	if activeConn == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not connected"})
		return
	}

	table := c.Param("table")
	cols, err := activeConn.GetColumns(table)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cols)
}