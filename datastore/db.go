package datastore

import (
    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
)

var Db, err = gorm.Open("mysql", "root:@/bukilist?charset=utf8&parseTime=True")

func init() {
    if err != nil {
        panic(err.Error())
    }

    Db.DB()
    Db.AutoMigrate(Bukiwish{})
}
