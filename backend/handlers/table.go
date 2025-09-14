package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func TableDataHandler(c *gin.Context) {
	if activeConn == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not connected"})
		return
	}

	table := c.Param("table")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "50"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	rows, err := activeConn.GetRows(table, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, rows)
}