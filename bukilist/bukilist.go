package main

import (
    "flag"
    "github.com/deini/bukilist/api"
    "log"
    "net/http"
)

var port = flag.String("port", "3000", "Port to use")

func main() {
    m := http.NewServeMux()
    m.Handle("/api/", http.StripPrefix("/api", api.Handler()))
    m.Handle("/", http.FileServer(http.Dir("./static/")))

    flag.Parse()
    log.Println("What happens on port " + *port + " stays on port " + *port)

    panic(http.ListenAndServe(":"+*port, m))
}
