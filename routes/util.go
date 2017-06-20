package routes

import (
  "gopkg.in/gin-gonic/gin.v1"
  "net/http"
)

  // gin.H is a shortcut for map[string]interface{}
func SomeJSON(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
}
