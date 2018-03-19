package user

import (
	"db-service/common/model"
	"log"
	"testing"
)

func TestUser(t *testing.T) {
	// db.Init()
	var user *model.User = new(model.User)
	user.Name = "段凌霄"
	user.Email = "jamesduanling@hotmail.com"
	user.PortraitURL = "github.com/jamesduan/123"
	user.Sex = "男"
	user.CreateTime = "2018-3-19 13:56"
	user.UpdateTime = "2018-3-19 13:56"
	// Create(user)
	// user := model.User{Name: "jamesduan", Email: "jamesduanling@hotmail.com", Sex: "男", CreateTime: "2018-3-19 13:56", UpdateTime: "2018-3-19 13:56"}
	_, err := Create(user)
	if err != nil {
		log.Println(err)
	}
}
