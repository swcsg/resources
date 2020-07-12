package router

import (
	"github.com/gin-gonic/gin"
	"resource/pkg/controller/app_dependence"
	"resource/pkg/controller/crd_dependence"
)

func RegisterRoutes() *gin.Engine {

	router := gin.New()

	r := router.Group("/api/v1/resources")
	{
		r.POST("/deployment/dependence", app_dependence.AppDependence)
		r.POST("/statefulset/dependence",crd_dependence.CRDDependence)
	}
	return router
}
