package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func handleCatchAll(router *gin.Engine) {
	router.GET("/*context", func(c *gin.Context) {
		urlPath := path.Join(wwwRoot, c.Request.URL.Path)

		if c.Request.URL.Path != "/login.go" {
			c.File(urlPath)
			return
		}

		session := sessions.Default(c)

		user := session.Get("user")
		if user != nil {
			c.Redirect(http.StatusFound, "/index.html")
			c.Abort()
			return
		}

		c.File(path.Join(wwwRoot, "/auth/login.html"))
	})
}
