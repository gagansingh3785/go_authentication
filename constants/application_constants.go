package constants

const (
	// general constants

	// app constants
	PORT                = "3002"
	AUTHENTICATION_HOST = ""

	//database constants
	DB_PORT = "5432"
	DB_HOST = "localhost"
	//DB_HOST     = "10.128.0.4"
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "go_authentication"
	DB_SSLMODE  = "disable"

	//database table constants
	USER_TABLE     = "app_user"
	SESSIONS_TABLE = "app_sessions"

	//sendgrid constants
	SENDGRID_API_HOST     = "SENDGRID_API_HOST"
	SENDGRID_API_ENDPOINT = "SENDGRID_API_ENDPOINT"
	SENDGRID_API_KEY      = "SENDGRID_API_KEY"

	//Request Headers
	SESSION_COOKIE = "Session-Cookie"
	CONTENT_TYPE   = "application/json"

	//miscellaneous constants
	EMPTY_STRING = ""

	// keys
	MY_SECRET = "MY_SECRET"
)
