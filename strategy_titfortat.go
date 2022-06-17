package main

type TitForTat struct {
	prisoner           Prisoner
	prevOpponentAction Action
}

func NewTitForTat() (tft TitForTat) {
	return TitForTat{prisoner: NewPrisoner("Tit For Tat"), prevOpponentAction: : Silent}
}

func (tft TitForTat) Decide() (action Action, err error) {
	tft.prisoner.action = tft.prevtft.prevOpponentAction

	tft.prisoner.Print()
	action = tft.prisoner.action
	return
}

func (tft TitForTat) SetOpponentDecision(action Action) (err error) {\
	tft.prevOpponentAction = action
	return
}
