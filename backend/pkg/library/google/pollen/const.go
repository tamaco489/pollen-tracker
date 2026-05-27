package pollen

// pollenLevel は Google Pollen API の花粉指数の値域を表す型
type pollenLevel int

const (
	minPollenLevel pollenLevel = 0
	maxPollenLevel pollenLevel = 5
)

func (l pollenLevel) isValid() bool {
	return l >= minPollenLevel && l <= maxPollenLevel
}

func (l pollenLevel) toInt() int {
	return int(l)
}
