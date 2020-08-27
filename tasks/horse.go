package tasks

import (
	"fmt"
	"strconv"
)

type Horse struct {
}

func (h Horse) Run(data []string) string {
	x, err := strconv.ParseInt(data[0], 10, 64)

	if err != nil {
		return ""
	}

	return h.calc(uint64(x))
}

func (h Horse) calc(x uint64) string {
	horse := uint64(1) << x
	horseL := horse & uint64(0xFEFEFEFEFEFEFEFE)
	horseR := horse & uint64(0x7F7F7F7F7F7F7F7F)

	mask := (horseL << 15) | (horseR << 17)  |
			(horseL << 6)  | (horseR << 10)  |
			(horseL >> 10) | (horseR >> 6)   |
			(horseL >> 17) | (horseR >> 15)

	count := 0
	tmp := mask
	for tmp > 0 {
		count++

		// Считаем количество единичных битов
		tmp &= tmp - 1
	}

	return fmt.Sprintf("%d %d", count, mask)
}
