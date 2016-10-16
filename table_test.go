package csvp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func testTable() Table {
	return Table{
		cells: [][]Cell{
			[]Cell{*NewCell("1-1"), *NewCell("2-1"), *NewCell("3-1")},
			[]Cell{*NewCell("1-2"), *NewCell("2-2"), *NewCell("3-2")},
		},
	}
}

func Test_ColumnCount(t *testing.T) {
	table := testTable()
	assert.Equal(t, 3, table.ColumnCount(), "Column count incorrect")
}

func Test_RowCount(t *testing.T) {
	table := testTable()
	assert.Equal(t, 2, table.RowCount(), "Row count incorrect")
}

func Test_rangeDescriptorB1(t *testing.T) {
	testRangeDescriptor(
		t,
		"C1",
		tableRange{left: 2, right: 2, upper: 0, lower: 0},
	)
}

func Test_rangeDescriptorB1B2(t *testing.T) {
	testRangeDescriptor(
		t,
		"B1:B2",
		tableRange{left: 1, right: 1, upper: 0, lower: 1},
	)
}

func Test_rangeDescriptorB1B(t *testing.T) {
	testRangeDescriptor(
		t,
		"B1:B",
		tableRange{left: 1, right: 1, upper: 0, lower: 1},
	)
}

func Test_rangeDescriptorB2C(t *testing.T) {
	testRangeDescriptor(
		t,
		"B2:C",
		tableRange{left: 1, right: 2, upper: 1, lower: 1},
	)
}

func Test_rangeDescriptorA2_2(t *testing.T) {
	testRangeDescriptor(
		t,
		"A2:2",
		tableRange{left: 0, right: 2, upper: 1, lower: 1},
	)
}

func Test_rangeDescriptorB2A1(t *testing.T) {
	testRangeDescriptor(
		t,
		"C2:B1",
		tableRange{left: 1, right: 2, upper: 0, lower: 1},
	)
}

func Test_Subtable(t *testing.T) {
	table := testTable()
	subtable, err := table.Subtable("A1:A")

	assert.NoError(t, err)
	assert.NotNil(t, subtable)
	assert.Equal(t, 1, subtable.ColumnCount())
	assert.Equal(t, 2, subtable.RowCount())
	assert.Equal(t, "1-1", subtable.cells[0][0].content)
	assert.Equal(t, "1-2", subtable.cells[1][0].content)
}

func Test_Subtable2(t *testing.T) {
	table := testTable()
	subtable, err := table.Subtable("A2:2")

	assert.NoError(t, err)
	assert.NotNil(t, subtable)
	assert.Equal(t, 3, subtable.ColumnCount())
	assert.Equal(t, 1, subtable.RowCount())
	assert.Equal(t, "1-2", subtable.cells[0][0].content)
	assert.Equal(t, "2-2", subtable.cells[0][1].content)
	assert.Equal(t, "3-2", subtable.cells[0][2].content)
}

func testRangeDescriptor(t *testing.T, desc string, expected tableRange) {
	table := testTable()
	tr, err := table.rangeForDescriptor(desc)
	assert.NoError(t, err, `Column descriptor "%s" should have worked`, desc)
	assert.Equal(t, expected, tr, `Column descriptor "%s" should refer to %#v, got %#v`, desc, expected, tr)
}
