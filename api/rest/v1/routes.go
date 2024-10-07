package v1

import (
	"github.com/gin-gonic/gin"
	v1handlers "github.com/zenorachi/youtube-task/api/rest/v1/handlers"
)

func InitRoutes(
	router *gin.RouterGroup,
	channelsHandler *v1handlers.ChannelsHandler,
) {
	router.GET("/channels", channelsHandler.GetChannels)
	router.POST("/channels", channelsHandler.InsertChannels)
}
