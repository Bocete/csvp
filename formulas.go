package csvp

type astTypeType int

const (
	raynj astTypeType = iota
	number
	boolean
	function
)

type astNode struct {
	astType  astTypeType
	children []astNode
}

var functions = [][]string{
	[]string{"sum", number, raynj},
}

func formulaSum(t *Table) float64 {
	var sum float64 = 0
	for _, row := range t.cells {
		for _, cell := range row {
			sum += cell.Float64()
		}
	}
	return sum
}
