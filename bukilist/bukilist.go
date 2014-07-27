package main

import (
    "flag"
    "github.com/deini/bukilist"
    "github.com/deini/bukilist/api"
    "github.com/deini/bukilist/datastore"
    "github.com/deini/bukilist/router"
    // "github.com/gorilla/mux"
    "net/http"
)

// import (
//     "encoding/json"
//     "flag"
//     _ "github.com/go-sql-driver/mysql"
//     "github.com/gorilla/mux"
//     "github.com/jinzhu/gorm"
//     "io/ioutil"
//     "log"
//     "net/http"
//     "time"
// )

var port = flag.String("port", "3000", "Port to use")

// var db, err = gorm.Open("mysql", "root:@/bukilist?charset=utf8&parseTime=True")

// type Bukiwish struct {
//     Id          int64
//     Name        string `sql:"not null;unique;type:varchar(100)"`
//     Description string
//     Rate        int
//     Worth       bool
//     Location    string
//     Geolocation string
//     CreatedAt   time.Time
// }

// func newBukiwish(w http.ResponseWriter, r *http.Request) {
//     var bukiwish Bukiwish
//     body, _ := ioutil.ReadAll(r.Body)
//     json.Unmarshal(body, &bukiwish)
//     db.Create(&bukiwish)
// }

// func getBukiwishes(w http.ResponseWriter, r *http.Request) {
//     bukiwishes := []Bukiwish{}
//     db.Find(&bukiwishes)

//     response, err := json.Marshal(bukiwishes)
//     if err != nil {
//         panic(err.Error())
//     }
//     w.WriteHeader(200)
//     w.Write(response)
// }

// func getBukiwish(w http.ResponseWriter, r *http.Request) {
//     var bukiwish Bukiwish
//     params := mux.Vars(r)
//     id := params["id"]

//     if db.Find(&bukiwish, id).RecordNotFound() {
//         http.Error(w, "Record not found", 404)
//         return
//     }

//     response, err := json.Marshal(bukiwish)
//     if err != nil {
//         panic(err.Error())
//     }
//     w.WriteHeader(200)
//     w.Write(response)
// }

// func updateBukiwish(w http.ResponseWriter, r *http.Request) {
//     var bukiwish, currentBukiwish Bukiwish
//     params := mux.Vars(r)
//     id := params["id"]

//     body, _ := ioutil.ReadAll(r.Body)
//     json.Unmarshal(body, &bukiwish)

//     if db.Find(&currentBukiwish, id).RecordNotFound() {
//         http.Error(w, "Record not found", 404)
//         return
//     }
//     db.Model(&currentBukiwish).Updates(bukiwish)
// }

// func deleteBukiwish(w http.ResponseWriter, r *http.Request) {
//     var bukiwish Bukiwish
//     params := mux.Vars(r)
//     id := params["id"]
//     db.Find(&bukiwish, id)
//     db.Delete(&bukiwish)
//     w.WriteHeader(202)
// }

func main() {
    //     srtr := rtr.PathPrefix("/api").Subrouter()
    //     rtr.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
    //     http.Handle("/", rtr)

    //     flag.Parse()
    //     log.Println("What happens on port " + *port + " stays on port " + *port)
    //     panic(hthesrcttp.ListenAndServe(":"+*port, nil))
    m := http.NewServeMux()
    m.Handle("/api/", http.StripPrefix("/api", api.Handler()))
    // m.Handle("/", app.Handler())
    flag.Parse()
    log.Println("What happens on port " + *port + " stays on port " + *port)
    panic(hthesrcttp.ListenAndServe(":"+*port, nil))
}