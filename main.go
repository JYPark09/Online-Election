package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	logFile, err := os.OpenFile(time.Now().Format("./logs/2006_01_02_15_04_05")+".log", os.O_WRONLY|os.O_CREATE, 0775)
	if err != nil {
		log.Fatalln(err)
	}

	mw := io.MultiWriter(os.Stderr, logFile)
	log.SetOutput(mw)

	initDatabaseManager("database.db")

	loadAllElections()

	port := 52525
	if len(os.Args) == 2 {
		if port, err = strconv.Atoi(os.Args[1]); err != nil {
			log.Fatalln("invalid port ", os.Args[1])
		}
	}

	srv := startServer(":" + strconv.Itoa(port))

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		cmd := scanner.Text()

		if cmd == "stop" {
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()

			if err = srv.Shutdown(ctx); err != nil {
				log.Fatalln("[http] shutdown failed ", err)
			}

			break
		} else if cmd == "reload_elections" {
			loadAllElections()
		} else if cmd == "begin" {
			scanner.Scan()
			id, err := strconv.Atoi(scanner.Text())
			if err != nil || !beginElection(id) {
				fmt.Println("Cannot begin election.")
			}
		} else if cmd == "end" {
			scanner.Scan()
			id, err := strconv.Atoi(scanner.Text())
			if err != nil || !endElection(id) {
				fmt.Println("Cannot end election.")
			}
		} else if cmd == "list_election" {
			for idx, elect := range elections {
				fmt.Printf("%d - %s\n", idx, elect.Name)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	shutdownDatabaseManager()
}
