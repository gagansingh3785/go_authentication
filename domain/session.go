package domain

type Session struct {
	UserID    string `DB:"user_id"`
	SessionID string `DB:"session_id"`
}

func (s *Session) GetSessionID() string {
	return s.SessionID
}

func (s *Session) GetSessionUserID() string {
	return s.UserID
}
