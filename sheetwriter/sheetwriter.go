package sheetwriter

import (
	"fmt"

	"google.golang.org/api/sheets/v4"
)

type AxisX struct {
	Letter string
	Alias  string
}

type AxisY struct {
	Row   int
	Alias string
}

type CellCoordinate struct {
	X AxisX
	Y AxisY
}

type CellWriter interface {
	Value() string
}

type SheetData struct {
	Cells map[CellCoordinate]CellWriter
}

func (s *SheetData) ToValueRanges(sheetName string) []*sheets.ValueRange {
	var ranges []*sheets.ValueRange
	for coord, cell := range s.Cells {
		ranges = append(ranges, &sheets.ValueRange{
			Range:  fmt.Sprintf("%s!%s%d", sheetName, coord.X.Letter, coord.Y.Row),
			Values: [][]any{{cell.Value()}},
		})
	}
	return ranges
}
