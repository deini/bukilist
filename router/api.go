package router

import "github.com/gorilla/mux"

func API() *mux.Router {
    m := mux.NewRouter()
    m.Path("/bukiwishes").Methods("GET").Name(Bukiwishes)
    // m.Path("/bukiwishes").Methods("POST").Name(CreateBukiwish)
    // m.Path("/bukiwishes/{Id:.+}").Methods("GET").Name(Bukiwish)
    // m.Path("/bukiwishes/{Id:.+}").Methods("PUT").Name(UpdateBukiwish)
    // m.Path("/bukiwishes/{Id:.+}").Methods("DELETE").Name(DeleteBukiwish)
    return m
}
