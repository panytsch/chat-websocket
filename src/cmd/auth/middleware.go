package auth

import (
	"chat-websocket/src/cmd/db/models/user"
	"errors"
)

func auth(name string, pass string) (*user.User, error) {
	us := user.FindByNamePass(name, pass)
	if us == nil {
		return nil, errors.New("not found")
	}
	return us, nil
}

func Register(name string, pass string) (*user.User, error) {
	us, err := auth(name, pass)
	if err != nil {
		buf := name + pass
		us = &user.User{
			Name:     name,
			Password: pass,
			Token:    Encrypt(buf, DefaultSecret),
		}
		us.SaveNew()
	}
	return us, nil
}
