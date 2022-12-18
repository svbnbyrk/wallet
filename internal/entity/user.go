package entity

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
