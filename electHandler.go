package main

import (
	"html/template"
	"net/http"
	"strconv"
)

func electViewHandler(writer http.ResponseWriter, request *http.Request) {
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

	if elect.Status != DURING {
		produceMsg(writer, "투표 가능 시간이 아닙니다.")
		return
	}

	t, _ := template.ParseFiles("template/election.html")

	t.Execute(writer, &elections[id])
}

func electHandler(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()

	id, err := strconv.Atoi(request.PostForm.Get("userid"))
	if err != nil {
		produceMsg(writer, "알 수 없는 오류가 발생했습니다.")
		return
	}

	pw := request.PostForm.Get("passwd")

	eid, err := strconv.Atoi(request.PostForm.Get("eid"))
	if err != nil {
		produceMsg(writer, "알 수 없는 오류가 발생했습니다.")
		return
	}

	candi := request.PostForm.Get("candi")

	if vote(id, pw, eid, candi) {
		produceMsg(writer, "투표가 완료되었습니다.")
	} else {
		produceMsg(writer, "투표에 실패하였습니다.")
	}
}
