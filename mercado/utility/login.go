package minhasFuncoes

import "fmt"

// =========== login ===============
func Login() bool {
	var login bool
	var tentativas int
	var tentativasMax int = 7

	for tentativas < tentativasMax {
		// entradas
		fmt.Println("digite seu usuario")
		var usuario string
		fmt.Scanln(&usuario)

		fmt.Println("digite seu senha")
		var senha string
		fmt.Scanln(&senha)

		// chama a finçao verificaçao login
		login = verifacaoLogin(usuario, senha)

		if !login { // se o login for falso repete o loop
			fmt.Println("usuario ou senha incoreto")
			tentativas++ // almenta o numero da variavel tentativas
			continue

		}

		if login {
			return true // se o login for true faz o login

		}

	}

	if tentativas == tentativasMax {
		// da um block e finaliza o sistema

		fmt.Println("voce atingiu  o limite de tentativas")

		return false
	}

	return false
}

//========= final do bloco login ========

// ========== funçao verificaçao de ususario e senha ============
func verifacaoLogin(usuario, senha string) bool {
	// mapa de usuario e senha
	usuarioCadastrado := map[string]string{
		"admin": "1234",
	}

	senhaCorreta, existe := usuarioCadastrado[usuario] //verifica se o usuario existe

	if !existe || senha != senhaCorreta {

		// verifica se a senha é igual a senha do usuario se
		// nao for passa false para a variavel login

		fmt.Println("usuario ou senha incorreta")
		return false

	}

	fmt.Printf("bemvindo %s", usuario) //faz o login

	return true
}

//========= final do bloco login ========
