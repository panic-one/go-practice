package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var n = flag.Bool("n", false, "行番号をつけるかどうか")

func main() {
	flag.Parse()
	linenumber := 1
	filepath := flag.Args()
	for _, a := range filepath {
		sf, err := os.Open(a)
		if err != nil {
			fmt.Fprintln(os.Stderr, "ファイル読み込みに失敗しました:", err)
		}
		linenumber = readLines(sf, n, linenumber)
		defer sf.Close()
	}
}

func readLines(fp *os.File, n *bool, s int) int {
	scanner := bufio.NewScanner(fp)
	// 1行ずつ読み込んで繰り返す
	for scanner.Scan() {
		//1行分を出力する
		if *n {
			fmt.Printf("%4d: %s\n", s, scanner.Text())
			s++
		} else {
			fmt.Println(scanner.Text())
		}
	}
	// まとめてエラー処理をする
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "読み込みに失敗しました:", err)
	}
	return s
}
