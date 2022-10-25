package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	checkEnv()

	store := initStore()

	router := configure(gin.Default())

	handleLoginPost(router, store)
	handleCatchAll(router)

	serve(router)
}
