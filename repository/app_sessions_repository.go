package repository

import (
	"database/sql"
	"github.com/gagansingh3785/go_authentication/constants"
	"github.com/gagansingh3785/go_authentication/database"
	"github.com/gagansingh3785/go_authentication/domain"
)

//queries

const findAllSessionsQuery = "SELECT * FROM " + constants.SESSIONS_TABLE
const findByIDSessionsQuery = "SELECT user_id, session_id FROM " + constants.SESSIONS_TABLE + " WHERE user_id = $1"
const createSessionQuery = "INSERT INTO " + constants.SESSIONS_TABLE + " (user_id, session_id) VALUES ($1, $2) RETURNING user_id, session_id"
const findByUserIDSessionQuery = "INSERT INTO " + constants.SESSIONS_TABLE + " (user_id, session_id) WHERE user_id = $1"

func CreateSession(userID, sessionID string) (domain.Session, error) {
	session := domain.Session{}
	row := database.DBConn.QueryRow(createSessionQuery, userID, sessionID)
	err := row.Scan(session)
	if err != nil {
		if err == sql.ErrNoRows {
			return session, constants.ErrSQLNoRows
		}
		return session, err
	}
	return session, nil
}

func GetSessionFromUserID(username string) (domain.Session, error) {
	session := domain.Session{}
	row := database.DBConn.QueryRow(findByUserIDSessionQuery)
	err := row.Scan(session)
	if err != nil {
		if err == sql.ErrNoRows {
			return session, constants.ErrSQLNoRows
		}
		return session, err
	}
	return session, nil
}
