package main

import (
    "flag"
    "github.com/deini/bukilist/api"
    "log"
    "net/http"
)

var port = flag.String("port", "3000", "Port to use")

func main() {
    //     srtr := rtr.PathPrefix("/api").Subrouter()
    //     rtr.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
    //     http.Handle("/", rtr)

    //     flag.Parse()
    //     log.Println("What happens on port " + *port + " stays on port " + *port)
    //     panic(hthesrcttp.ListenAndServe(":"+*port, nil))
    m := http.NewServeMux()
    m.Handle("/api/", http.StripPrefix("/api", api.Handler()))
    // m.Handle("/api/", api.Handler())
    // m.Handle("/", app.Handler())
    flag.Parse()
    log.Println("What happens on port " + *port + " stays on port " + *port)
    panic(http.ListenAndServe(":"+*port, nil))
}
