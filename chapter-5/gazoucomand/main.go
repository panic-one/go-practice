package main

import (
	"flag"
	"fmt"
	"go-practice/chapter-5/gazoucomand/gazou"
)

var inputFormat = flag.String("i", "jpg", "変換前の画像の形式の指定")
var outputFormat = flag.String("o", "png", "変換後の画像の形式の指定")

func main() {
	flag.Parse()
	dir := flag.Arg(0)
	g := gazou.NewGazou()
	if err := g.Convert(dir, *inputFormat, *outputFormat); err != nil {
		fmt.Println("エラー", err)
	}
}
