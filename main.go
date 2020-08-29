package main

import (
	"algo-3/tasks"
	tester "github.com/MihailShev/algo-tester"
)

const pathToKingTask = "test-data/king"
const pathToHorseTask = "test-data/horse"
const pathToFenTask = "test-data/fen"

func main() {

	kingTask := tester.NewTester(tasks.King{}, pathToKingTask)
	kingTask.RunTestWithCount(10)

	horseTask := tester.NewTester(tasks.Horse{}, pathToHorseTask)
	horseTask.RunTestWithCount(10)

	fenTask := tester.NewTester(tasks.Fen{}, pathToFenTask)
	fenTask.RunTestWithCount(10)
}
