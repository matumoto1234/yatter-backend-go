package usecase

// パッケージ(package 〇〇)を import
import (
	"context"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type Status interface {
	AddStatus(ctx context.Context, content string, account *object.Account) (*AddStatusDTO, error)
}

type status struct {
	db *sqlx.DB
	statusRepo repository.Status
}

type AddStatusDTO struct {
	Status *object.Status
}

var _ Status = (*status)(nil)

func NewStatus(db *sqlx.DB, statusRepo repository.Status) *status {
	return &status{
		db:          db,
		statusRepo: statusRepo,
	}
}

func (s *status) AddStatus(ctx context.Context, content string, account *object.Account) (*AddStatusDTO, error) {
	fmt.Println("hello")
	fmt.Println("content",content)
	fmt.Println("account",account)
	// objectパッケージのNewStatus関数を呼び出し、ステータスを作成
	status := object.NewStatus(content, account)
	fmt.Println("status",status)
	
	// トランザクションを開始する関数
	tx, err := s.db.Beginx()
	if err != nil {
		return nil, err
	}

	defer func() {
		// recover() は panic() が発生した場合に、panic()をキャッチする
		// panic() が発生した場合、トランザクションをロールバックする
		if err := recover(); err != nil {
			tx.Rollback()
		}

		tx.Commit()
	}()

	if err := s.statusRepo.AddStatus(ctx, tx, status); err != nil {
		return nil, err
	}

	return &AddStatusDTO{Status: status}, nil

}