package domain

type Erupting struct {
	*AbstractState
}

func NewErupting(role Role) *Erupting {
	return &Erupting{
		NewAbstractState(3, "爆發", role),
	}
}

func (state *Erupting) checkStateIfExpired() bool {
	if state.round >= state.expiredRound {
		state.role.applyState(NewTeleport(state.role))
		return true
	}
	return false
}

func (state *Erupting) attack() {
	for _, enemy := range state.role.getAllEnemies() {
		enemy.attacked()
	}
}
