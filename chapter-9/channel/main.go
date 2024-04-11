package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func input(r io.Reader) <-chan string {
	// TODO: チャネルを作る
	ch1 := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			// TODO: チャネルに読み込んだ文字列を送る
			ch1 <- s.Text()
		}
		// TODO: チャネルを閉じる
		defer close(ch1)
	}()
	// TODO: チャネルを返す
	return ch1
}

func main() {
	ch := input(os.Stdin)
	for {
		fmt.Print(">")
		fmt.Println(<-ch)
	}
}
