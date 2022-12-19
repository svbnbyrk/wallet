package usecase_test

import (
	"context"
	"testing"

	"github.com/svbnbyrk/wallet/internal/entity"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/svbnbyrk/wallet/internal/usecase"
)

func wallet(t *testing.T) (*usecase.WalletUseCase, *MockWalletRepository) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	wallet := NewMockWalletRepository(mockCtl)

	walletUsecase := usecase.NewWalletUseCase(wallet)

	return walletUsecase, wallet
}

func TestWalletStore(t *testing.T) {
	t.Parallel()

	walletUc, walletRepo := wallet(t)

	tests := []test{
		{
			name: "result with no error",
			mock: func() {
				walletRepo.EXPECT().Store(context.Background(), entity.Wallet{}).Return(nil)
			},
			res: nil,
			err: nil,
		},
		{
			name: "result with error",
			mock: func() {
				walletRepo.EXPECT().Store(context.Background(), entity.Wallet{}).Return(errInternalServErr)
			},
			res: entity.Wallet{},
			err: errInternalServErr,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			err := walletUc.Store(context.Background(), entity.Wallet{})

			require.ErrorIs(t, err, tc.err)
		})
	}
}
