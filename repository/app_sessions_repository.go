package repository

import (
	"database/sql"
	"fmt"
	"github.com/gagansingh3785/go_authentication/constants"
	"github.com/gagansingh3785/go_authentication/database"
	"github.com/gagansingh3785/go_authentication/domain"
	"time"
)

//queries

const findAllSessionsQuery = "SELECT * FROM " + constants.SESSIONS_TABLE
const findByIDSessionsQuery = "SELECT user_id, session_id FROM " + constants.SESSIONS_TABLE + " WHERE user_id = $1"
const createSessionQuery = "INSERT INTO " + constants.SESSIONS_TABLE + " (user_id, session_id) VALUES ($1, $2) RETURNING user_id, session_id"
const findByUserIDSessionQuery = "SELECT user_id, session_id FROM " + constants.SESSIONS_TABLE + "  WHERE user_id = $1"
const deleteSessionByUserIDQuery = "DELETE FROM " + constants.SESSIONS_TABLE + " WHERE user_id = $1"
const deleteSessionBySessionIDQuery = "DELETE FROM " + constants.SESSIONS_TABLE + " WHERE session_id = $1"
const clearOldSessionsQuery = "DELETE FROM " + constants.SESSIONS_TABLE + " WHERE $1 - created_time > '1 days 00:00:00'"

func CreateSession(userID, sessionID string) (domain.Session, error) {
	session := domain.Session{}
	row := database.DBConn.QueryRow(createSessionQuery, userID, sessionID)
	err := row.Scan(&session.UserID, &session.SessionID)
	if err != nil {
		if err == sql.ErrNoRows {
			return session, constants.ErrSQLNoRows
		}
		return session, err
	}
	return session, nil
}

func GetSessionFromUserID(userID string) (domain.Session, error) {
	session := domain.Session{}
	row := database.DBConn.QueryRow(findByUserIDSessionQuery, userID)
	err := row.Scan(&session.UserID, &session.SessionID)
	if err != nil {
		if err == sql.ErrNoRows {
			return session, constants.ErrSQLNoRows
		}
		return session, err
	}
	return session, nil
}

func DeleteSessionByUserID(userID string) error {
	row := database.DBConn.QueryRow(deleteSessionByUserIDQuery, userID)
	err := row.Scan()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return err
	}
	return nil
}

func DeleteSessionBySessionID(sessionID string) error {
	row := database.DBConn.QueryRow(deleteSessionBySessionIDQuery, sessionID)
	err := row.Scan()
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Logout &&&&&&&&", err.Error())
			return nil
		}
		return err
	}
	return nil
}

func ClearOldSessions(currentTime time.Time) error {
	rows, err := database.DBConn.Query(clearOldSessionsQuery, currentTime)
	if err != nil {
		fmt.Println("ClearSessionQuery: ", err.Error())
		return err
	}
	defer rows.Close()
	return nil
}
