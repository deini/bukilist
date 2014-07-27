package datastore

import (
    "encoding/json"
    "errors"
    "time"
)

var (
    ErrBukiwishNotFound = errors.New("Bukiwish not found")
)

type Bukiwish struct {
    Id          int64
    Name        string `sql:"not null;unique;type:varchar(100)"`
    Description string
    Rate        int
    Worth       bool
    Location    string
    Geolocation string
    CreatedAt   time.Time
}

type BukiwishService struct {}

func (v BukiwishService) Get(id int) (Bukiwish, error) {
    var bukiwish Bukiwish
    if err := Db.Find(&bukiwish, id).Error; err != nil {
        return bukiwish, ErrBukiwishNotFound
    }
    return bukiwish, nil
}

func (v BukiwishService) GetAll() ([]*Bukiwish, error) {
    var bukiwishes []*Bukiwish
    if err := Db.Find(&bukiwishes).Error; err != nil {
        return nil, err
    }
    return bukiwishes, nil
}

func (v BukiwishService) Create(body []byte) (Bukiwish, error) {
    var bukiwish Bukiwish
    if err := json.Unmarshal(body, &bukiwish); err != nil {
        return bukiwish, err
    }
    if err := Db.Create(&bukiwish).Error; err != nil {
        return bukiwish, err
    }
    return bukiwish, nil
}

func (v BukiwishService) Delete(id int) error {
    var bukiwish Bukiwish
    if err := Db.Find(&bukiwish, id).Error; err != nil {
        return err
    }
    if err := Db.Delete(&bukiwish).Error; err != nil {
        return err
    }
    return nil
}

func (v BukiwishService) Update(id int, body []byte) (Bukiwish, error) {
    var bukiwish, currentBukiwish Bukiwish
    if err := json.Unmarshal(body, &bukiwish); err != nil {
        return bukiwish, err
    }
    if err := Db.Find(&currentBukiwish, id).Error; err != nil {
        return bukiwish, err
    }
    if err := Db.Model(&currentBukiwish).Updates(bukiwish).Error; err != nil {
        return bukiwish, err
    }
    return currentBukiwish, nil
}
