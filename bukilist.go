package main

import (
    "encoding/json"
    _ "github.com/go-sql-driver/mysql"
    "github.com/gorilla/mux"
    "github.com/jinzhu/gorm"
    "io/ioutil"
    "log"
    "net/http"
    "time"
)

var db, err = gorm.Open("mysql", "root:@/bukilist?charset=utf8&parseTime=True")

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

func newBukiwish(w http.ResponseWriter, r *http.Request) { //, data interface{}) {
    var bukiwish Bukiwish
    buf, _ := ioutil.ReadAll(r.Body)
    json.Unmarshal(buf, &bukiwish)
    db.Create(&bukiwish)
}

func getBukiwishes(w http.ResponseWriter, r *http.Request) {
    bukiwishes := []Bukiwish{}
    db.Find(&bukiwishes)

    response, err := json.Marshal(bukiwishes)
    if err != nil {
        panic(err.Error())
    }
    w.WriteHeader(200)
    w.Write(response)
}

func getBukiwish(w http.ResponseWriter, r *http.Request) {
    var bukiwish Bukiwish
    params := mux.Vars(r)
    id := params["id"]
    db.Find(&bukiwish, id)

    response, err := json.Marshal(bukiwish)
    if err != nil {
        panic(err.Error())
    }
    w.WriteHeader(200)
    w.Write(response)
}

func updateBukiwish(w http.ResponseWriter, r *http.Request) {
    var bukiwish, currentBukiwish Bukiwish
    params := mux.Vars(r)
    id := params["id"]

    buf, _ := ioutil.ReadAll(r.Body)
    json.Unmarshal(buf, &bukiwish)

    db.Find(&currentBukiwish, id)
    db.Model(&currentBukiwish).Updates(bukiwish)
}

func deleteBukiwish(w http.ResponseWriter, r *http.Request) {
    var bukiwish Bukiwish
    params := mux.Vars(r)
    id := params["id"]
    db.Find(&bukiwish, id)
    db.Delete(&bukiwish)
    w.WriteHeader(202)
}

func main() {

    if err != nil {
        panic(err.Error())
    }

    db.DB()
    db.CreateTable(Bukiwish{})
    db.AutoMigrate(Bukiwish{})

    rtr := mux.NewRouter()
    srtr := rtr.PathPrefix("/api").Subrouter()
    srtr.HandleFunc("/bukiwishes/{id:[0-9]+}", getBukiwish).Methods("GET")
    srtr.HandleFunc("/bukiwishes/{id:[0-9]+}", updateBukiwish).Methods("PUT")
    srtr.HandleFunc("/bukiwishes/{id:[0-9]+}", deleteBukiwish).Methods("DELETE")
    srtr.HandleFunc("/bukiwishes", getBukiwishes).Methods("GET")
    srtr.HandleFunc("/bukiwishes", newBukiwish).Methods("POST")
    rtr.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
    http.Handle("/", rtr)

    log.Println("Listening...")
    http.ListenAndServe(":3000", nil)
}
