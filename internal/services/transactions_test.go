package services_test

import (
	"testing"
)

func TestTransactionService(t *testing.T) {
	// assert := assert.New(t)

	// ctrl := gomock.NewController(t)
	// defer ctrl.Finish()

	// ctx := t.Context()
	// db := mock_db.NewMockDB(ctrl)

	// repo := mock_repositories.NewMockTransactions(ctrl)
	// s := services.NewTransactionService(db, repo)

	// TODO: due to lack of time, I couldn't fininsh these tests on time.
	// should be done later.

	// 	t.Run("should init transactions service", func(t *testing.T) {
	// 		assert.NotNil(s)
	// 	})

	// t.Run("should create transactions", func(t *testing.T) {
	// 	d := &dtos.CreateAccount{DocumentNumber: "1"}
	// 	dbdata := &models.Account{DocumentNumber: d.DocumentNumber}

	// 	tx := mock_db.NewMockDriverTransaction(ctrl)
	// 	tx.EXPECT().Commit(gomock.Any()).Return(nil)
	// 	db.EXPECT().BeginTx(gomock.Any()).Return(tx, nil)

	// 	repo.EXPECT().Create(gomock.Any(), gomock.Any(), dbdata).Return(nil)

	// 	err := s.Create(ctx, d)
	// 	assert.NoError(err)
	// })

	// t.Run("should get account", func(t *testing.T) {
	// 	dbdata := &models.Account{Id: 1, DocumentNumber: "1"}
	// 	expected := &dtos.Account{Id: dbdata.Id, DocumentNumber: dbdata.DocumentNumber}

	// 	repo.EXPECT().Get(gomock.Any(), gomock.Any(), dbdata.Id).Return(dbdata, nil)

	// 	result, err := s.Get(ctx, dbdata.Id)
	// 	assert.NoError(err)
	// 	assert.Equal(expected, result)
	// })

	// t.Run("should not get account", func(t *testing.T) {
	// 	repo.EXPECT().Get(gomock.Any(), gomock.Any(), int64(1)).Return(nil, errors.New("not found"))

	// 	_, err := s.Get(ctx, int64(1))
	// 	assert.Error(err)
	// })

	// t.Run("should fail due to commit failure", func(t *testing.T) {
	// 	d := &dtos.CreateAccount{DocumentNumber: "1"}
	// 	dbdata := &models.Account{DocumentNumber: d.DocumentNumber}

	// 	tx := mock_db.NewMockDriverTransaction(ctrl)
	// 	tx.EXPECT().Commit(gomock.Any()).Return(errors.New("err"))
	// 	db.EXPECT().BeginTx(gomock.Any()).Return(tx, nil)

	// 	repo.EXPECT().Create(gomock.Any(), gomock.Any(), dbdata).Return(nil)

	// 	err := s.Create(ctx, d)
	// 	assert.Error(err)
	// })

	// t.Run("should fail due to panic", func(t *testing.T) {
	// 	d := &dtos.CreateAccount{DocumentNumber: "1"}
	// 	dbdata := &models.Account{DocumentNumber: d.DocumentNumber}

	// 	tx := mock_db.NewMockDriverTransaction(ctrl)

	// 	tx.EXPECT().Commit(gomock.Any()).Return(errors.New("err"))
	// 	tx.EXPECT().Rollback(gomock.Any()).Return(errors.New("err"))
	// 	db.EXPECT().BeginTx(gomock.Any()).Return(tx, nil)

	// 	repo.EXPECT().Create(gomock.Any(), gomock.Any(), dbdata).
	// 		DoAndReturn(func(ctx, conn, account any) error {
	// 			panic("some unexpected error ocurred, and transaction must rollback")
	// 		})

	// 	err := s.Create(ctx, d)
	// 	assert.Error(err)
	// })
}
