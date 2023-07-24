package models

import (
    "github.com/go-xorm/xorm"
_ "github.com/go-sql-driver/mysql"
)

var engine *xorm.Engine

func init () {
	var err error
	engine, err = xorm.NewEngine("mysql", "fuga:pass@/bookshelf")
    if err != nil {
        panic(err)
    }
}

type User struct {
	ID       int    `json:"id" xorm:"'id'"`
	Username string `json:"name" xorm:"'nickname'"`
}

func NewUser (id int, username string) User {
	return User{
		ID: id,
		Username: username,
	}
}

type UserRepository struct {
}

func NewUserRepository () UserRepository {
	return UserRepository{}
}

func (u UserRepository) GetById (id int) *User {
	var user = User{ID: id}
	has, _ := engine.Get(&user)
	if has {
		return &user
	}
	return nil
}
