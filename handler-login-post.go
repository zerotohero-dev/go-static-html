package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func handleLoginPost(router *gin.Engine, store *UserStore) {
	router.POST(urlLogin, func(c *gin.Context) {
		session := sessions.Default(c)
		u := session.Get(user)
		if u != nil {
			c.String(http.StatusBadRequest, "Please log out first.")
			return
		}

		username := c.PostForm("username")
		password := c.PostForm("password")
		if username == "" || password == "" {
			c.String(http.StatusBadRequest, "Empty username or password.")
			return
		}

		users := store.Users
		found := false
		for _, u := range users {
			if u.Username == username && u.Password == password {
				found = true
				session.Set(string(user), username)
				break
			}
		}

		if !found {
			c.String(http.StatusBadRequest, "Incorrect username or password.")
			return
		}

		if err := session.Save(); err != nil {
			c.String(http.StatusInternalServerError, "Failed to save session")
			return
		}

		c.Redirect(http.StatusFound, urlHome)
		c.Abort()
	})
}
