package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func main() {
	clienteWerm := clientes.Titular{Nome: "Wermenson", CPF: "000.000.000-00", Profissao: "Desenvolvedor"}
	contaDoWerm := contas.ContaCorrente{Titular: clienteWerm, NumeroAgencia: 123, NumeroConta: 123456}
	contaDoWerm.Depositar(1000)

	fmt.Println(contaDoWerm.ObterSaldo())

	PagarBoleto(&contaDoWerm, 100)

	fmt.Println(contaDoWerm.ObterSaldo())
}

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}
