package v1handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/zenorachi/youtube-task/api/rest/apiutils"
	"github.com/zenorachi/youtube-task/internal/models"
	"github.com/zenorachi/youtube-task/internal/services"
	"github.com/zenorachi/youtube-task/internal/utils"
	"net/http"
)

type ChannelsHandler struct {
	srv services.IChannelsService
}

func NewChannelsHandler(srv services.IChannelsService) *ChannelsHandler {
	return &ChannelsHandler{
		srv: srv,
	}
}

type getChannelsReq struct {
	Topic      string  `form:"topic" binding:"required"`
	MaxResults int64   `form:"max_results" binding:"required"`
	Language   *string `form:"language"`
	Filename   string  `form:"filename"`
}

type getChannelsResp struct {
	ID            string `json:"id"`
	Topic         string `json:"topic"`
	Title         string `json:"title"`
	Subscriptions uint64 `json:"subscriptions"`
}

// GetChannels
//
//	@Summary	Get YouTube channels
//	@Tags		v1,channels
//	@Produce	json
//	@Param		topic		query		string	true	"channel topic"
//	@Param		max_results	query		int		true	"channels count"
//	@Param		language	query		string	false	"channels language"
//	@Param		filename	query		string	false	"csv filename (if want to save results to csv)"
//	@Success	200			{object}	[]getChannelsResp
//	@Failure	400			{object}	apiutils.ErrorResponse
//	@Failure	500			{object}	apiutils.ErrorResponse
//	@Router		/api/v1/channels [get]
func (h *ChannelsHandler) GetChannels(c *gin.Context) {
	var req getChannelsReq
	if err := c.BindQuery(&req); err != nil {
		apiutils.NotSuccessResponse(c, http.StatusBadRequest, err)
		return
	}

	channels, err := h.srv.GetChannels(c, &models.ChannelRequest{
		Topic:    req.Topic,
		MaxRes:   req.MaxResults,
		Lan:      req.Language,
		FileName: req.Filename,
	})
	if err != nil {
		apiutils.NotSuccessResponse(c, http.StatusInternalServerError, err)
		return
	}

	var res []getChannelsResp
	for _, channel := range channels {
		res = append(res, getChannelsResp{
			ID:            channel.ID,
			Topic:         channel.Topic,
			Title:         channel.Title,
			Subscriptions: channel.Subscriptions,
		})
	}

	if len(req.Filename) > 0 {
		if err = apiutils.WriteToCsv(
			c,
			req.Filename,
			utils.ConvertModelToStr2xSlice(channels),
			[]string{"id", "topic", "title", "subscriptions"}...,
		); err != nil {
			apiutils.NotSuccessResponse(c, http.StatusInternalServerError, err)
			return
		}
	} else {
		c.JSON(http.StatusOK, &res)
	}
}

type insertChannelsReq struct {
	Topic      string  `json:"topic" binding:"required"`
	MaxResults int64   `json:"max_results" binding:"required"`
	Language   *string `json:"language"`
}

// InsertChannels
//
//	@Summary	Insert YouTube channels to database
//	@Tags		v1,channels
//	@Produce	json
//	@Param		data	body		insertChannelsReq	true	"data for searching channels to insert"
//	@Success	201		{object}	nil
//	@Failure	400		{object}	apiutils.ErrorResponse
//	@Failure	500		{object}	apiutils.ErrorResponse
//	@Router		/api/v1/channels [post]
func (h *ChannelsHandler) InsertChannels(c *gin.Context) {
	var req insertChannelsReq
	if err := c.BindJSON(&req); err != nil {
		apiutils.NotSuccessResponse(c, http.StatusBadRequest, err)
		return
	}

	err := h.srv.InsertChannelsToDB(c, req.Topic, req.MaxResults, req.Language)
	if err != nil {
		apiutils.NotSuccessResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, nil)
}
