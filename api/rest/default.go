package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Healthcheck
//
//	@Summary	Healthcheck route
//	@Tags		default
//	@Success	200	{object}	nil
//	@Router		/healthcheck [get]
func healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
