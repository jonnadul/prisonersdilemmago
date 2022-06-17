package main

type Strategy interface {
	Decide() (Action, error)
	SetOpponentDecision(Action) error
}
