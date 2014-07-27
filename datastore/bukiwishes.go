package datastore

import (
    "encoding/json"
    "errors"
    "time"
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

type BukiwishesService interface {
    // fail??
    CreateBukiwish() (Bukiwish, error)
    GetBukiwishes() ([]*Bukiwish, error)
}

var (
    ErrBukiwishNotFound = errors.New("Bukiwish not found")
)

func GetBukiwish(id int) (Bukiwish, error) {
    var bukiwish Bukiwish

    if err := Db.Find(&bukiwish, id).Error; err != nil {
        return bukiwish, ErrBukiwishNotFound
    }

    return bukiwish, nil
}

func GetBukiwishes() ([]*Bukiwish, error) {
    var bukiwishes []*Bukiwish
    query := Db.Find(&bukiwishes)

    if query.Error != nil {
        return nil, query.Error
    }

    return bukiwishes, nil
}

func CreateBukiwish(body []byte) (Bukiwish, error) {
    var bukiwish Bukiwish

    if err := json.Unmarshal(body, &bukiwish); err != nil {
        return bukiwish, err
    }

    if err := Db.Create(&bukiwish).Error; err != nil {
        return bukiwish, err
    }

    return bukiwish, nil
}

func DeleteBukiwish(id int) error {
    var bukiwish Bukiwish

    if err := Db.Find(&bukiwish, id).Error; err != nil {
        return err
    }

    if err := Db.Delete(&bukiwish).Error; err != nil {
        return err
    }

    return nil
}

func UpdateBukiwish(id int, body []byte) (Bukiwish, error) {
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
