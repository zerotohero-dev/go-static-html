package main

import (
	"github.com/gin-gonic/gin"
)

func serve(router *gin.Engine) {
	if err := router.Run(host); err != nil {
		panic("Poop!")
	}
}
