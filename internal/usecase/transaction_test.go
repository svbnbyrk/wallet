package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/svbnbyrk/wallet/internal/entity"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/svbnbyrk/wallet/internal/usecase"
)

var errInternalServErr = errors.New("internal server error")

type test struct {
	name string
	mock func()
	res  interface{}
	err  error
}

func transaction(t *testing.T) (*usecase.TransactionUseCase, *MockTransactionRepository, *MockWalletRepository) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	transaction := NewMockTransactionRepository(mockCtl)
	wallet := NewMockWalletRepository(mockCtl)
	exchange := NewMockExchangeRepository(mockCtl)

	translation := usecase.NewTransactionUseCase(transaction, wallet, exchange)

	return translation, transaction, wallet
}

func TestHistory(t *testing.T) {
	t.Parallel()

	transaction, repo, _ := transaction(t)

	tests := []test{
		{
			name: "empty result",
			mock: func() {
				repo.EXPECT().GetHistory(context.Background()).Return(nil, nil)
			},
			res: []entity.Transaction(nil),
			err: nil,
		},
		{
			name: "result with error",
			mock: func() {
				repo.EXPECT().GetHistory(context.Background()).Return(nil, errInternalServErr)
			},
			res: []entity.Transaction(nil),
			err: errInternalServErr,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			res, err := transaction.History(context.Background())

			require.Equal(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}
