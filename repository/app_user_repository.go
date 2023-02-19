package repository

import (
	"database/sql"
	"fmt"
	"github.com/gagansingh3785/go_authentication/constants"
	"github.com/gagansingh3785/go_authentication/database"
	"github.com/gagansingh3785/go_authentication/domain"
)

const findByIDUserQuery = "SELECT user_id, username, password_hash, email, phone, salt, role FROM " + constants.USER_TABLE + " WHERE user_id = $1"
const findByUsernameUserQuery = "SELECT user_id, username, password_hash, email, phone, salt, role FROM " + constants.USER_TABLE + " WHERE username = $1"
const findByEmailUserQuery = "SELECT user_id, username, password_hash, email, phone, salt, role FROM " + constants.USER_TABLE + " WHERE email = $1"
const findByEmailOrUsernameUserQuery = "SELECT user_id, username, password_hash, email, salt, phone, role FROM " + constants.USER_TABLE + " WHERE email = $1 OR username = $1"
const createNewUserQuery = "INSERT INTO " + constants.USER_TABLE + " (username, email, salt, phone, password_hash, role) VALUES ($1, $2, $3, $4, $5, $6) RETURNING (user_id, username, email, salt, phone, password_hash, role)"

func GetUserByUserID(userID string) (domain.User, error) {
	user := domain.User{}
	row := database.DBConn.QueryRow(findByIDUserQuery, userID)
	if err := row.Scan(user); err != nil {
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
	if err := row.Scan(user); err != nil {
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
	if err := row.Scan(user); err != nil {
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
	if err := row.Scan(user); err != nil {
		fmt.Println(err.Error())
		if err == sql.ErrNoRows {
			return user, constants.ErrSQLNoRows
		}
		return user, err
	}
	return user, nil
}

func CreateNewUser(username, email, salt, phone, passwordHash string, role int64) (domain.User, error) {
	user := domain.User{}
	row := database.DBConn.QueryRow(createNewUserQuery, username, email, salt, phone, passwordHash, role)
	if err := row.Scan(user); err != nil {
		fmt.Println(err.Error())
		if err == sql.ErrNoRows {
			return user, constants.ErrSQLNoRows
		}
		return user, err
	}
	return user, nil
}
