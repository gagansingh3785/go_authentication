package constants

const (
	// general constants

	// app constants
	PORT                = "3002"
	AUTHENTICATION_HOST = ""

	//database table constants
	USER_TABLE      = "app_user"
	SESSIONS_TABLE  = "app_sessions"
	ARTICLE_TABLE   = "app_articles"
	TAGS_TABLE      = "app_tags"
	ARTICLES_X_TAGS = "app_articles_app_tags"

	//Request Headers
	SESSION_COOKIE     = "X-Session-Id"
	CONTENT_TYPE_KEY   = "Content-Type"
	CONTENT_TYPE_VALUE = "application/json"

	//miscellaneous constants
	EMPTY_STRING = ""

	// keys
	MY_SECRET = "MY_SECRET"

	PAGE_SIZE = 10
)
