package service

import (
    //go get -v -u github.com/gorilla/mux
    "MANGAGO/logService"
    "github.com/gorilla/mux"
)

// Function that returns a pointer to a mux.Router we can use as a handler.
func NewRouter() *mux.Router {
    // Create an instance of the Gorilla router
    router := mux.NewRouter().StrictSlash(true)

    // Iterator over the routes we declared in routes.go and attach them to the router instance
    for _, route := range routes {
        // Attach each route, uses a Builder-like pattern to set each route up.
        logService.OUT("loading ROUTE:", route.Method, route.Pattern, route.Name)
        router.Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(route.HandlerFunc)
    }

    return router
}