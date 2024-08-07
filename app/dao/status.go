package dao

import (
	"context"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Status
	status struct {
		// sqlx.DB: database/sql.DBをラップし、より簡単に使えるようにしたもの
		db *sqlx.DB
	}
)

var _ repository.Status = (*status)(nil)

// NewStatus : Create status repository
func NewStatus(db *sqlx.DB) *status {
	return &status{db: db}
}

// AddStatus : データベースにステータスを追加
func (s *status) AddStatus(ctx context.Context, tx *sqlx.Tx, status *object.Status)error{
	_, err := s.db.ExecContext((ctx), "insert into status (account, content, created_at) values (?, ?, ?)",
		status.Account, status.Content, status.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to insert status: %w", err)
	}

	return nil
}