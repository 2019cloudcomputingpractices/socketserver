package model

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var dbstr = "root:wszdwlhw51868@tcp(182.254.206.244:3306)/test?charset=utf8"

func CreateUser(username string, password string) error {
	db, err := sql.Open("mysql", dbstr)
	if err != nil {
		return errors.New("database error")
	}
	stmt, err := db.Prepare("INSERT INTO user SET username=?,password=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(username, password)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("username has been registed")
	}
	return nil
}

func Login(username string, password string) error {
	db, err := sql.Open("mysql", dbstr)
	rows, err := db.Query("SELECT * FROM user where username=? and password=?", username, password)
	if err != nil {
		return errors.New("database error")
	}
	if rows.Next() {
		return nil
	} else {
		return errors.New("username and password mismatch")
	}
}
