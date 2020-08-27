package tasks

import (
	"fmt"
	"strconv"
)

type King struct {
}

func (k King) Run(data []string) string {
	x, err := strconv.ParseInt(data[0], 10, 64)

	if err != nil {
		return ""
	}

	return k.calc(uint64(x))
}

func (k King) calc(x uint64) string {
	king := uint64(1) << x

	kingL := king & uint64(0xFEFEFEFEFEFEFEFE)
	kingR := king & uint64(0x7F7F7F7F7F7F7F7F)

	mask := (kingL << 7) | (king << 8) | (kingR << 9) |
		    (kingL >> 1) |               (kingR << 1) |
			(kingL >> 9) | (king >> 8) | (kingR >> 7)

	count := 0
	tmp := mask
	for tmp > 0 {
		count++

		// Считаем количество единичных битов
		tmp &= tmp - 1
	}

	return fmt.Sprintf("%d %d", count, mask)
}
