package domain

type Teleport struct {
	*AbstractState
}

func NewTeleport(role Role) *Teleport {
	return &Teleport{
		NewAbstractState(1, "瞬身", role),
	}
}

func (state *Teleport) checkStateIfExpired() bool {
	if state.round >= state.expiredRound {
		state.role.getMap().moveRoleToRandomPosition(state.role)
		state.role.applyState(NewNormal(state.role))
		return true
	}
	return false
}
