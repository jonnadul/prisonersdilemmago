package main

import (
	"fmt"
)

type Sentence struct {
	forA int
	forB int
}

type Game struct {
	policy map[Action](map[Action]Sentence)
}

func NewGame() Game {
	policy := make(map[Action](map[Action]Sentence))
	policy[Silent] = make(map[Action]Sentence)
	policy[Silent][Silent] = Sentence{forA:-1, forB:-1}
	policy[Silent][Betray] = Sentence{forA:-3, forB:0}

	policy[Betray] = make(map[Action]Sentence)
	policy[Betray][Silent] = Sentence{forA:0, forB:-3}
	policy[Betray][Betray] = Sentence{forA:-2, forB:-2}

	return Game{policy:policy}
}

func (g *Game) CrossExamine(pA Prisoner, pB Prisoner) Sentence {
	return g.policy[pA.action][pB.action]
}

func main() {
	pA := NewPrisoner("A")
	pB := NewPrisoner("B")

	pA.Print()
	pB.Print()

	g := NewGame()

	sentence := g.CrossExamine(pA, pB)

	fmt.Println("PrisonerA sentenced to ", sentence.forA)
	fmt.Println("PrisonerB sentenced to ", sentence.forB)

	pA.StaySilent()
	pB.Betray()

	pA.Print()
	pB.Print()

	sentence = g.CrossExamine(pA, pB)

	fmt.Println("PrisonerA sentenced to ", sentence.forA)
	fmt.Println("PrisonerB sentenced to ", sentence.forB)
}
