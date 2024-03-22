package main

import (
	"fmt"
	"os"
	"flag"
)
var number = flag.Int("n","1", "行番号")
var n int 
func init() {
	flag.IntVar(&n, "n", 1, "行番号")
}
func main() {
	filepath := os.Args[1:]
	for _, a := range filepath {
		sf, err := os.Open(a)
		if err != nil {
			fmt.Fprintln(os.Stdout,a)
		}
		defer sf.Close
	}


	flag.Parse()
	fmt.Println(strings.Repeat_(*number, n))
}
