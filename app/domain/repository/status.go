package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"

	"github.com/jmoiron/sqlx"
)

type Status interface {
	AddStatus(ctx context.Context, tx *sqlx.Tx, status *object.Status) error
	FindByID(ctx context.Context, id int) (*object.Status, error)
	FindAccountByID(ctx context.Context, id int) (*object.Account, error)
}