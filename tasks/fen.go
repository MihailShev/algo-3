package tasks

import (
	"algo-3/utils"
	"fmt"
	"strings"
)

type Fen struct {
}

func (f Fen) Run(data []string) string {
	position := data[0]

	return f.calc(position)
}

func (f Fen) calc(position string) string {
	boards := utils.ParseFen(position)

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
