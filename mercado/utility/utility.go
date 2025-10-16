package minhasFuncoes

import "fmt"

// =======variavel global mapa========
var produtos = make(map[string]float64)

//====================================

//=========  cadastramento de produtos ======

func CadastroProduto() {

	for {
		//entrada
		fmt.Println("digite o nome do produto, para encerrar digite end")
		var produtoNome string
		fmt.Scanln(&produtoNome)
		//finalisa o cadastro
		if produtoNome == "end" {
			fmt.Println("cadastro completo")
			break

		}
		//entrada valor
		fmt.Println("digite o valor do produto")
		var produtoValor float64
		fmt.Scanln(&produtoValor)

		// verifica se ja existe o produto
		_, existe := produtos[produtoNome]

		//se existe avisa e volta para o loop
		if existe {
			fmt.Println("produto ja cadastrado")
			continue
		}

		// cadastra o produto e o valor
		produtos[produtoNome] = produtoValor
	}

}

// ======== fim do bloco de cadastramento de produtos =========

//================ vendas====================

func Vendas() {

	var soma, Valorpagar float64
	var compraConcluida bool

	for {
		// entrada gralmente e o id do produto ou nome
		fmt.Println("digite um valor, para finalisar a compra digite (concluir) para cancelar a compra digite (cancelar)")
		var produtoNome string
		fmt.Scanln(&produtoNome)

		//segue para o pagamento
		if produtoNome == "concluir" {
			fmt.Println("segindo para pagameto")
			Valorpagar = soma // recebe o valor total a pagar
			compraConcluida = true
			break
		}
		// cancela a compra
		if produtoNome == "cancela" {
			fmt.Println("compra cancelada")
			compraConcluida = false
			break

		}
		// verifica se o produto existe
		valor, existe := produtos[produtoNome]

		if !existe {
			fmt.Println("nao existe")
			continue
		}

		soma += valor // soma e armazena o valor

	}
	// chama a funçao pagamento
	if compraConcluida {
		Pagamento(Valorpagar)

	} else {
		return
	}

}

//================ pagamento ====================

func Pagamento(Valorpagar float64) {

	var compra_concluida bool // confirma a compra

	for {
		// entrada
		fmt.Printf("valor a total e : %f , qual a forma de pagameto (debito/ credito/ pix/ dinheiro)", Valorpagar)
		var tipo_pagamento string
		fmt.Scanln(&tipo_pagamento)

		// as opçoes
		switch tipo_pagamento {

		case "debito": // no debito so imprime
			fmt.Printf("valor a pagar é : %f", Valorpagar)
			compra_concluida = true

		case "credito": //chama afunçao calculo de juros
			Valorpagar = calculojuros(Valorpagar)

		case "pix": // chama um pequeno desconto
			Valorpagar = Valorpagar - (Valorpagar * 0.05)
			fmt.Printf("valor a pagar é : %f", Valorpagar)
			compra_concluida = true

		case "dinheiro": // no dinheiro so printa a mensagem
			fmt.Printf("valor a pagar é : %f", Valorpagar)
			compra_concluida = true

		default: // restart do loop
			fmt.Printf("essa opçao nao existe")
			compra_concluida = false
		}

		// abre um dialgo com o funcionario
		if compra_concluida {
			// entrada
			for {
				fmt.Println("se deseja encerrar o caixa digite (encerrar) para próxima compra digite (seguir):")
				var trabalho string
				fmt.Scanln(&trabalho)
				// opçoes

				switch trabalho {
				case "encerrar": // retorna ao menu principal
					return
				case "seguir": //continua o trabalho
					Vendas()
					return
				default:
					fmt.Println("opção não existe")
				}
			}

			// volta para o menu
		}
	}

}

//===== fim do bloco pagameto =======

//================= juros======================
func calculojuros(Valorpagar float64) float64 {
	//juros simples
	juros := 0.09

	// entrada para saber o numero dr pacelas
	fmt.Println("quantidade de pacelas")
	var parcelas int
	fmt.Scanln(&parcelas)
	//calculo de juros
	Valorpagar = Valorpagar * (1 + (juros/100)*float64(parcelas))

	valorpacela := Valorpagar / float64(parcelas)

	fmt.Printf("valor a pagar é : %f, o valor de cada pacela e igual %f x%d", Valorpagar, valorpacela, parcelas)
	return Valorpagar
}

//============= final do bloco juros =======
