package usecase

import (
	"context"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type Timeline interface {
	GetPublic(ctx context.Context, onlyMedia bool, sinceId int, limit int) (*TimelineDto, error)
}

type TimelineDto struct {
	Timeline *object.Timeline
}

type timeline struct {
	db *sqlx.DB
	timelineRepo repository.Timeline
}

var _ Timeline = (*timeline)(nil)

func NewTimeline(db *sqlx.DB, timelineRepo repository.Timeline) *timeline {
	return &timeline{
		db: db,
		timelineRepo: timelineRepo,
	}
}

func (t *timeline) GetPublic(ctx context.Context, onlyMedia bool, sinceId int, limit int) (*TimelineDto, error) {
	timeline, err := t.timelineRepo.FindPublicTimelines(ctx, onlyMedia, sinceId, limit)
	if err != nil {
		return nil, err
	}

	return &TimelineDto{
		Timeline: timeline,
	}, nil
}
