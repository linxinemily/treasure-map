package domain

type Invincible struct {
	*AbstractState
}

func NewInvincible(role Role) *Invincible {
	return &Invincible{
		NewAbstractState(2, "無敵", role),
	}
}

func (state *Invincible) attacked() {
	//
}
