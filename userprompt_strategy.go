package main

import (
	"bufio"
	"fmt"
	"os"
)

type UserPrompt struct {
	prisoner Prisoner
	reader   *bufio.Reader
}

func NewUserPrompt() UserPrompt {
	return UserPrompt{prisoner: NewPrisoner("User Prompt"), reader: bufio.NewReader(os.Stdin)}
}

func (up UserPrompt) Decide() (action Action, err error) {
	var prompt_action rune

	for prompt_action != 's' && prompt_action != 'b' {
		fmt.Println("What should Prisoner " + up.prisoner.name + " do? [s/b]")
		prompt_action, _, err = up.reader.ReadRune()
		if err != nil {
			return
		}
	}

	if prompt_action == 's' {
		up.prisoner.action = Silent
	} else {
		up.prisoner.action = Betray
	}

	up.prisoner.Print()

	action = up.prisoner.action

	return
}

func (up UserPrompt) SetOpponentDecision(action Action) (err error) {
	return
}
