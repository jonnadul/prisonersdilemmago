package main

type Always struct {
	prisoner Prisoner
	silent   bool
}

func NewAlways(silent bool) (al Always) {
	if silent {
		al = Always{prisoner: NewPrisoner("Always Silent"), silent: silent}
	} else {
		al = Always{prisoner: NewPrisoner("Always Betray"), silent: silent}
	}
	return
}

func (al Always) Decide() (action Action, err error) {

	if al.silent {
		al.prisoner.action = Silent
	} else {
		al.prisoner.action = Betray
	}

	al.prisoner.Print()
	action = al.prisoner.action
	return
}

func (al Always) SetOpponentDecision(action Action) (err error) {
	return
}
