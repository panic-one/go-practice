package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var n = flag.Bool("n", false, "行番号をつけるかどうか")
var linenumber int = 1

func main() {
	flag.Parse()
	filepath := os.Args[1:]
	for _, a := range filepath {
		sf, err := os.Open(a)
		if err != nil {
			fmt.Fprintln(os.Stderr, "ファイル読み込みに失敗しました:", err)
		}
		readLines(sf)
		defer sf.Close()
	}
}

func readLines(fp *os.File) {
	scanner := bufio.NewScanner(fp)
	// 1行ずつ読み込んで繰り返す
	for scanner.Scan() {
		//1行分を出力する
		fmt.Println(scanner.Text())
		fmt.Printf("%5d: %s\n", linenumber, scanner.Text())
		linenumber ++
	}
	// まとめてエラー処理をする
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "読み込みに失敗しました:", err)
	}
}
