package service

import (
    "MANGAGO/logService"
    "net/http"
)
func StartWebServer(port string) {
    r := NewRouter()
    http.Handle("/", r)
    
    logService.OUT("Starting HTTP service at " + port)
    err := http.ListenAndServe(":"+port, nil) // Goroutine will block here

    if err != nil {
        logService.OUT("An error occured starting HTTP listener at port " + port)
        logService.OUT("Error: " + err.Error())
    }
}

