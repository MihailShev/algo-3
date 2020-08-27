package main

import (
	"algo-3/tasks"
	tester "github.com/MihailShev/algo-tester"
)

const pathToKingTask = "test-data/king"
const pathToHorseTask = "test-data/horse"

func main() {

	//kingTask := tester.NewTester(tasks.King{}, pathToKingTask)
	//kingTask.RunTestWithCount(10)

	horseTask := tester.NewTester(tasks.Horse{}, pathToHorseTask)
	horseTask.RunTestWithCount(10)
}
