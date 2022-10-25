package main

import "os"

var secret = []byte(os.Getenv("GO_STATIC_SESSION_SECRET")) // long_secret_string
var wwwRoot = os.Getenv("GO_STATIC_WWW_ROOT")              // $WORKSPACE/go-static-html/public
var host = os.Getenv("GO_STATIC_HOST")                     // localhost:8080
var storePath = os.Getenv("GO_STATIC_STORE_PATH")          // $WORKSPACE/go-static-html/users.json
var urlHome = os.Getenv("GO_STATIC_URL_HOME")              // /index.html
var urlLogin = os.Getenv("GO_STATIC_URL_LOGIN")            // /login.go
var urlLoginHtml = os.Getenv("GO_STATIC_LOGIN_HTML")       // $WORKSPACE/go-static-html/public/auth/login/index.html

func checkEnv() {
	if len(secret) == 0 {
		panic("Empty secret!")
	}

	if len(wwwRoot) == 0 {
		panic("Empty wwwRoot!")
	}

	if len(host) == 0 {
		panic("Empty host!")
	}

	if len(storePath) == 0 {
		panic("Empty store path!")
	}

	if len(urlHome) == 0 {
		panic("Empty home url")
	}

	if len(urlLogin) == 0 {
		panic("Empty login url")
	}

	if len(urlLoginHtml) == 0 {
		panic("Empty login html file path")
	}
}
