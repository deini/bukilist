package api

import (
    "fmt"
    "github.com/deini/bukilist/router"
    "github.com/gorilla/mux"
    "log"
    "net/http"
)

func Handler() *mux.Router {
    m := router.API()
    m.Get(router.Bukiwishes).Handler(handler(getBukiwishes))
    // m.Get(router.Bukiwish).Handler(handler(getBukiwish))
    // m.Get(router.CreateBukiwish).Handler(handler(createBukiwish))
    // m.Get(router.UpdateBukiwish).Handler(handler(updateBukiwish))
    // m.Get(router.DeleteBukiwish).Handler(handler(deleteBukiwish))
    return m
}

type handler func(http.ResponseWriter, *http.Request) error

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    err := h(w, r)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "error: %s", err)
        log.Println(err)
    }
}
