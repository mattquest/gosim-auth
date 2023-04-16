package auth

type User struct {
	ID       string
	Email    string
	Password string
}

type UserStore interface {
	CreateUser(user *User) error
	GetUserByEmail(email string) (*User, error)
	UserExists(email string) (bool, error)
}
