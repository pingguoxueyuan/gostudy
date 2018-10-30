package session

import "fmt"

var (
	sessionMgr SessionMgr
)

//provider:
//1. memory， 返回一个内存的session管理类
//2. redis, 返回一个redis的session管理类
func Init(provider string, addr string, options ...string) (err error) {

	switch provider {
	case "memory":
		sessionMgr = NewMemorySessionMgr()
	case "redis":
		sessionMgr = NewRedisSessionMgr()
	default:
		err = fmt.Errorf("not support")
		return
	}

	err = sessionMgr.Init(addr, options...)
	return
}

func CreateSession() (session Session, err error) {
	return sessionMgr.CreateSession()
}

func Get(sessionId string) (session Session, err error) {
	return sessionMgr.Get(sessionId)
}
