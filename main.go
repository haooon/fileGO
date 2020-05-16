package main

import (
    "fmt"
    "MANGAGO/service"
)

var appName = "accountservice"

func main() {
    fmt.Printf("Starting %v\n", appName)
    service.StartWebServer("6767")
}