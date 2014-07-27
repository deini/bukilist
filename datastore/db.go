package datastore

import (
    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
)

var Db

func init() {
    Db, err = gorm.Open("mysql", "root:@/bukilist?charset=utf8&parseTime=True")

    if err != nil {
        panic(err.Error())
    }

    Db.DB()
    Db.AutoMigrate(Bukiwish{})
}
