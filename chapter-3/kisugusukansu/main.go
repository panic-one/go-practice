package main

func gusu(number int) bool {
    if number%2 == 0 {
		return true
	} else {
        return false
    }
}


func main() {
	for i := 1; i <= 100; i++ {
		
		print(i)
		if gusu(i) {
			println("-偶数")
		} else {
            println("-奇数")
		}
	}
}
