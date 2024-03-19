package main

import "fmt"

func main() {
	sum := float32(5 + 6 + 3)
	avg := sum / 3
	if avg > 4.5 {
		fmt.Println("good")
	}
}