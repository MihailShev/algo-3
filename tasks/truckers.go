package tasks

import (
	"algo-3/utils"
	"fmt"
)

type Truckers struct {
}

func (t Truckers) Run(data []string) string {
	position := data[0]

	return t.Calc(position)
}

func (t Truckers) Calc(position string) string {
	boards := utils.ParseFen(position)
	res := fmt.Sprintf("%d\r\n%d\r\n%d", t.rookSteps(boards), t.bishopSteps(boards),
		t.queenSteps(boards))

	return res
}

func (t Truckers) rookSteps(boards []uint64) uint64 {
	start := boards[utils.WhiteRook]
	whiteFigures, blackFigures := utils.FillFigureMasks(boards, utils.WhiteRook)

	return utils.VerticalSteps(start, whiteFigures, blackFigures)
}

func (t Truckers) queenSteps(boards []uint64) uint64 {
	start := boards[utils.WhiteQueen]
	whiteFigures, blackFigures := utils.FillFigureMasks(boards, utils.WhiteQueen)

	verticalSteps := utils.VerticalSteps(start, whiteFigures, blackFigures)
	diagonalSteps := utils.DiagonalSteps(start, whiteFigures, blackFigures)

	return verticalSteps | diagonalSteps
}

func (t Truckers) bishopSteps(boards []uint64) uint64 {
	start := boards[utils.WhiteBishop]
	whiteFigures, blackFigures := utils.FillFigureMasks(boards, utils.WhiteBishop)

	return utils.DiagonalSteps(start, whiteFigures, blackFigures)
}
