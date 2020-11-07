package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type ELECTION_STATUS int

const (
	READY = iota
	DURING
	DONE
)

type Election struct {
	Name   string
	Status ELECTION_STATUS

	Candidates []string

	ID int
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

	return election
}

func loadAllElections() {
	elections = nil

	root := "./elections/"
	files, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatalln("[election] cannot load elections ", err)
	}

	id := 0
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		election := loadElection(root + file.Name())
		election.ID = id

		elections = append(elections, election)

		id++
	}

	log.Println("[election] load all elections done")
}
