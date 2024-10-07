package models

const (
	YTTypeChannel = "channel"
)

type ChannelRequest struct {
	Topic    string
	MaxRes   int64
	Lan      *string
	FileName string
}

type Channel struct {
	ID            string `json:"id"`
	Topic         string `json:"topic"`
	Title         string `json:"title"`
	Subscriptions uint64 `json:"subscriptions"`
}
