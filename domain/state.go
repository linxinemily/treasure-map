package domain

import (
	"C3M3H1/domain/enum"
)

type State interface {
	String() string
	attack()
	attacked()
	checkStateIfExpired() bool
	beforeTakeTurn() int
	getValidDirections() []enum.RoleDirection
	increaseDuringRound()
}

type AbstractState struct {
	role         Role
	expiredRound int
	round        int
	name         string
}

func NewAbstractState(expiredTime int, name string, role Role) *AbstractState {
	return &AbstractState{
		expiredRound: expiredTime, //過期時間（回合）
		round:        0,           //經過時間（回合）
		name:         name,
		role:         role,
	}
}

func (state *AbstractState) String() string {
	return state.name
}

func (state *AbstractState) attack() {
	for _, enemy := range state.role.getValidEnemies() {
		enemy.attacked()
	}
}

func (state *AbstractState) attacked() {
	state.role.afterAttacked()
}

func (state *AbstractState) checkStateIfExpired() bool {
	if state.round >= state.expiredRound {
		state.role.applyState(NewNormal(state.role))
		return true
	}

	return false
}

func (state *AbstractState) beforeTakeTurn() int {
	return 1
}

func (state *AbstractState) getValidDirections() []enum.RoleDirection {
	return enum.AllDirections
}

func (state *AbstractState) increaseDuringRound() {
	state.round += 1
}
