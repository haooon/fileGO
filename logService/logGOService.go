package logService

import (
	"fmt"
	"log"
)

func init(){
	fmt.Println("================ init logService ===================")
}

func init() {
    log.SetPrefix("TRACE: ")
    //log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
    log.SetFlags(log.Ldate | log.Lmicroseconds)
}

func OUT(param ... string) {
	var msg = ""
	for _, msg_i := range param{
		msg += msg_i
	}
	log.Println(msg)
}

func log_to_file() {
	return
}