package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func RunGame(pA Strategy, pB Strategy, numIter int) (err error) {

	fmt.Println("Starting Game!")

	pAScore := 0
	pBScore := 0

	for i := 0; i < numIter; i++ {
		var pAAction Action
		var pBAction Action

		fmt.Println("Iteration " + strconv.Itoa(i))

		pAAction, err = pA.Decide()
		if err != nil {
			return
		}

		pBAction, err = pB.Decide()
		if err != nil {
			return
		}

		examine := NewExamine()

		sentence := examine.CrossExamine(pAAction, pBAction)

		pAScore += sentence.forA
		pBScore += sentence.forB

		fmt.Println("Prisoner A sentenced to " + strconv.Itoa(sentence.forA) + " current total " + strconv.Itoa(pAScore))
		fmt.Println("Prisoner B sentenced to " + strconv.Itoa(sentence.forB) + " current total " + strconv.Itoa(pBScore))

		pA.SetOpponentDecision(pBAction)
		pB.SetOpponentDecision(pAAction)
	}

	return nil
}

func main() {
	upA := UserPrompt{prisoner: NewPrisoner("Alice"), reader: bufio.NewReader(os.Stdin)}
	upB := UserPrompt{prisoner: NewPrisoner("Bob"), reader: bufio.NewReader(os.Stdin)}

	RunGame(upA, upB, 10)
}
