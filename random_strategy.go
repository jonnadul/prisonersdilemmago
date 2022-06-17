package main

import (
	"math/rand"
	"time"
)

type Random struct {
	prisoner Prisoner
}

func NewRandom() (rd Random) {
	rand.Seed(time.Now().UnixNano())

	return Random{prisoner: NewPrisoner("Random")}
}

func (rd Random) Decide() (action Action, err error) {
	if rand.Intn(2) == 1 {
		rd.prisoner.action = Silent
	} else {
		rd.prisoner.action = Betray
	}

	rd.prisoner.Print()
	action = rd.prisoner.action
	return
}

func (rd Random) SetOpponentDecision(action Action) (err error) {
	return
}
