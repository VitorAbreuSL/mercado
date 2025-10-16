package main

import (
	"fmt"
	m "projeto/utility"
)

func main() {

	var opcao int

	// login

	if m.Login() == false {
		fmt.Println("voce atingiu o limite de tentativas")
		return
	}
	// menu

	for {
		// menu
		m.Menu(opcao)

		if m.Menu(opcao) == false {
			break
		}

	}
}
