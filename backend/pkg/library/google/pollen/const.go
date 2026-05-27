package pollen

// PollenLevel は Google Pollen API の花粉指数の値域を表す型
type PollenLevel int

const (
	MinPollenLevel PollenLevel = 0
	MaxPollenLevel PollenLevel = 5
)

func (l PollenLevel) IsValid() bool {
	return l >= MinPollenLevel && l <= MaxPollenLevel
}

func (l PollenLevel) ToInt() int {
	return int(l)
}
