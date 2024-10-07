package services

import (
	"context"
	"fmt"
	"github.com/zenorachi/youtube-task/internal/models"
	"github.com/zenorachi/youtube-task/internal/repository"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type IChannelsService interface {
	SetRepo(repo repository.IChannelRepository)
	GetChannels(ctx context.Context, data *models.ChannelRequest) ([]*models.Channel, error)
	InsertChannelsToDB(ctx context.Context, topic string, maxRes int64, lan *string) error
}

type channelsService struct {
	youtube *youtube.Service
	repo    repository.IChannelRepository
}

func NewChannelsService(ctx context.Context, apiKey string) (IChannelsService, error) {
	s, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}

	return &channelsService{
		youtube: s,
	}, nil
}

func (c *channelsService) SetRepo(repo repository.IChannelRepository) {
	c.repo = repo
}

// GetChannels - list of channels with ID and Title.
// Params:
// 1) topic - channel topic;
// 2) maxRes - counter of channels;
// 3) lan (optional) - channels language.
func (c *channelsService) GetChannels(ctx context.Context, data *models.ChannelRequest) ([]*models.Channel, error) {
	call := c.youtube.Search.List([]string{"id", "snippet"}).
		Context(ctx).
		Q(data.Topic).
		Type(models.YTTypeChannel).
		MaxResults(data.MaxRes)

	if lan := data.Lan; lan != nil {
		call.RelevanceLanguage(*lan)
	}

	response, err := call.Do()
	if err != nil {
		return nil, err
	}

	channels := make([]*models.Channel, 0)

	for _, item := range response.Items {
		if item.Id.Kind == fmt.Sprintf("youtube#%s", models.YTTypeChannel) {
			subsCall := c.youtube.Channels.List([]string{"statistics"}).Id(item.Id.ChannelId)
			subsResp, err := subsCall.Do()
			if err != nil {
				return nil, err
			}

			channels = append(channels, &models.Channel{
				ID:            item.Id.ChannelId,
				Topic:         data.Topic,
				Title:         item.Snippet.Title,
				Subscriptions: subsResp.Items[0].Statistics.SubscriberCount,
			})
		}
	}
	
	return channels, nil
}

func (c *channelsService) InsertChannelsToDB(ctx context.Context, topic string, maxRes int64, lan *string) error {
	channels, err := c.GetChannels(ctx, &models.ChannelRequest{
		Topic:  topic,
		MaxRes: maxRes,
		Lan:    lan,
	})
	if err != nil {
		return err
	}

	return c.repo.WithTransaction(func() error {
		for _, channel := range channels {
			if err = c.repo.InsertYTChannel(ctx, channel.ID, channel.Topic, channel.Title, channel.Subscriptions); err != nil {
				return err
			}
		}

		return nil
	})
}
