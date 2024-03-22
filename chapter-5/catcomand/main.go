package main

import (
	"fmt"
	"os"
)

func main() {
	filepath := os.Args[1:]
	for _, a := range filepath {
		sf, err := os.Open(a)
		if err != nil {
			fmt.Fprintln(os.Stdout,a)
		}
		defer sf.Close
	}
}
