package csvp

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

var columnNamePattern = `[A-Z]+`
var rowNamePattern = `\d+`

var columnNameRegexp = regexp.MustCompile(columnNamePattern)
var rangeRegexp = regexp.MustCompile(fmt.Sprintf(`(%s)(%s)(?:\:(%s)?(%s)?)?`, columnNamePattern, rowNamePattern, columnNamePattern, rowNamePattern))

// Essentially just a matrix of cells, indexed by row first
type Table struct {
	cells [][]Cell
}

type tableRange struct {
	upper int
	lower int
	left  int
	right int
}

var InvalidTableRange = tableRange{-1, -1, -1, -1}

// Returns the number of columns in the table
// If the table is empty (has 0 rows), this function will return 0.
// It will not return 0 under any other circumstance.
func (table Table) ColumnCount() int {
	if table.IsEmpty() {
		return 0
	} else {
		return len(table.cells[0])
	}
}

// Returns the number of rows in the table
func (table Table) RowCount() int {
	return len(table.cells)
}

// If the table is empty, it has 0 cells
func (table Table) IsEmpty() bool {
	return len(table.cells) == 0
}

func (table Table) columnForDescriptor(descriptor string) (int, error) {
	if !columnNameRegexp.MatchString(descriptor) {
		return -1, fmt.Errorf(`Incorrect column descriptor "%s"`, descriptor)
	}
	index := 0
	for _, rune := range descriptor {
		if rune >= 'A' && rune <= 'Z' {
			index = index*int('Z'-'A') + int(rune-'A')
		} else {
			return -1, fmt.Errorf("Column not identified for '%s'", descriptor)
		}
	}
	if index >= table.ColumnCount() {
		return -1, fmt.Errorf("Column index out of range")
	}
	return index, nil
}

func (t Table) Subtable(desc string) (*Table, error) {
	tr, err := t.rangeForDescriptor(desc)
	if err != nil {
		return nil, err
	}
	ret := new(Table)

	retRows := tr.lower - tr.upper + 1
	ret.cells = make([][]Cell, retRows)
	for i := 0; i < retRows; i++ {
		ret.cells[i] = t.cells[tr.upper+i][tr.left : tr.right+1]
	}
	return ret, nil
}

func (table Table) rangeForDescriptor(descriptor string) (tableRange, error) {
	match := rangeRegexp.FindStringSubmatch(descriptor)

	var col1, row1, col2, row2 int
	var err error
	col1, err = table.columnForDescriptor(match[1])
	if err != nil {
		return InvalidTableRange, err
	}

	if match[2] != "" {
		row1, err = strconv.Atoi(match[2])
		if err != nil {
			return InvalidTableRange, err
		}
		row1 -= 1
	} else {
		row1 = 0
	}

	if match[3] != "" {
		col2, err = table.columnForDescriptor(match[3])
		if err != nil {
			return InvalidTableRange, err
		}
	} else {
		if match[4] == "" {
			col2 = col1
		} else {
			col2 = table.ColumnCount() - 1
		}
	}

	if match[4] != "" {
		row2, err = strconv.Atoi(match[4])
		if err != nil {
			return InvalidTableRange, err
		}
		row2 -= 1
	} else {
		if match[3] == "" {
			row2 = row1
		} else {
			row2 = table.RowCount() - 1
		}
	}

	if col1 > col2 {
		col1, col2 = col2, col1
	}

	if row1 > row2 {
		row1, row2 = row2, row1
	}

	if row1 < 0 || row2 >= table.RowCount() || col1 < 0 || col2 >= table.ColumnCount() {
		return InvalidTableRange, errors.New("Range out of bounds")
	}

	return tableRange{left: col1, upper: row1, right: col2, lower: row2}, nil
}
