package routes

import (
  "gopkg.in/gin-gonic/gin.v1"
  "net/http"
  "encoding/json"
)

type User struct {
  Name string `json:"name"`
  Email string `json:"email"`
}

func Users(router *gin.Engine) {
  router.GET("/user/:name", getUser)
  router.GET("/user/:name/*action", userAction)
  router.POST("/user", create)
}

func getUser(c *gin.Context) {
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


func create(c *gin.Context) {
  var user User
  e := c.BindJSON(&user)
  if e != nil {
    c.Error(e)  
    return
  }
  
  r, e := json.Marshal(user)
  if e != nil {
    c.Error(e)
  }

  if len(c.Errors) == 0 {
    c.String(http.StatusOK, "Name: %s\n%s", user.Name, r)  
  } else {
    c.AbortWithStatus(400)
  }
}
