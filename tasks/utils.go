package tasks

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

func defineFigure(f string) int {
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
