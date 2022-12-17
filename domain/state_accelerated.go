package domain

type Accelerated struct {
	*AbstractState
}

func NewAccelerated(role Role) *Accelerated {
	return &Accelerated{
		NewAbstractState(3, "加速", role),
	}
}

func (state *Accelerated) attacked() {
	state.role.applyState(NewNormal(state.role))
}

func (state *Accelerated) beforeTakeTurn() int {
	return 2
}
