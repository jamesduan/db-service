package user

import (
	"db-service/common/model"
	"db-service/db"
	"fmt"
	"log"
)

func Create(user *model.User) error {
	sql := fmt.Sprintf(
		"insert into user(name, email, sex, portrait_url, create_time, update_time) values ('%s', '%s', '%s', '%s', '%s', '%s')",
		user.Name,
		user.Email,
		user.Sex,
		user.PortraitURL,
		user.CreateTime,
		user.UpdateTime,
	)
	// log.Println(sql)
	_, err := db.DB.Exec(sql)
	if err != nil {
		log.Println("exec", sql, "fail", err)
		return err
	}
	return nil
}

func Update(user *model.User) error {
	return nil
}

func Delete(userId int) {

}

func SelectUser(userId int) error {
	sql := fmt.Sprintf("select * from user where id=%d", userId)
	// log.Println(sql)
	rows, err := db.DB.Query(sql)
	log.Println(rows)
	// for row := range rows.Next() {
	// 	log.Println(row)
	// }
	for rows.Next() {
		var (
			id          int
			name        string
			email       string
			sex         string
			portraitURL string
			createTime  string
			updateTime  string
		)
		err := rows.Scan(&id, &name, &email, &sex, &portraitURL, &createTime, &updateTime)
		if err != nil {
			log.Println(err)
		}
		log.Println(id, name, email, sex, portraitURL, createTime, updateTime)
	}
	if err != nil {
		log.Println("exec", sql, "fail", err)
		return err
	}
	defer rows.Close()
	return nil
}

func selectUserByName(username string) {

}
