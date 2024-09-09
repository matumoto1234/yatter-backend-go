package usecase

import (
	"context"
	"testing"
	"yatter-backend-go/app/domain/object"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

type MockStatusRepo struct {
	mockCreateFunc func(ctx context.Context, tx *sqlx.Tx ,status *object.Status) error
}

func (m *MockStatusRepo) AddStatus(ctx context.Context, tx *sqlx.Tx, status *object.Status) error {
	return m.mockCreateFunc(ctx, tx, status)
}

func (m *MockStatusRepo) FindByID(ctx context .Context, id int) (*object.Status, error) {
	return nil, nil
}

func (m *MockStatusRepo) FindAccountByID(ctx context.Context, id int) (*object.Account, error) {
	return nil, nil
}

func TestStatusUsecase_AddStatus(t *testing.T) {
	ctx := context.Background()

	t.Run("正常系 : Status情報を正常に返すこと", func(t *testing.T){
		content := "test content"
		account := object.Account{
			ID: 1,
			Username: "testtest",
		}

		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}

		sqlxDB := sqlx.NewDb(db, "sqlmock")


		mock.ExpectBegin()

		mockRepo := MockStatusRepo{
			mockCreateFunc: func(ctx context.Context, tx *sqlx.Tx,status *object.Status) error {
				return nil
			},
		}

		sut := NewStatus(sqlxDB, &mockRepo)

		got , err := sut.AddStatus(ctx, content, &account)
		assert.NoError(t, err)
		want := object.NewStatus(content, &account)

		assert.Equal(t, want.AccountID, got.Status.AccountID)
		assert.Equal(t, want.Content, got.Status.Content)
		
		
	})

	t.Run("異常系 : StatusRepository.AddStatus() でエラーが発生した場合、エラーを返すこと", func(t *testing.T){
		// todo: 異常系のテストを書く
	})
}