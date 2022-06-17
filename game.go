package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func RunGame(pA Prisoner, pB Prisoner, numIter int) (err error) {

	fmt.Println("Starting Game!")

	pAScore := 0
	pBScore := 0

	for i := 0; i < numIter; i++ {
		var pAAction rune
		var pBAction rune

		fmt.Println("Iteration " + strconv.Itoa(i))

		for pAAction != 's' && pAAction != 'b' {
			fmt.Println("What should Prisoner " + pA.name + " do? [s/b]")
			reader := bufio.NewReader(os.Stdin)
			pAAction, _, err = reader.ReadRune()
			if err != nil {
				return err
			}
		}

		if pAAction == 's' {
			pA.action = Silent
		} else {
			pA.action = Betray
		}

		for pBAction != 's' && pBAction != 'b' {
			fmt.Println("What should Prisoner " + pB.name + " do? [s/b]")
			reader := bufio.NewReader(os.Stdin)
			pBAction, _, err = reader.ReadRune()
			if err != nil {
				return err
			}
		}

		if pBAction == 's' {
			pB.action = Silent
		} else {
			pB.action = Betray
		}

		pA.Print()
		pB.Print()

		examine := NewExamine()

		sentence := examine.CrossExamine(pA.action, pB.action)

		pAScore += sentence.forA
		pBScore += sentence.forB

		fmt.Println("Prisoner " + pA.name + " sentenced to " + strconv.Itoa(sentence.forA) + " current total " + strconv.Itoa(pAScore))
		fmt.Println("Prisoner " + pB.name + " sentenced to " + strconv.Itoa(sentence.forB) + " current total " + strconv.Itoa(pBScore))
	}

	return nil
}

func main() {
	pA := NewPrisoner("Alice")
	pB := NewPrisoner("Bob")

	RunGame(pA, pB, 10)
}
