package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func InitLogs() {
	logFile, err := os.OpenFile("api_gateway.log", os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("fail to open api_gateway log file", err)
		return
	}

	multiWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiWriter)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func ErrorLog(err string) {

}
