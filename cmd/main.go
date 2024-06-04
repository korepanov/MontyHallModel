package main

import (
	"fmt"
	"math/rand"
)

const experimentsNumber = 1000000

func main() {
	noChangeStrategyWins := 0
	changeStrategyWins := 0

	randomGenerator := rand.New(rand.NewSource(42))

	for i := 0; i < experimentsNumber; i++ {
		winnerNumber := randomGenerator.Intn(3)

		doors := make(map[int]int)
		doors[0] = 1
		doors[1] = 1
		doors[2] = 1
		doors[winnerNumber] = 2

		selectedDoorNumber := randomGenerator.Intn(3)
		var remainingDoorsNumbers []int

		for idx := range doors {
			if idx != selectedDoorNumber {
				remainingDoorsNumbers = append(remainingDoorsNumbers, idx)
			}
		}

		indexToOpen := randomGenerator.Intn(2)

		if doors[remainingDoorsNumbers[indexToOpen]] != 2 {
			delete(doors, remainingDoorsNumbers[indexToOpen])
		} else {
			remainingDoorsNumbers = append(remainingDoorsNumbers[:indexToOpen], remainingDoorsNumbers[indexToOpen+1:]...)
			delete(doors, remainingDoorsNumbers[0])
		}

		if doors[selectedDoorNumber] == 2 {
			noChangeStrategyWins++
			continue
		}

		delete(doors, selectedDoorNumber)

		for _, val := range doors {
			if val == 2 {
				changeStrategyWins++
			}
		}

	}

	fmt.Println("wins with no change: " + fmt.Sprintf("%d%%", int((float64(noChangeStrategyWins)/experimentsNumber)*100)))
	fmt.Println("wins with change: " + fmt.Sprintf("%d%%", int((float64(changeStrategyWins)/experimentsNumber)*100)))

}
