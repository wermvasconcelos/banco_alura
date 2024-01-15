package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func main() {
	clienteWerm := clientes.Titular{"Wermenson", "000.000.000-00", "Desenvolvedor"}
	contaDoWerm := contas.ContaCorrente{clienteWerm, 123, 123456, 100}
	fmt.Println(contaDoWerm)

}
