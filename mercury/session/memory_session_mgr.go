package session

import (
	"sync"

	uuid "github.com/satori/go.uuid"
)

type MemorySessionMgr struct {
	sessionMap map[string]Session
	rwlock     sync.RWMutex
}

func NewMemorySessionMgr() SessionMgr {
	sr := &MemorySessionMgr{
		sessionMap: make(map[string]Session, 1024),
	}

	return sr
}

func (s *MemorySessionMgr) Init(addr string, options ...string) (err error) {
	return
}

func (s *MemorySessionMgr) Get(sessionId string) (session Session, err error) {
	s.rwlock.RLock()
	defer s.rwlock.RUnlock()

	session, ok := s.sessionMap[sessionId]
	if !ok {
		err = ErrSessionNotExist
		return
	}

	return
}

func (s *MemorySessionMgr) CreateSession() (session Session, err error) {
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
