package minhasFuncoes

import "fmt"

func Menu(opcao int) bool {
	// menu de seleçao
	for {

		fmt.Println("menu inicial")
		fmt.Println("cadastro (1)/ vendas(2) / encerrar(3))")
		fmt.Scanln(&opcao) // variavel de seleçao do menu

		switch opcao {
		case 1:
			CadastroProduto()
		case 2:
			Vendas()
		case 3:
			fmt.Println("encerrando o sisterma")
			return false
		default:
			fmt.Println("opção invalida")

		}
	}
}
