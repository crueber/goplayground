package routes

import (
  "gopkg.in/gin-gonic/gin.v1"
  "net/http"
  // "encoding/json"
)

func socketUp(c *gin.Context) {
  c.String(http.StatusOK, "Here we are.")
}

func Socket(router *gin.Engine) {
  router.GET("/socket", socketUp)
}
