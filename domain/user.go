package domain

type User struct {
	UUID         string `db:"uuid"`
	Username     string `db:"username"`
	PasswordHash string `db:"password_hash"`
	Email        string `db:"email"`
	Phone        string `db:"phone"`
	Salt         string `db:"salt"`
	Role         int64  `db:"role"`
}

func (u User) GetUsername() string {
	return u.Username
}

func (u User) GetPassword() string {
	return u.PasswordHash
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
