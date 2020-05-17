package service

import (
    "net/http"
    "MANGAGO/fileServer"
)

// Define a single route, e.g. a human readable name, HTTP method and the pattern the function that will execute when the route is called.

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

// Defines the type Routes which is just an array (slice) of Route structs.
type Routes []Route

var routes = Routes{
    Route{
        "GetAccount", // Name
        "GET",        // HTTP method
        "/accounts/{accountId}", // Route pattern
        func(w http.ResponseWriter, r *http.Request) {
            w.Header().Set("Content-Type", "application/json; charset=UTF-8")
            w.Write([]byte("{\"result\":\"OK\"}"))
        },
    },
    Route{
        "fileget", // Name
        "GET",        // HTTP method
        "/upload", // Route pattern
        func(w http.ResponseWriter, r *http.Request) {
            fileServer.GETuploadHandler(w, r)
        },
    },
    Route{
        "filepost", // Name
        "POST",        // HTTP method
        "/uploadp", // Route pattern
        func(w http.ResponseWriter, r *http.Request) {
            fileServer.POSTuploadHandler(w, r)
        },
    },
    Route{
        "filepostoption", // Name
        "OPTIONS",        // HTTP method
        "/uploadp", // Route pattern
        func(w http.ResponseWriter, r *http.Request) {
            fileServer.OPTIONSPOSTuploadHandler(w, r)
        },
    },
}