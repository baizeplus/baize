package session

import (
	"baize/app/constant/sessionStatus"
	"baize/app/utils/session/sessionCache"
	"github.com/gin-gonic/gin"
)

type Manager struct {
	Propagator
	Store
}

func NewManger() *Manager {
	return &Manager{
		Propagator: sessionCache.NewPropagator(),
		Store:      sessionCache.NewStore(),
	}
}

func (m *Manager) GetSession(ctx *gin.Context) (Session, error) {
	val, ok := ctx.Get(sessionStatus.SessionKey)
	if ok {
		return val.(*sessionCache.Session), nil
	}
	sessId, err := m.Extract(ctx)
	if err != nil {
		return nil, err
	}
	sess, err := m.Get(ctx, sessId)
	if err != nil {
		return nil, err
	}
	ctx.Set(sessionStatus.SessionKey, sess)
	return sess, nil
}
func (m *Manager) InitSession(ctx *gin.Context, userId int64) (Session, error) {
	sess, err := m.Generate(ctx, userId)
	if err != nil {
		return nil, err
	}
	err = m.Refresh(ctx, sess.Id())
	if err != nil {
		return nil, err
	}
	ctx.Set(sessionStatus.SessionKey, sess)
	return sess, err
}
func (m *Manager) InitAppSession(ctx *gin.Context, userId int64) (Session, error) {
	sess, err := m.Generate(ctx, userId)
	if err != nil {
		return nil, err
	}
	ctx.Set(sessionStatus.SessionKey, sess)
	return sess, err
}
func (m *Manager) RemoveSession(ctx *gin.Context) {
	sess, err := m.GetSession(ctx)
	if err != nil {
		return
	}
	_ = m.Store.Remove(ctx, sess.Id())

}
func (m *Manager) RefreshSession(ctx *gin.Context) error {
	sess, err := m.GetSession(ctx)
	if err != nil {
		return err
	}
	return m.Refresh(ctx, sess.Id())
}
