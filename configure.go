package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func configure(router *gin.Engine) *gin.Engine {
	router.LoadHTMLGlob("**/*.html")
	router.LoadHTMLGlob("**/**/*.html")
	router.Use(sessions.Sessions("session", cookie.NewStore(secret)))
	router.Use(authMiddleware)
	return router
}
