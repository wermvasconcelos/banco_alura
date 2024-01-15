package main

import (
	"banco/contas"
	"fmt"
)

func main() {
	contaDaLizz := contas.ContaCorrente{Titular: "Lizz", Saldo: 300}
	contaDoWerm := contas.ContaCorrente{Titular: "Werm", Saldo: 300}

	fmt.Println(contaDoWerm.Transferir(200, &contaDaLizz))
	fmt.Println(contaDaLizz, contaDoWerm)

}
