package domain

type User struct {
	Username string `db:"username"`
	Password string `db:"password"`
	Email    string `db:"email"`
	Phone    string `db:"phone"`
}

func (u User) GetUsername() string {
	return u.Username
}

func (u User) GetPassword() string {
	return u.Password
}

func (u User) GetEmail() string {
	return u.Email
}

func (u User) GetPhone() string {
	return u.Phone
}

func (u User) EncryptPassword() string {
	return ""
}
