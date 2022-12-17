package domain

import "math"

type Normal struct {
	*AbstractState
}

func NewNormal(role Role) *Normal {
	return &Normal{
		NewAbstractState(math.MaxInt, "正常", role),
	}
}
