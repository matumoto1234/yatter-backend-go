package dao

import (
	"context"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	timeline struct {
		db *sqlx.DB
	}
)

var _ repository.Timeline = (*timeline)(nil)

func NewTimeline(db *sqlx.DB) *timeline {
	return &timeline{db: db}
}

func (t *timeline) FindPublicTimelines(ctx context.Context, onlyMedia bool, sinceId int, limit int) (*object.Timeline, error) {
	timeline := new(object.Timeline)
	query := "select * from status where id > ? order by id desc limit ?"
	err := t.db.SelectContext(ctx, &timeline.Timeline, query, sinceId, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to select timeline: %w", err)
	}

	return timeline, nil
}