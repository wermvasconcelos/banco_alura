package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDaGeovanna := ContaCorrente{}
	contaDaGeovanna.titular = "Geovanna"
	contaDaGeovanna.saldo = 500

	fmt.Println(contaDaGeovanna.saldo)
	fmt.Println(contaDaGeovanna.Sacar(200))
	fmt.Println(contaDaGeovanna.saldo)

}

func (c *ContaCorrente) Depositar(valorDeposito float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
	}

}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	}

	return "Saldo insuficiente"
}
