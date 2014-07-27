package datastore

import (
    //     // "fmt"
    // "github.com/deini/bukilist"
    //     "strings"
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
    // Get a bukiwish
    // Get(id int) (*Bukiwish, error)

    // Get all bukiwishes
    // GetAll() ([]*Bukiwish, error)
    GetBukiwishes() ([]*Bukiwish, error)
    // NewBukiwish(bukiwish *Bukiwish) (created bool, err error)
}

// var (
//     ErrBukiwishNotFound = errors.New("Bukiwish not found")
// )

// func GetAll() ([]*Bukiwish, error) {
//     bukiwishes, err := datastore.GetBukiwishes()
//     if err != nil {
//         return nil, err
//     }

//     return bukiwishes, nil
// }

func GetBukiwishes() ([]*Bukiwish, error) {
    var bukiwishes []*Bukiwish
    query := Db.Find(&bukiwishes)

    if query.Error != nil {
        return nil, query.Error
    }

    return bukiwishes, nil
}
