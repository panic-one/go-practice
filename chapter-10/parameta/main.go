package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

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
	fmt.Fprintf(w, "%sさんの運勢は%sです。\n", r.FormValue("p"), omikuji())
}

func main() {
	http.HandleFunc("/", omikujiHandler)
	http.ListenAndServe(":8080", nil)
}
