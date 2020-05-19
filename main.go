package main

import (
    "MANGAGO/service"
    "MANGAGO/logService"
    _"MANGAGO/dbService"
)

var appName = "accountservice"

func main() {
    logService.OUT("Starting %v\n", appName)
    service.StartWebServer("6767")
}