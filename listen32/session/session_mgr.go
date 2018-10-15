package session

import (
	"sync"

	uuid "github.com/satori/go.uuid"
)

type SessionMgr struct {
	sessionMap map[string]Session
	rwlock     sync.RWMutex
}

func (s *SessionMgr) Get(sessionId string) (session Session, err error) {
	s.rwlock.RLock()
	defer s.rwlock.RUnlock()

	session, ok := s.sessionMap[sessionId]
	if !ok {
		err = ErrSessionNotExist
		return
	}

	return
}

func (s *SessionMgr) CreateSession() (session Session, err error) {
	s.rwlock.Lock()
	defer s.rwlock.Unlock()

	id, err := uuid.NewV4()
	if err != nil {
		return
	}

	sessionId := id.String()
	session = NewMemorySession(sessionId)

	s.sessionMap[sessionId] = session
	return
}
