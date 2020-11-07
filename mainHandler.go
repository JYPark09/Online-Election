package main

import (
	"html/template"
	"net/http"
)

func mainHandler(writer http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("template/index.html")

	var data struct {
		During []Election
		Ready  []Election
		Done   []Election
	}

	for _, election := range elections {
		if election.Status == DURING {
			data.During = append(data.During, election)
		} else if election.Status == DONE {
			data.Done = append(data.Done, election)
		} else if election.Status == READY {
			data.Ready = append(data.Ready, election)
		}
	}

	t.Execute(writer, &data)
}
