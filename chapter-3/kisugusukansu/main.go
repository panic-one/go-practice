package main

func kisu(number int) bool {
    if number%2 == 0 {
		return false
	} else {
        return true
    }
}


func main() {
	for i := 1; i <= 100; i++ {
		
		print(i)
		if kisu(i) {
			println("-偶数")
		} else {
            println("-奇数")
		}
	}
}
