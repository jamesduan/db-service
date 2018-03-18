package user

import (
	"db-service/common/model"
	"db-service/db"
	"fmt"
	"log"
)

func Create(user *model.User) (bool, error) {
	sql := fmt.Sprintf(
		"insert into user(name, email, sex, portrait_url, create_time, update_time) values ('%s', '%s', '%s', '%s', '%s', '%s')",
		user.Name,
		user.Email,
		user.Sex,
		user.PortraitURL,
		user.CreateTime,
		user.UpdateTime,
	)
	_, err := db.DB.Exec(sql)
	if err != nil {
		log.Println("exec", sql, "fail", err)
		return false, err
	} else {
		return true, nil
	}
}

func Update(user *model.User) {

}

func Delete(userId int) {

}

func selectUser(userId int) {

}

func selectUserByName(username string) {

}
