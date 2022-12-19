package usecase_test

import (
	"context"
	"testing"

	"github.com/svbnbyrk/wallet/internal/entity"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/svbnbyrk/wallet/internal/usecase"
)

func User(t *testing.T) (*usecase.UserUseCase, *MockUserRepository) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	user := NewMockUserRepository(mockCtl)

	userUsecase := usecase.NewUserUseCase(user)

	return userUsecase, user
}

func TestUserStore(t *testing.T) {
	t.Parallel()

	userUc, userRepo := User(t)

	tests := []test{
		{
			name: "result with no error",
			mock: func() {
				userRepo.EXPECT().Store(context.Background(), entity.User{}).Return(nil)
			},
			res: nil,
			err: nil,
		},
		{
			name: "result with error",
			mock: func() {
				userRepo.EXPECT().Store(context.Background(), entity.User{}).Return(errInternalServErr)
			},
			res: entity.User{},
			err: errInternalServErr,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			err := userUc.Store(context.Background(), entity.User{})

			require.ErrorIs(t, err, tc.err)
		})
	}
}
