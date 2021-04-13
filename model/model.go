package model

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var Db *sqlx.DB

func init() {
	db, err := sqlx.Open("mysql", "root:980531@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln(err.Error())
	}
	Db = db
}

type Info struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
	Note string `json:"note"`
	Unix int64  `json:"unix"`
}

func InfoAdd(mod *Info) error {
	result, err := Db.Exec("insert into info (`name`,path,note,unix) values (?,?,?,?)", mod.Name, mod.Path, mod.Note, mod.Unix)
	if err != nil {
		return err
	}
	id, _ := result.LastInsertId()
	if id < 1 {
		return errors.New("添加失败")
	}
	return nil
}

func InfoGet(id int64) (Info, error) {
	mod := Info{}
	err := Db.Get(&mod, "select * from info where id = ?", id)
	return mod, err
}

func InfoList() ([]Info, error) {
	mods := make([]Info, 0, 8)
	err := Db.Select(&mods, "select * from info")
	return mods, err
}

func InfoDel(id int64) error {
	res, err := Db.Exec("delete from info where id = ?", id)
	if err != nil {
		return err
	}
	rows, _ := res.RowsAffected()
	if rows != 1 {
		return errors.New("删除失败")
	}
	return nil
}
