package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type ELECTION_STATUS int

const (
	NOT_START = iota
	DURING
	DONE
)

type Election struct {
	Name   string
	Status ELECTION_STATUS

	Candidates []string
}

var elections []Election

func loadElection(filepath string) Election {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalln("[election] cannot load election ", err)
	}

	var election Election
	err = json.Unmarshal(file, &election)
	if err != nil {
		log.Fatalln("[election] cannot unmarshal election ", err)
	}

	fmt.Println(election)

	return election
}

func loadAllElections() {
	elections = nil

	root := "./elections/"
	files, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatalln("[election] cannot load elections ", err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		elections = append(elections, loadElection(root+file.Name()))
	}

	log.Println("[election] load all elections done")
}
