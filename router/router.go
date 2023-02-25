package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/jary-287/gopass-svc-api/router/v1"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	api := r.Group("/api/v1")
	{
		//test
		api.GET("/service", v1.GetAllSvc)
		api.POST("/service", v1.AddSvc)
		api.DELETE("/service", v1.DeleteSvc)
		api.PUT("/service", v1.UpdateSvc)
		api.GET("/service/:id", v1.GetSvcByID)
	}
	return r

}
