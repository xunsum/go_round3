package utils

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

//todo: 挂载 session 中间件，未使用

type SessionLoader struct{}

func (loader *SessionLoader) LoadSession(r *gin.Engine) {
	store := cookie.NewStore([]byte("abdkfgyt65$3uiobjhkl^"))
	store.Options(sessions.Options{
		Domain:   "/",
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
	})
	r.Use(sessions.Sessions("mySession", store))
}
