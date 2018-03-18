package rpc

import (
	"db-service/common/model"
	"db-service/db/user"
	"log"
)

func (rpcuser *User) CreateUser(req model.User, res *model.User) error {
	_, err := user.Create(&req)
	if err != nil {
		log.Println(err)
	}
	return nil
}
