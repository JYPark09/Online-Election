package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
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

	fname string
	users []int
	votes []string
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

	election.fname = filepath
	election.Status = READY

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
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".json") {
			continue
		}

		election := loadElection(root + file.Name())
		election.ID = id

		elections = append(elections, election)

		id++
	}

	log.Println("[election] load all elections done")
}

func getElection(id int) *Election {
	if len(elections) <= id {
		return nil
	}

	return &elections[id]
}

func beginElection(id int) bool {
	elect := getElection(id)

	if elect == nil {
		log.Println("[election] invalid election. " + strconv.Itoa(id))
		return false
	}

	if elect.Status != READY {
		return false
	}

	elect.users = nil
	elect.votes = nil

	elect.Status = DURING

	log.Println("[election] begin election " + elect.Name)

	return true
}

func endElection(id int) bool {
	elect := getElection(id)

	if elect == nil {
		log.Println("[election] invalid election. " + strconv.Itoa(id))

		return false
	}

	var result struct {
		Users []int
		Votes []string
	}

	result.Users = elect.users
	result.Votes = elect.votes

	file, _ := json.MarshalIndent(result, "", "  ")
	err := ioutil.WriteFile(elect.fname+"_result", file, 0775)
	if err != nil {
		log.Fatalln("[election] cannot end election ", err)
	}

	file, _ = json.MarshalIndent(elect, "", "  ")
	err = ioutil.WriteFile(elect.fname, file, 0775)
	if err != nil {
		log.Fatalln("[election] cannot end election ", err)
	}

	elect.Status = DONE

	log.Println("[election] end election " + elect.Name)
	log.Println("[election] result")
	log.Printf("%d voted\n", len(result.Users))

	res := getResult(elect)
	for name, count := range res {
		log.Printf("%s - %d\n", name, count)
	}

	return true
}

func getResult(elect *Election) map[string]int {
	res := make(map[string]int)

	if elect.Status != DONE {
		return nil
	}

	for _, name := range elect.Candidates {
		res[name] = 0
	}

	for _, name := range elect.votes {
		res[name]++
	}

	return res
}

func vote(id int, pw string, eid int, candi string) bool {
	if !checkUserPassword(id, pw) {
		return false
	}

	elect := getElection(eid)
	if elect == nil {
		return false
	}

	for _, user := range elect.users {
		if user == id {
			return false
		}
	}

	elect.users = append(elect.users, id)
	elect.votes = append(elect.votes, candi)

	return true
}
