package repository

import (
	"database/sql"
	"fmt"
	"github.com/gagansingh3785/go_authentication/constants"
	"github.com/gagansingh3785/go_authentication/database"
	"github.com/gagansingh3785/go_authentication/domain"
)

const baseUserQuery = "SELECT user_id, username, passwordhash, email, phone, role, salt FROM "

const findByIDUserQuery = baseUserQuery + constants.USER_TABLE + " WHERE user_id = $1"
const findByUsernameUserQuery = baseUserQuery + constants.USER_TABLE + " WHERE username = $1"
const findByEmailUserQuery = baseUserQuery + constants.USER_TABLE + " WHERE email = $1"
const findByEmailOrUsernameUserQuery = baseUserQuery + constants.USER_TABLE + " WHERE email = $1 OR username = $1"
const createNewUserQuery = "INSERT INTO " + constants.USER_TABLE + " (username, email, salt, phone, passwordhash, role) VALUES ($1, $2, $3, $4, $5, $6) RETURNING (user_id)"

func GetUserByUserID(userID string) (domain.User, error) {
	user := domain.User{}
	row := database.DBConn.QueryRow(findByIDUserQuery, userID)
	if err := row.Scan(&user.UUID, &user.Email, &user.PasswordHash, &user.Phone, &user.Role, &user.UUID, &user.Salt); err != nil {
		fmt.Println(err.Error())
		if err == sql.ErrNoRows {
			return user, constants.ErrSQLNoRows
		}
		return user, err
	}
	return user, nil
}

func GetUserByEmail(email string) (domain.User, error) {
	user := domain.User{}
	row := database.DBConn.QueryRow(findByEmailUserQuery, email)
	if err := row.Scan(&user.UUID, &user.Username, &user.PasswordHash, &user.Email, &user.Phone, &user.Role, &user.Salt); err != nil {
		fmt.Println(err.Error())
		if err == sql.ErrNoRows {
			return user, constants.ErrSQLNoRows
		}
		return user, err
	}
	return user, nil
}

func GetUserByEmailOrUsername(userIdentifier string) (domain.User, error) {
	user := domain.User{}
	row := database.DBConn.QueryRow(findByEmailOrUsernameUserQuery, userIdentifier)
	if err := row.Scan(&user.UUID, &user.Username, &user.PasswordHash, &user.Email, &user.Phone, &user.Role, &user.Salt); err != nil {
		fmt.Println(err.Error())
		if err == sql.ErrNoRows {
			return user, constants.ErrSQLNoRows
		}
		return user, err
	}
	return user, nil
}

func GetUserByUsername(username string) (domain.User, error) {
	user := domain.User{}
	row := database.DBConn.QueryRow(findByUsernameUserQuery, username)
	if err := row.Scan(&user.UUID, &user.Username, &user.PasswordHash, &user.Email, &user.Phone, &user.Role, &user.Salt); err != nil {
		fmt.Println(err.Error())
		if err == sql.ErrNoRows {
			return user, constants.ErrSQLNoRows
		}
		return user, err
	}
	return user, nil
}

func CreateNewUser(username, email, salt, phone, passwordHash string, role int) (domain.User, error) {
	user := domain.User{}
	row := database.DBConn.QueryRow(createNewUserQuery, username, email, salt, phone, passwordHash, role)
	if err := row.Scan(&user.UUID); err != nil {
		fmt.Println(err.Error())
		if err == sql.ErrNoRows {
			return user, constants.ErrSQLNoRows
		}
		return user, err
	}
	return user, nil
}
