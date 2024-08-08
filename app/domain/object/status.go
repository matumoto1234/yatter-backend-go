package object

import (
	"time"
)

type Status struct {
	ID        int       `json:"id,omitempty"`
	AccountID  int      `json:"account,omitempty" db:"account"`
	URL       *string   `json:"url,omitempty" db:"url"`
	Content   string    `json:"status"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
}

func NewStatus(content string, account *Account) *Status {
	return &Status{
		Content:   content,
		AccountID:   int(account.ID),
		CreatedAt: time.Now(),
	}
}
