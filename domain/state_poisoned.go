package domain

type Poisoned struct {
	*AbstractState
}

func NewPoisoned(role Role) *Poisoned {
	return &Poisoned{
		NewAbstractState(3, "中毒", role),
	}
}

func (state *Poisoned) beforeTakeTurn() int {
	state.role.subtractHP(15)
	state.role.checkRoleHP(state.role)
	return 1
}
