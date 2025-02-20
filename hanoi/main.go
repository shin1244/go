package main

import "fmt"

func hanoi(n, start, end, temp int) int {
	if n == 1 {
		return 1
	}
	moves := hanoi(n-1, start, temp, end)
	moves++
	moves += hanoi(n-1, temp, end, start)
	return moves
}

func main() {
	totalMoves := hanoi(40, 1, 3, 2)
	fmt.Println(totalMoves)
}
