package apiutils

import (
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zenorachi/youtube-task/internal/utils"
	"os"
)

type ErrorResponse struct {
	Detail string `json:"detail"`
}

func NotSuccessResponse(c *gin.Context, status int, err error) {
	utils.Silent(c.Error(err))

	c.AbortWithStatusJSON(
		status,
		ErrorResponse{
			Detail: err.Error(),
		},
	)
}

func WriteToCsv(c *gin.Context, name string, data [][]string, headers ...string) error {
	c.Writer.Header().Set("Content-Type", "text/csv; charset=utf-8")
	c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.csv", name))

	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer func() {
		utils.Silent(file.Close())
		utils.Silent(os.Remove(name))
	}()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	utils.Silent(writer.Write(headers))
	for _, row := range data {
		if err = writer.Write(row); err != nil {
			return err
		}
	}

	writer.Flush()

	csvData, err := os.ReadFile(name)
	if err != nil {
		return err
	}

	_, err = c.Writer.WriteString(string(csvData))
	return err
}
