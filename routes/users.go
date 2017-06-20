package routes

import (
  "gopkg.in/gin-gonic/gin.v1"
  "net/http"
)

func User(c *gin.Context) {
  name := c.Param("name")
  alt := c.DefaultQuery("alt", "Guest")

  c.String(http.StatusOK, "Hello %s (aka %s)", name, alt)
}

func UserAction(c *gin.Context) {
  name := c.Param("name")
  action := c.Param("action")
  message := name + " is " + action
  c.String(http.StatusOK, message)
}