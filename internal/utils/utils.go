package utils

import (
	"github.com/zenorachi/youtube-task/internal/models"
	"strconv"
)

func Silent(...interface{}) {}

// ConvertModelToStr2xSlice TODO: use reflexion to use any models
func ConvertModelToStr2xSlice(data []*models.Channel) [][]string {
	var result [][]string

	for _, v := range data {
		result = append(result, []string{v.ID, v.Topic, v.Title, strconv.FormatUint(v.Subscriptions, 10)})
	}

	return result
}
