package main 

import (
	"fmt"
	"os"
	"bufio"
	"io"
	"time"
	"strconv"
	
	"github.com/tjarratt/babble"
)

func input(r io.Reader) <-chan string{
	channel := make(chan string)
	go func() {
		strings := bufio.NewScanner(r)
		for strings.Scan() {
			channel <- strings.Text()
		}
	}()
	return channel
}

func words() string{
	babbler := babble.NewBabbler()
    babbler.Count = 1
	return babbler
}

func main() {
	ch_rcv := input(os.Stdin)
	n := 0
	t := 1
	t1 :=time.After(time.Duration(t) * time.Minute)
	q := words()
	select {
        case <-t1: // 制限時間が来た時の処理
            fmt.Println("Finished!" + " Your score is " + strconv.Itoa(n) + " points! Good job:)")
            break words()
        case x := <-ch_rcv: // お題と入力した値の比較処理
            if x == q {
                fmt.Println("OK!")
                n += 1
            } else {
                fmt.Println("NG :(")
            }
        }
}