package csvp

import "strconv"

type numericCacheStatus byte

const (
	NCS_UNKNOWN numericCacheStatus = iota
	NCS_NOTNUMERIC
	NCS_FLOAT
)

type Cell struct {
	content string

	ncs          numericCacheStatus
	float64cache float64
}

func NewCell(content string) *Cell {
	return &Cell{content: content}
}

func (c Cell) cacheNumerics() {
	if c.ncs == NCS_UNKNOWN {
		s := NCS_NOTNUMERIC

		var err error
		c.float64cache, err = strconv.ParseFloat(c.content, 64)
		if err == nil {
			s = NCS_FLOAT
		}

		c.ncs = s
	}
}

func (c Cell) IsNumeric() bool {
	c.cacheNumerics()
	return c.ncs == NCS_FLOAT
}

func (c Cell) Float64() float64 {
	if c.IsNumeric() {
		return c.float64cache
	} else {
		return 0
	}
}
