package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/stackpath/backend-developer-tests/rest-service/global"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/models/handlers"
)

// BackendRouter ... backend router to handle various APIs
func BackendRouter() (*gin.Engine, error) {
	restRouter := gin.Default()
	// configure cors as needed for FE/BE interactions: For now defaults

	configCors := cors.DefaultConfig()
	configCors.AllowAllOrigins = true
	configCors.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	restRouter.Use(cors.New(configCors))
	// help us recover from panic and log
	restRouter.Use(gin.Recovery())

	// TODO: inti route handlers

	//
	if !global.Options.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	
	//.....................................................................
	// healthcheck is need by Kubernetes to test readiness of containers
	//.....................................................................
	restRouter.GET("/healthz", handlers.HealthCheckHandler)
	version1 := restRouter.Group("/v1")
	{
		// this api should check if a query is made else return all
		version1.GET("/people", handlers.AllPeopleHandler)
		version1.GET("/people/:id", handlers.PersonWithIDHandler)
	}
	return restRouter, nil
}
