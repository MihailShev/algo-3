package main

import (
	"algo-3/tasks"
	"fmt"
	tester "github.com/MihailShev/algo-tester"
)

const pathToKingTask = "test-data/king"
const pathToHorseTask = "test-data/horse"
const pathToFenTask = "test-data/fen"
const pathToTruckersTask = "test-data/truckers"

func main() {

	fmt.Printf("*** Test King tasks ***\n\n")
	kingTask := tester.NewTester(tasks.King{}, pathToKingTask)
	kingTask.RunTestWithCount(10)

	fmt.Printf("\n*** Test horse tasks ***\n\n")
	horseTask := tester.NewTester(tasks.Horse{}, pathToHorseTask)
	horseTask.RunTestWithCount(10)

	fmt.Printf("\n*** Test FEN tasks ***\n\n")
	fenTask := tester.NewTester(tasks.Fen{}, pathToFenTask)
	fenTask.RunTestWithCount(10)

	fmt.Printf("\n*** Test Truckers tasks ***\n\n")
	truckersTask := tester.NewTester(tasks.Truckers{}, pathToTruckersTask)
	truckersTask.RunTestWithCount(10)

	fmt.Printf("\n*** Finish ***\n")
	_, _ = fmt.Scanf(" ")
}
