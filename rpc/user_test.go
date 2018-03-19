package rpc

import (
	"db-service/common/model"
	"fmt"
	"os"
	"testing"
	"time"
)

const RPCServer = "127.0.0.1:6030"

// const Port = 6030

func TestGetUser(t *testing.T) {
	timeout := time.Duration(1) * time.Second
	client, err := JsonRpcClient("tcp", RPCServer, timeout)
	if err != nil {
		fmt.Println("dial failed:", err)
		os.Exit(1)
	}
	response := new(model.User)
	client.Call("User.GetUser", 3, &response)
}
func TestUser(t *testing.T) {

	timeout := time.Duration(1) * time.Second
	client, err := JsonRpcClient("tcp", RPCServer, timeout)
	if err != nil {
		fmt.Println("dial failed:", err)
		os.Exit(1)
	}
	// db.Init()
	var user *model.User = new(model.User)
	user.Name = "张乐"
	user.Email = "jamesduanling@hotmail.com"
	user.PortraitURL = "github.com/jamesduan/123"
	user.Sex = "男"
	user.CreateTime = "2018-3-19 13:56"
	user.UpdateTime = "2018-3-19 13:56"
	response := new(model.User)
	client.Call("User.CreateUser", user, &response)
	// Create(user)
	// user := model.User{Name: "jamesduan", Email: "jamesduanling@hotmail.com", Sex: "男", CreateTime: "2018-3-19 13:56", UpdateTime: "2018-3-19 13:56"}
	// _, err := Create(user)
	// if err != nil {
	// 	log.Println(err)
	// }
}
