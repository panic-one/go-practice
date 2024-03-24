package main

import (
	"flag"
	"fmt"
	"go-practice/chapter-5/gazoucomand/gazou"
)

var inputformat = flag.String("i", "jpg", "変換前の画像の形式の指定")
var outputformat = flag.String("o", "png", "変換後の画像の形式の指定")

func main() {
	flag.Parse()
	dir := flag.Arg(0)
	g := gazou.NewGazou()
	if err := g.ConvertJPGToPNG(dir, *inputformat, *outputformat); err != nil {
		fmt.Println("エラー", err)
	}
}
