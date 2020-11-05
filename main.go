package main

import (
	"io"
	"log"
	"os"
	"time"
)

func main() {
	logFile, err := os.OpenFile(time.Now().Format("./logs/2006_01_02_15_04_05")+".log", os.O_WRONLY|os.O_CREATE, 0775)
	if err != nil {
		log.Fatalln(err)
	}

	mw := io.MultiWriter(os.Stderr, logFile)
	log.SetOutput(mw)
}
