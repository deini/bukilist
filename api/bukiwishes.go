package api

import (
    "fmt"
    "github.com/deini/bukilist/datastore"
    "net/http"
)

func getBukiwishes(w http.ResponseWriter, r *http.Request) error {
    fmt.Println("ENTRE")
    bukiwishes, err := datastore.GetBukiwishes()
    if err != nil {
        return err
    }
    if bukiwishes == nil {
        bukiwishes = []*datastore.Bukiwish{}
    }

    return writeJSON(w, bukiwishes)
}
