package main

import (
	"github.com/YogeLiu/medical/api"
	_ "github.com/YogeLiu/medical/log"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()
	engine.Use(gin.Recovery())
	router := engine.Group("/medical")
	api.Router(router)
	engine.Run(":8081")
}
