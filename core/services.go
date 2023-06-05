package core

import (
	"errors"
	"log"
)

func Login(username string, password string) error {
	log.Printf("Login with username %v", username)
	if username == "idev" && password == "idev" {
		return nil
	}
	return errors.New("Username or password is wrong")
}

func Register(username string, password string) error {
	log.Printf("Register with username %v", username)
	return nil
}