package main

import (
	"fmt"
	"time"

	greeting "github.com/tenntenn/greeting"
	greeting2 "github.com/tenntenn/greeting/v2"
)

func main() {
	fmt.Println(greeting.Do())
	fmt.Println(greeting2.Do(time.Now()))
}
