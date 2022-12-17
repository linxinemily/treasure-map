package domain

type StateGenerator interface {
	generate(role Role) State
}

type StateInvincibleGenerator struct{}

func (g StateInvincibleGenerator) generate(role Role) State {
	return NewInvincible(role)
}

type StatePoisonedGenerator struct{}

func (g StatePoisonedGenerator) generate(role Role) State {
	return NewPoisoned(role)
}

type StateAcceleratedGenerator struct{}

func (g StateAcceleratedGenerator) generate(role Role) State {
	return NewAccelerated(role)
}

type StateHealingGenerator struct{}

func (g StateHealingGenerator) generate(role Role) State {
	return NewHealing(role)
}

type StateOrderlessGenerator struct{}

func (g StateOrderlessGenerator) generate(role Role) State {
	return NewOrderless(role)
}

type StateStockpileGenerator struct{}

func (g StateStockpileGenerator) generate(role Role) State {
	return NewStockpile(role)
}

type StateEruptingGenerator struct{}

func (g StateEruptingGenerator) generate(role Role) State {
	return NewErupting(role)
}

type StateTeleportGenerator struct{}

func (g StateTeleportGenerator) generate(role Role) State {
	return NewTeleport(role)
}
