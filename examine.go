package main

type Sentence struct {
	forA float64
	forB float64
}

type Examine struct {
	policy map[Action](map[Action]Sentence)
}

func NewExamine() Examine {
	policy := make(map[Action](map[Action]Sentence))
	policy[Silent] = make(map[Action]Sentence)
	policy[Silent][Silent] = Sentence{forA: -1, forB: -1}
	policy[Silent][Betray] = Sentence{forA: -3, forB: 0}

	policy[Betray] = make(map[Action]Sentence)
	policy[Betray][Silent] = Sentence{forA: 0, forB: -3}
	policy[Betray][Betray] = Sentence{forA: -2, forB: -2}

	return Examine{policy: policy}
}

func (e *Examine) CrossExamine(aAction Action, bAction Action) Sentence {
	return e.policy[aAction][bAction]
}

/*
func main() {
	pA := NewPrisoner("A")
	pB := NewPrisoner("B")

	pA.Print()
	pB.Print()

	e := NewExamine()

	sentence := e.CrossExamine(pA.action, pB.action)

	fmt.Println("PrisonerA sentenced to ", sentence.forA)
	fmt.Println("PrisonerB sentenced to ", sentence.forB)

	pA.StaySilent()
	pB.Betray()

	pA.Print()
	pB.Print()

	sentence = e.CrossExamine(pA.action, pB.action)

	fmt.Println("PrisonerA sentenced to ", sentence.forA)
	fmt.Println("PrisonerB sentenced to ", sentence.forB)
}
*/
