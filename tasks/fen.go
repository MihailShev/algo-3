package tasks

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const binary = float64(2)
const lineLength = 8
const fieldCount = 64

type Fen struct {
}

func (f Fen) Run(data []string) string {
	position := data[0]

	return f.calc(position)
}

func (f Fen) calc(position string) string {
	boards := []uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	handled := strings.Split(position, "/")

	fieldNum := fieldCount

	for _, s := range handled {
		pow := fieldNum - lineLength
		fieldNum -= lineLength

		for _, v := range s {

			n, err := strconv.Atoi(string(v))

			if err != nil {
				f := defineFigure(string(v))
				boards[f] = uint64(math.Pow(binary, float64(pow))) | boards[f]
				pow++
			} else {
				pow += n
			}
		}
	}

	return f.convertResToStr(boards)

}

func (f Fen) convertResToStr(boards []uint64) string {
	var res strings.Builder

	for i, v := range boards {
		res.WriteString(fmt.Sprint(v))
		if i != len(boards)-1 {
			res.WriteString("\r\n")
		}
	}

	return res.String()
}
