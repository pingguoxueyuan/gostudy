package main

import "fmt"

var (
	coins = 100
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func calcCoin(username string) int {
	var sum int = 0
	for _, char := range username {
		switch char {
		case 'a', 'A':
			sum = sum + 1
		case 'e', 'E':
			sum = sum + 1
		case 'i', 'I':
			sum = sum + 2
		case 'o', 'O':
			sum = sum + 3
		case 'u', 'U':
			sum = sum + 5
		}
	}

	return sum
}

func dispatchCoin() int {
	var left int = coins
	for _, username := range users {
		allCoin := calcCoin(username)
		left = left - allCoin
		value, ok := distribution[username]
		if !ok {
			distribution[username] = allCoin
		} else {
			distribution[username] = value + allCoin
		}
	}
	return left
}

func main() {

	left := dispatchCoin()
	for username, coin := range distribution {
		fmt.Printf("user:%s have %d coins\n", username, coin)
	}

	fmt.Printf("left coin:%d\n", left)
}
