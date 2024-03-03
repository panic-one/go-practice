package main

import (
	"fmt"
	"math/rand"
	"time"
)

func saikoro() int {
	t := time.Now().UnixNano()
    rand.Seed(t)
    s := rand.Intn(6) + 1
    println(s)
	return s
}

func main() {
	s := saikoro()
	if  s == 6 {
		fmt.Println(s,"大吉")
	} else if s == 5 || s == 4 {
		fmt.Println(s,"中吉")
	} else if s == 3 || s == 2 {
		fmt.Println(s,"吉")
	} else {
		fmt.Println(s,"凶")
	}
}