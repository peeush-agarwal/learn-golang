package main

import (
	"fmt"

	"example/restful-apis/cmd/blueprint/apis"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		v1.GET("/users/:id", apis.GetUser)
	}

	r.Run(fmt.Sprintf(":%v", 5000))
}
