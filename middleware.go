package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func authMiddleware(c *gin.Context) {
	p := c.Request.URL.Path

	if p == "/login.go" {
		c.Next()
		return
	}
	if !strings.HasSuffix(strings.ToLower(p), ".html") && len(p) > 0 && p[len(p)-1] != '/' {
		c.Next()
		return
	}

	session := sessions.Default(c)
	u := session.Get(string(user))
	if u == nil {
		c.Redirect(http.StatusFound, "/login.go")
		c.Abort()
		return
	}

	c.Next()
}
