package main

func swap2(x, y *int) {
	tmp := *x
	*x = *y
	*y = tmp
}

func main() {
	n, m := 10, 20

	swap2(&n, &m)
	println(n, m)
}
