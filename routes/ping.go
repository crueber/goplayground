package routes

import (
  "gopkg.in/gin-gonic/gin.v1"
)

func ping(c *gin.Context) {
  c.JSON(200, gin.H{
    "message": "pong",
  })
}

func Pinger(router *gin.Engine) {
  router.GET("/ping", ping)  
}
