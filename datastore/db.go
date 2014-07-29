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

type DatabaseService struct {
    Model interface {}
}


func (v DatabaseService) Get (id int) (interface {}, error) {
    ans := v.Model
    if err := Db.Find(ans, id).Error; err != nil {
        return nil, ErrBukiwishNotFound
    }
    return ans, nil
}
