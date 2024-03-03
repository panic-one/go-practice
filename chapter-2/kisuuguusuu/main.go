package main
import "fmt"

func main() {
	for i := 0; i <=100; i=i+1{
	    if a := i%2; a == 0 {
	        fmt.Println(i,"偶数")
		} else {
			fmt.Println(i,"奇数")
		}
	}
}