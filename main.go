package main

import "fmt"

func main() {
	g := NewGame(false)
	totalGames := 0
	totalWin := 0
	for i := 1; i <= 1000; i++ {
		win := g.TestGame()
		if win {

			totalWin = totalWin + 1
		}
		totalGames = totalGames + 1
	}
	fmt.Println("Total total Games: ", totalGames)
	fmt.Println("Total total Win: ", totalWin)
	fmt.Println("Total percent: ", float32(totalWin)/float32(totalGames)*100)
}
