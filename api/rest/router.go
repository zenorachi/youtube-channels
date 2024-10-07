package rest

import (
	"github.com/zenorachi/youtube-task/api/rest/middlewares"
	v1 "github.com/zenorachi/youtube-task/api/rest/v1"
	v1handlers "github.com/zenorachi/youtube-task/api/rest/v1/handlers"
	_ "github.com/zenorachi/youtube-task/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(channelsHandler *v1handlers.ChannelsHandler) *gin.Engine {
	router := gin.New()
	router.Use(cors.Default())
	router.Use(middlewares.LogsMiddleware)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/healthcheck", healthcheck)

	v1Api := router.Group("/api/v1")
	v1.InitRoutes(v1Api, channelsHandler)

	return router
}
