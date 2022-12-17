package domain

type Stockpile struct {
	*AbstractState
}

func NewStockpile(role Role) *Stockpile {
	return &Stockpile{
		NewAbstractState(2, "蓄力", role),
	}
}

func (state *Stockpile) checkStateIfExpired() bool {
	if state.round >= state.expiredRound {
		state.role.applyState(NewErupting(state.role))
		return true
	}
	return false
}

func (state *Stockpile) attacked() {
	state.role.applyState(NewNormal(state.role))
}
