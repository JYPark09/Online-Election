package main

import (
	"bufio"
	"context"
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

	srv := startServer(":52525")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "stop" {
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()

			if err = srv.Shutdown(ctx); err != nil {
				log.Fatalln("[http] shutdown failed ", err)
			}

			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
}
