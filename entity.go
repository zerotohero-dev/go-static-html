package main

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserStore struct {
	Users []Credential `json:"users"`
}

type SessionKey string

const user SessionKey = "user"
