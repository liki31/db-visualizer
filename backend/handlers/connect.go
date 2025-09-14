package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/liki31/db-visualizer/backend/db"
)

var activeConn *db.DBConnection

func ConnectHandler(c *gin.Context) {
	type ConnReq struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Database string `json:"database"`
	}

	var req ConnReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	conn, err := db.ConnectPostgres(req.Host, req.Port, req.User, req.Password, req.Database, )
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	activeConn = conn
	c.JSON(http.StatusOK, gin.H{"message": "Connected successfully"})
}