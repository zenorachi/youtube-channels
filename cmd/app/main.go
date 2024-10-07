package main

import (
	"github.com/zenorachi/youtube-task/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		panic(err)
	}
}
