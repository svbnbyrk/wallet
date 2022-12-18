package entity

type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUser(name, email string) User {
	return User{
		Name:name,
		Email:email,
	}
}