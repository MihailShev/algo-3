package utils

import (
	"math"
	"strconv"
	"strings"
)

const binary = float64(2)
const lineLength = 8
const fieldCount = 64

type WhiteFigures = uint64
type BlackFigures = uint64
type Figure = int

const leftWall = uint64(0x0101010101010101)
const rightWall = uint64(0x8080808080808080)

const (
	WhitePawn   = iota // белая пешка
	WhiteKnight = iota // белый конь
	WhiteBishop = iota // белый офицер
	WhiteRook   = iota // белая ладья
	WhiteQueen  = iota // белая королева
	WhiteKing   = iota // белый король

	BlackPawn   = iota // черная пешка
	BlackKnight = iota // черный конь
	BlackBishop = iota // черный офицер
	BlackRook   = iota // черная ладья
	BlackQueen  = iota // черная королева
	BlackKing   = iota // черный король
)

func DefineFigure(f string) int {
	switch f {
	case "r":
		return BlackRook
	case "n":
		return BlackKnight
	case "b":
		return BlackBishop
	case "q":
		return BlackQueen
	case "k":
		return BlackKing
	case "p":
		return BlackPawn
	case "R":
		return WhiteRook
	case "N":
		return WhiteKnight
	case "B":
		return WhiteBishop
	case "Q":
		return WhiteQueen
	case "K":
		return WhiteKing
	case "P":
		return WhitePawn
	}

	return -1
}

// Парсер FEN нотации
func ParseFen(position string) []uint64 {
	boards := []uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	handled := strings.Split(position, "/")

	fieldNum := fieldCount

	for _, s := range handled {
		pow := fieldNum - lineLength
		fieldNum -= lineLength

		for _, v := range s {

			n, err := strconv.Atoi(string(v))

			if err != nil {
				f := DefineFigure(string(v))
				boards[f] = uint64(math.Pow(binary, float64(pow))) | boards[f]
				pow++
			} else {
				pow += n
			}
		}
	}

	return boards
}

// Расчет вертикальных шагов для ферзя и ладьи
func VerticalSteps(start uint64, whiteFigures, blackFigures uint64) uint64 {
	verticalUp := calcVerticalUp(start, whiteFigures, blackFigures)
	verticalDown := calcVerticalDown(start, whiteFigures, blackFigures)
	verticalLeft := calcVerticalLeft(start, whiteFigures, blackFigures)
	verticalRight := calcVerticalRight(start, whiteFigures, blackFigures)

	return verticalRight | verticalLeft | verticalDown | verticalUp
}

// Расчет диагональных шагов для ферзя и слона
func DiagonalSteps(start uint64, whiteFigures, blackFigures uint64) uint64 {

	diagonalLeftUp := calcDiagonalLeftUp(start, whiteFigures, blackFigures)
	diagonalRightUp := calcDiagonalRightUp(start, whiteFigures, blackFigures)
	diagonalLeftDown := calcDiagonalLeftDown(start, whiteFigures, blackFigures)
	diagonalRightDown := calcDiagonalRightDown(start, whiteFigures, blackFigures)

	return diagonalRightUp | diagonalLeftUp | diagonalLeftDown | diagonalRightDown
}

// Создает две битовые макски со всеми белыми фигурами (за исключением одной белой фигуры) и  со всеми черными фигурами
func FillFigureMasks(boards []uint64, exclude Figure) (WhiteFigures, BlackFigures) {
	var white WhiteFigures = 0
	var black BlackFigures = 0

	for i := 0; i < len(boards); i++ {
		if i == exclude {
			continue
		}

		if i < 6 {
			white |= boards[i]
		} else {
			black |= boards[i]
		}
	}

	return white, black
}

// Расчет возможных шагов по вертикали вверх
func calcVerticalUp(field, whiteFigures, blackFigures uint64) uint64 {
	res := uint64(0)

	for field > 0 {
		field <<= 8

		if isWhiteFigureHere(field, whiteFigures) {
			break
		}

		res = res | field

		if isBlackFigureHere(field, blackFigures) {
			break
		}
	}

	return res
}

// Расчет возможных шагов по вертикали вниз
func calcVerticalDown(field, whiteFigures, blackFigures uint64) uint64 {
	res := uint64(0)

	for field > 0 {
		field >>= 8

		if isWhiteFigureHere(field, whiteFigures) {
			break
		}

		res = res | field

		if isBlackFigureHere(field, blackFigures) {
			break
		}
	}

	return res
}

// Расчет возможных шагов по вертикали влево
func calcVerticalLeft(field, whiteFigures, blackFigures uint64) uint64 {
	res := uint64(0)

	for field > 0 {

		if isLeftWall(field) {
			break
		}

		field >>= 1

		if isWhiteFigureHere(field, whiteFigures) {
			break
		}

		res = res | field

		if isBlackFigureHere(field, blackFigures) {
			break
		}
	}

	return res
}

// Расчет возможных шагов по вертикали вправо
func calcVerticalRight(field, whiteFigures, blackFigures uint64) uint64 {
	res := uint64(0)

	for field > 0 {

		if isRightWall(field) {
			break
		}

		field <<= 1

		if isWhiteFigureHere(field, whiteFigures) {
			break
		}

		res = res | field

		if isBlackFigureHere(field, blackFigures) {
			break
		}
	}

	return res
}

// Расчет возможных шагов по диагонали влево и вверх
func calcDiagonalLeftUp(field, whiteFigures, blackFigures uint64) uint64 {
	res := uint64(0)

	for field > 0 {

		if isLeftWall(field) {
			break
		}

		field <<= 7

		if isWhiteFigureHere(field, whiteFigures) {
			break
		}

		res = res | field

		if isBlackFigureHere(field, blackFigures) {
			break
		}
	}

	return res
}

// Расчет возможных шагов по диагонали влево и вниз
func calcDiagonalLeftDown(field, whiteFigures, blackFigures uint64) uint64 {
	res := uint64(0)

	for field > 0 {

		if isLeftWall(field) {
			break
		}

		field >>= 9

		if isWhiteFigureHere(field, whiteFigures) {
			break
		}

		res = res | field

		if isBlackFigureHere(field, blackFigures) {
			break
		}
	}

	return res
}

// Расчет возможных шагов по диагонали вправо и вверх
func calcDiagonalRightUp(field, whiteFigures, blackFigures uint64) uint64 {
	res := uint64(0)

	for field > 0 {

		if isRightWall(field) {
			break
		}

		field <<= 9

		if isWhiteFigureHere(field, whiteFigures) {
			break
		}

		res = res | field

		if isBlackFigureHere(field, blackFigures) {
			break
		}
	}

	return res
}

// Расчет возможных шагов по диагонали вправо и вниз
func calcDiagonalRightDown(field, whiteFigures, blackFigures uint64) uint64 {
	res := uint64(0)

	for field > 0 {

		if isRightWall(field) {
			break
		}

		field >>= 7

		if isWhiteFigureHere(field, whiteFigures) {
			break
		}

		res = res | field

		if isBlackFigureHere(field, blackFigures) {
			break
		}
	}

	return res
}

// Проверяет есть ли на клетке белая фигура
func isWhiteFigureHere(field uint64, whiteFigures WhiteFigures) bool {
	return field&whiteFigures == field
}

// Проверяет есть ли на клетке черная фигура
func isBlackFigureHere(field uint64, blackFigures BlackFigures) bool {
	return field&blackFigures == field
}

// Проверяет находится ли поле у левой стены
func isLeftWall(field uint64) bool {
	return field&leftWall == field
}

// Проверяет находится ли поле у правой стены
func isRightWall(field uint64) bool {
	return field&rightWall == field
}
