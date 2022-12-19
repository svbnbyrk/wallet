package entity

import "context"

type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email" binding:"required,email"`
}

func NewUser(name, email string) User {
	return User{
		Name:  name,
		Email: email,
	}
}

// User Usecase
type UserUseCase interface {
	Store(ctx context.Context, u User) error
}

// User Repository
type UserRepository interface {
	Store(ctx context.Context, u User) error
	Update(ctx context.Context, u User) error
}
