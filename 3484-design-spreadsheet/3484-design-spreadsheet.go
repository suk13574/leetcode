import (
	"strconv"
	"strings"
)

type Spreadsheet struct {
	sheet map[byte]map[int]int
}

func Constructor(rows int) Spreadsheet {
	s := Spreadsheet{sheet: make(map[byte]map[int]int, 26)}

	for i := 0; i < 26; i++ {
		s.sheet['A'+byte(i)] = make(map[int]int)
	}

	return s
}

func (s *Spreadsheet) SetCell(cell string, value int) {
	label, row := getLabelRow(cell)

	(*s).sheet[label][row] = value
}

func (s *Spreadsheet) ResetCell(cell string) {
	label, row := getLabelRow(cell)

	(*s).sheet[label][row] = 0
}

func (s *Spreadsheet) GetValue(formula string) int {
	getCell := func(cell string) int {
		label, row := getLabelRow(cell)
		return (*s).sheet[label][row]
	}

	fs := strings.Split(formula[1:], "+")

	X, ok1 := strconv.Atoi(fs[0])
	if ok1 != nil {
		X = getCell(fs[0])
	}

	Y, ok2 := strconv.Atoi(fs[1])
	if ok2 != nil {
		Y = getCell(fs[1])
	}
	return X + Y
}

func getLabelRow(cell string) (byte, int) {
	label := cell[0]
	row, _ := strconv.Atoi(cell[1:])

	return label, row
}
