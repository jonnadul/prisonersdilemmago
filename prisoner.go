package main

import (
	"fmt"
)

type Action int64

const (
	Silent Action = iota
	Betray
)

type Prisoner struct {
	name string
	action Action
}

func NewPrisoner(name string) Prisoner {
	return Prisoner{name:name, action:Silent}
}

func (p *Prisoner) StaySilent() {
	p.action = Silent
}

func (p *Prisoner) Betray() {
	p.action = Betray
}

func (p *Prisoner) Print() {
	fmt.Print("Prisoner ", p.name)

	if p.action == Silent {
		fmt.Println(" stays silent.")
	} else {
		fmt.Println(" betrays.")
	}
}
/*
func main() {
	prisonerA := NewPrisoner("A")
	prisonerA.Print()
	prisonerA.Betray()
	prisonerA.Print()
	prisonerA.StaySilent()
	prisonerA.Print()
}
*/
