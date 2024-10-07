package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/zenorachi/youtube-task/pkg/database/postgres"
)

type IChannelRepository interface {
	postgres.ITransactor
	InsertYTChannel(ctx context.Context, id, name, topic string, subscriptions uint64) error
}

type channelRepository struct {
	*postgres.Runner
}

func NewChannelRepository(db *sqlx.DB) IChannelRepository {
	return &channelRepository{
		Runner: postgres.NewRunner(db),
	}
}

const insertYTChannelQuery = `
	INSERT INTO yt_channels (channel_id, topic, title, subscriptions)
	VALUES ($1, $2, $3, $4)
	ON CONFLICT (channel_id) DO UPDATE
	SET topic = excluded.topic,
	    title = excluded.title,
	    subscriptions = excluded.subscriptions,
		updated_at = NOW()
`

func (c *channelRepository) InsertYTChannel(ctx context.Context, id, name, topic string, subscriptions uint64) error {
	_, err := c.Exec.ExecContext(ctx, insertYTChannelQuery, id, name, topic, subscriptions)
	if err != nil {
		return err
	}

	return nil
}
