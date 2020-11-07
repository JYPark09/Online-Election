package main

import (
	"html/template"
	"net/http"
	"strconv"
)

func electViewHandler(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(request.FormValue("id"))

	if err != nil || len(elections) <= id {
		produceMsg(writer, "알 수 없는 오류가 발생했습니다.")
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

	if !checkUserPassword(id, pw) {
		produceMsg(writer, "올바르지 않은 사용자 정보입니다.")
		return
	}
	//candi := request.PostForm.Get("candi")

	produceMsg(writer, "투표가 완료되었습니다.")
}
