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

type BukiwishService struct{}

func (v BukiwishService) Get(id int) (*Bukiwish, error) {
    bukiwish := new(Bukiwish)
    service := DatabaseService{bukiwish}
    ans, _ := service.Get(id)
    ret := ans.(Bukiwish)
    return &ret, nil
}

func (v BukiwishService) GetAll() (*[]*Bukiwish, error) {
    bukiwishes := new([]*Bukiwish)

    if err := Db.Find(bukiwishes).Error; err != nil {
        return nil, err
    }
    return bukiwishes, nil
}

func (v BukiwishService) Create(body []byte) (*Bukiwish, error) {
    bukiwish := new(Bukiwish)

    if err := json.Unmarshal(body, bukiwish); err != nil {
        return nil, err
    }
    if err := Db.Create(bukiwish).Error; err != nil {
        return nil, err
    }
    return bukiwish, nil
}

func (v BukiwishService) Delete(id int) error {
    bukiwish := new(Bukiwish)

    if err := Db.Find(bukiwish, id).Error; err != nil {
        return err
    }
    if err := Db.Delete(bukiwish).Error; err != nil {
        return err
    }
    return nil
}

func (v BukiwishService) Update(id int, body []byte) (*Bukiwish, error) {
    bukiwish := new(Bukiwish)
    currentBukiwish := new(Bukiwish)

    if err := json.Unmarshal(body, bukiwish); err != nil {
        return nil, err
    }
    if err := Db.Find(currentBukiwish, id).Error; err != nil {
        return nil, err
    }
    if err := Db.Model(currentBukiwish).Updates(bukiwish).Error; err != nil {
        return nil, err
    }
    return currentBukiwish, nil
}
