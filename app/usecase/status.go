package usecase

// パッケージ(package 〇〇)を import
import (
	"context"
	"time"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type Status interface {
	AddStatus(ctx context.Context, content string, account *object.Account) (*AddStatusDTO, error)
	Get(ctx context.Context, id int) (*GetStatusWithAccount, error)
	
}

type status struct {
	db *sqlx.DB
	statusRepo repository.Status
}

// AddStatusDTO : ステータスを追加する際のデータ転送オブジェクトの構造体
// この実装だと、ドメイン層が変更した場合に、この構造体( AddStatus の返り値)も変更する必要がある
// なので、ドメイン層が変更しても AddStatus の返り値を変更しなくてもいいように、useCase 用の型を定義するのが良い
type AddStatusDTO struct {
	Status *object.Status
}

// useCase 用の型を定義する
// type AddStatusUseCaseDTO struct {
// 	ID        int       
// 	Account   *Account   
// 	URL       *string   
// 	Content   string    
// 	CreatedAt time.Time 
// }

type GetStatusWithAccount struct {
	ID int
	Account *object.Account
	URL *string
	Content string
	CreatedAt time.Time
}

var _ Status = (*status)(nil)

func NewStatus(db *sqlx.DB, statusRepo repository.Status) *status {
	return &status{
		db:          db,
		statusRepo: statusRepo,
	}
}

func (s *status) AddStatus(ctx context.Context, content string, account *object.Account) (*AddStatusDTO, error) {
	// objectパッケージのNewStatus関数を呼び出し、ステータスを作成
	status := object.NewStatus(content, account)
	
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

func (s *status) Get(ctx context.Context, id int) (*GetStatusWithAccount, error) {
	status, err := s.statusRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	accountId := status.AccountID

	account, err := s.statusRepo.FindAccountByID(ctx, accountId)
	if err != nil {
		return nil, err
	}

	return &GetStatusWithAccount{
		ID: status.ID,
		Account: account,
		URL: status.URL,
		Content: status.Content,
		CreatedAt: status.CreatedAt,
	}, nil
}