package main

import (
	"fmt"
	"os"
)

type Stringer interface {
	String() string
}

type Error string

func (e Error) Error() string {
	return string(e)
}

func ToStringer(v interface{}) (Stringer, error) {
	if a, ok := v.(Stringer); ok {
		return a, nil
	}
	return nil, Error("変換に失敗しました")
}

func main() {
	t := "aaa"
	if s, err := ToStringer(t); err != nil {
		fmt.Fprintln(os.Stderr, "err", err)
	} else {
		fmt.Println(s)
	}
}
