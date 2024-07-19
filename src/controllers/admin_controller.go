package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ProfileAdmin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ok": "ok"})
}
