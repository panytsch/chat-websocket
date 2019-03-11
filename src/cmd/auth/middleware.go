package auth

import (
	"chat-websocket/src/cmd/db/models/user"
	"errors"
	"log"
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
		log.Println("here")
		buf := name + pass
		token := Encrypt(buf, DefaultSecret)
		log.Println(token)
		us = &user.User{
			Name:     name,
			Password: pass,
			Token:    string(token),
		}
		log.Println("token length: ", len(us.Token))
		us.SaveNew()
		log.Println(us, "after save")
		if err != nil {
			log.Println(err.Error())
			return nil, errors.New("failed")
		}
	}
	return us, nil
}
