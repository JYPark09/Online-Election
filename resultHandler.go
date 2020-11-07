package main

import (
	"html/template"
	"net/http"
	"strconv"
)

func resultHandler(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(request.FormValue("id"))

	if err != nil {
		produceMsg(writer, "알 수 없는 오류가 발생했습니다.")
		return
	}

	elect := getElection(id)
	if elect == nil {
		produceMsg(writer, "알 수 없는 오류가 발생했습니다.")
		return
	}

	result := getResult(elect)

	var winner []string
	sum := 0.0
	maxValue := -999

	for name, value := range result {
		sum += float64(value)

		if value >= maxValue {
			if value > maxValue {
				winner = nil
			}

			winner = append(winner, name)
			maxValue = value
		}
	}

	type resultT struct {
		Candidate string
		Count     int
		Percent   int
	}

	var data struct {
		Name string

		Winner []string

		Result []resultT
	}

	data.Name = elect.Name

	data.Winner = winner

	for name, value := range result {
		var res resultT

		res.Candidate = name
		res.Count = value
		res.Percent = int(float64(value) / sum * 100.0)

		data.Result = append(data.Result, res)
	}

	t, _ := template.ParseFiles("template/result.html")

	t.Execute(writer, &data)
}
