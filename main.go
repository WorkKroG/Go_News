package main

import (
	"github.com/gin-gonic/gin"

	"./app"
)

func main() {
	r := gin.Default()

	app.MakeRoutes(r)
	app.Connect(r)

	r.Run(":8080")
}
