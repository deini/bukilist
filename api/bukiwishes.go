package api

import (
    "github.com/deini/bukilist/datastore"
    "github.com/gorilla/mux"
    "io/ioutil"
    "net/http"
    "strconv"
)

var dataInterface datastore.BukiwishService

func getBukiwish(w http.ResponseWriter, r *http.Request) error {
    id, err := strconv.Atoi(mux.Vars(r)["id"])

    if err != nil {
        return err
    }
    bukiwish, err := dataInterface.Get(id)

    if err != nil {
        return err
    }
    return writeJSON(w, bukiwish)
}

func getBukiwishes(w http.ResponseWriter, r *http.Request) error {
    bukiwishes, err := dataInterface.GetAll()

    if err != nil {
        return err
    }
    return writeJSON(w, bukiwishes)
}

func createBukiwish(w http.ResponseWriter, r *http.Request) error {
    body, err := ioutil.ReadAll(r.Body)

    if err != nil {
        return err
    }
    bukiwish, err := dataInterface.Create(body)

    if err != nil {
        return err
    }
    return writeJSON(w, bukiwish)
}

func deleteBukiwish(w http.ResponseWriter, r *http.Request) error {
    id, err := strconv.Atoi(mux.Vars(r)["id"])

    if err != nil {
        return err
    }
    if err := dataInterface.Delete(id); err != nil {
        return err
    }
    w.WriteHeader(http.StatusAccepted)
    return nil
}

func updateBukiwish(w http.ResponseWriter, r *http.Request) error {
    body, err := ioutil.ReadAll(r.Body)

    if err != nil {
        return err
    }
    id, err := strconv.Atoi(mux.Vars(r)["id"])

    if err != nil {
        return err
    }
    bukiwish, err := dataInterface.Update(id, body)

    if err != nil {
        return err
    }
    return writeJSON(w, bukiwish)
}
