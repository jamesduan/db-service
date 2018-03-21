package rpc

import (
	"db-service/common/model"
	"db-service/db/user"
	"log"
)

func (rpcuser *User) CreateUser(req *model.User, res *model.User) error {
	err := user.Create(req)
	if err != nil {
		log.Println(err)
	}
	res = nil
	return nil
}

func (rpcuser *User) GetUser(userId int, res *model.User) error {
	err := user.SelectUser(userId)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func (rpcuser *User) UpdateUser(req *model.User, res *model.User) error {
	err := user.Update(req)
	if err != nil {
		log.Println(err)
	}
	return nil
}
