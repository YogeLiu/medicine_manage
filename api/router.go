package api

import (
	"github.com/gin-gonic/gin"
)

func Router(router *gin.RouterGroup) {
	v1 := router.Group("/v1")
	detail := v1.Group("/medicine")
	{
		endpoint := NewDetailBackend()
		detail.POST("/add", endpoint.AddDetail)
		detail.GET("/list", endpoint.ListDetail)
		detail.PUT("/update", endpoint.AddDetail)
		detail.DELETE("/delete", endpoint.DeleteDetail)
	}
}
