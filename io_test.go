package csvp

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func genReader(s string) *TableReader {
	return NewTableReader(strings.NewReader(s))
}

func Test_ParseNormalCSV(t *testing.T) {
	reader := genReader("A1,B1\nA2,B2\n")
	assert.NotNil(t, reader)

	table, err := reader.ReadTable()

	assert.NoError(t, err)
	assert.NotNil(t, table)
	assert.Equal(t, 2, table.RowCount())
	assert.Equal(t, 2, table.ColumnCount())
}

func Test_ParseEmptyCSV(t *testing.T) {
	reader := genReader("")
	assert.NotNil(t, reader)

	table, err := reader.ReadTable()

	assert.NoError(t, err)
	assert.NotNil(t, table)
	assert.Equal(t, 0, table.RowCount())
	assert.Equal(t, 0, table.ColumnCount())
	assert.True(t, table.IsEmpty())
}

func Test_ParseInconsistentRowsCSV(t *testing.T) {
	reader := genReader("A1\nA1,A2")
	assert.NotNil(t, reader)

	_, err := reader.ReadTable()

	assert.Error(t, err)
}
