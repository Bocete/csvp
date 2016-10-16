package csvp

import (
	"encoding/csv"
	"errors"
	"io"
)

type TableReader struct {
	*csv.Reader
}

type TableWriter struct {
	*csv.Writer
}

func NewTableReader(r io.Reader) *TableReader {
	return &TableReader{
		Reader: csv.NewReader(r),
	}
}

func (tr TableReader) ReadTable() (*Table, error) {
	table := Table{}
	colCount := -1
	for {
		line, err := tr.Reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		if colCount == -1 {
			if len(line) == 0 {
				return nil, &csv.ParseError{Err: errors.New("Empty row in line 1"), Line: 1, Column: 0}
			}
			colCount = len(line)
		} else if colCount != len(line) {
			return nil, &csv.ParseError{Err: errors.New("Inconsistent column count")}
		}
		cells := make([]Cell, len(line))
		for i, content := range line {
			cells[i] = Cell{content: content}
		}
		table.cells = append(table.cells, cells)
	}
	return &table, nil
}

func NewTableWriter(r io.Writer) *TableWriter {
	return &TableWriter{
		Writer: csv.NewWriter(r),
	}
}

func (tw TableWriter) WriteTable(table Table) error {
	for _, row := range table.cells {
		strings := make([]string, len(row))
		for i, cell := range row {
			strings[i] = cell.content
		}
		err := tw.Writer.Write(strings)
		if err != nil {
			return err
		}
	}
	return nil
}
