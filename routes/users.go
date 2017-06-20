package routes

import (
  "gopkg.in/gin-gonic/gin.v1"
  "net/http"
)

func user(c *gin.Context) {
  name := c.Param("name")
  alt := c.DefaultQuery("alt", "Guest")

  c.String(http.StatusOK, "Hello %s (aka %s)", name, alt)
}

func userAction(c *gin.Context) {
  name := c.Param("name")
  action := c.Param("action")
  message := name + " is " + action
  c.String(http.StatusOK, message)
}

func Users(router *gin.Engine) {
  router.GET("/user/:name", user)
  router.GET("/user/:name/*action", userAction)
}
