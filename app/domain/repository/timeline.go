package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type Timeline interface {
	FindPublicTimelines(ctx context.Context, onlyMedia bool, sinceId int, limit int) (*object.Timeline, error)
}