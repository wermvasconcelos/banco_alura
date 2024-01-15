package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDaLizz := ContaCorrente{titular: "Lizz", saldo: 300}
	contaDoWerm := ContaCorrente{titular: "Werm", saldo: 300}

	fmt.Println(contaDoWerm.Transferir(200, &contaDaLizz))
	fmt.Println(contaDaLizz, contaDoWerm)

}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Deposito realizado com sucesso!", c.saldo
	}

	return "Valor do deposito menor que zero", c.saldo

}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	}

	return "Saldo insuficiente"
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < c.saldo && valorDaTransferencia > 0 {
		c.saldo -= valorDaTransferencia
		contaDestino.saldo += valorDaTransferencia
		return true
	}
	return false
}
