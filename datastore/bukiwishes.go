package datastore

import (
    //     // "fmt"
    "github.com/deini/bukilist"
    //     "strings"
    //     "time"
)

func GetBukiwishes() ([]*bukilist.Bukiwish, error) {
    var bukiwishes []*bukilist.Bukiwish
    query := Db.Find(&bukiwishes)

    if query.Error != nil {
        return nil, query.Error
    }

    return bukiwishes, nil
}
