package main

import (
	"fmt"
	inter "go-go-go-power-rangers/objetos/interface"
)

func andarDeTrem(c inter.Member) {
	c.ComprarNoShopTrem()
}

func main() {
	// log.Print("pao de batata")

	// variaveis.Explicacao()

	// operadores.Atribuicao()
	// operadores.Aritmeticos()
	// operadores.Atribuicao()
	// operadores.Logicos()

	// funcoes.FuncaoPublica()

	// estruturas.ExemplosDeEstruturasDeRepeticao()
	// estruturas.ExemplosDeEstruturasDeCondicional()

	// client := impl.Client{
	// 	Name: "Napoleao Bonaparte",
	// }

	// employee := impl.Employee{
	// 	Alias: "Diiiiiissssss",
	// }

	// andarDeTrem(client)
	// andarDeTrem(employee)

	// user := conversor.User{
	// 	Name:   "Abelardo",
	// 	Online: true,
	// }

	// conversor.ConverterUserToJSON(user)

	// conversor.ConverterUserFromJSON(`{"username":"userboladao","active":true}`)

	// chamando de maneira ASSINCRONA
	// go assincrono.Display("Caneta Azul")

	// chamando de maneira SINCRONA
	// assincrono.Display("Azul Caneta")

	done := make(chan bool)
	msg := make(chan string)

	go sendMessage(msg)
	go receiveMessage(msg, done)

	<-done
}

func sendMessage(msg chan string) {
	fmt.Println("sending message")
	msg <- "DIIISSSSS"
}

func receiveMessage(msg chan string, done chan bool) {
	fmt.Println("receiveing message ", <-msg)
	done <- false
}
