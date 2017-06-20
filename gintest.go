package main

import (
  "gopkg.in/gin-gonic/gin.v1"
  "github.com/crueber/goplayground/routes"
)

// func main() {
//   r := gin.Default()
//   r.GET("/ping", func(c *gin.Context) {
//     c.JSON(200, gin.H{
//       "message": "pong",
//     })
//   })
//   r.Run() // listen and serve on 0.0.0.0:8080
// }

// func main() {
//   // Creates a gin router with default middleware:
//   // logger and recovery (crash-free) middleware
//   router := gin.Default()

//   router.GET("/someGet", getting)
//   router.POST("/somePost", posting)
//   router.PUT("/somePut", putting)
//   router.DELETE("/someDelete", deleting)
//   router.PATCH("/somePatch", patching)
//   router.HEAD("/someHead", head)
//   router.OPTIONS("/someOptions", options)

//   // By default it serves on :8080 unless a
//   // PORT environment variable was defined.
//   router.Run()
//   // router.Run(":3000") for a hard coded port
// }

func someGet(c *gin.Context) {
  c.String(200, "It worked.")
}

func main() {
  router := gin.Default()

  router.GET("/user/:name", routes.User)
  router.GET("/user/:name/*action", routes.UserAction)
  router.GET("/someJSON", routes.SomeJSON)
  router.GET("/someGet", someGet)

  router.Run(":8080")
}