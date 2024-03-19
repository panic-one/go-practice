package main

func main() {
	sum := 0
	ns := []int{19, 86, 1, 12}
	for _, i := range ns {
		sum += i
	}
	println(sum)
}
