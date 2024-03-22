package main

import (
	"fmt"
	"os"
	"flag"
)
var number = flag.Int("n",0, "行番号")
var n int 
func init() {
	flag.IntVar(&n, "n", 0, "行番号")
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

if number >= 1 {
	flag.Parse()
	fmt.Println(int.Repeat_(*number, n))
}
}
