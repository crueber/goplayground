package main

import (
  "gopkg.in/gin-gonic/gin.v1"
  "github.com/crueber/goplayground/routes"
)

func someGet(c *gin.Context) {
  c.String(200, "It worked.")
}

func main() {
  router := gin.Default()

  routes.Users(router)
  routes.Util(router)
  routes.Ping(router)
  router.GET("/someGet", someGet)

  router.Run(":8080")
}