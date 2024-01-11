package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoWermenson := ContaCorrente{titular: "Wermenson",
		numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}

	contaDaLizz := ContaCorrente{"Lizz", 222, 111222, 1000}

	var contaDaLuna *ContaCorrente
	contaDaLuna = new(ContaCorrente)
	contaDaLuna.titular = "Luna"

	fmt.Println(contaDoWermenson)
	fmt.Println(contaDaLizz)
	fmt.Println(*contaDaLuna)
}
