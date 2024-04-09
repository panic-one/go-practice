package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

var temp = template.Must(template.New("").
	Parse("<html><body>{{Name}}さんの運勢は「<b>{{Omikuji}}</b>」です</body></html>"))

type Result struct {
	Name    string
	Omikuji string
}

func omikuji() string {
	t := time.Now().UnixNano()
	rand.Seed(t)
	s := rand.Intn(6) + 1
	if s == 6 {
		return "大吉"
	} else if s == 5 || s == 4 {
		return "中吉"
	} else if s == 3 || s == 2 {
		return "吉"
	} else {
		return "凶"
	}
}

func omikujiHandler(w http.ResponseWriter, r *http.Request) {
	result := Result{
		Name:    r.FormValue("p"),
		Omikuji: omikuji(),
	}
	temp.Execute(w, result)
}

func main() {
	http.HandleFunc("/", omikujiHandler)
	http.ListenAndServe(":8080", nil)
}
