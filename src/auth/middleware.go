package auth

import (
	"chat-websocket/src/db/models/user"
	"errors"
	"log"
)

func auth(name string, pass string) (*user.User, error) {
	us, err := user.FindByNamePass(name, pass)
	if err != nil {
		return nil, err
	}
	return us, nil
}

func Register(name string, pass string) (*user.User, error) {
	us, err := auth(name, pass)
	if err != nil {
		log.Println("here")
		buf := []byte(name + pass)
		us = &user.User{
			Name:     name,
			Password: pass,
			Token:    string(Encrypt(buf, DEFAULT_SECRET)),
		}
		_, err := us.SaveNew()
		log.Println(us)
		if err != nil {
			return nil, errors.New("failed")
		}
	}
	return us, nil
}
