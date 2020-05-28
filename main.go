package main

import (
	"fmt"
	"go-go-go-power-rangers/assincrono"
	"go-go-go-power-rangers/conversor"
	"go-go-go-power-rangers/estruturas"
	"go-go-go-power-rangers/funcoes"
	animal "go-go-go-power-rangers/objetos/composicao"
	impl "go-go-go-power-rangers/objetos/implementacoes"
	inter "go-go-go-power-rangers/objetos/interface"
	"go-go-go-power-rangers/operadores"
	"go-go-go-power-rangers/variaveis"
	"log"
)

func andarDeTrem(c inter.Member) {
	c.ComprarNoShopTrem()
}

func main() {
	log.Print("pao de batata")

	variaveis.Explicacao()

	operadores.Atribuicao()
	operadores.Aritmeticos()
	operadores.Atribuicao()
	operadores.Logicos()

	funcoes.FuncaoPublica()

	estruturas.ExemplosDeEstruturasDeRepeticao()
	estruturas.ExemplosDeEstruturasDeCondicional()

	client := impl.Client{
		Name: "Napoleao Bonaparte",
	}

	employee := impl.Employee{
		Alias: "Diiiiiissssss",
	}

	andarDeTrem(client)
	andarDeTrem(employee)

	user := conversor.User{
		Name:   "Abelardo",
		Online: true,
	}

	conversor.ConverterUserToJSON(user)

	conversor.ConverterUserFromJSON(`{"username":"userboladao","active":true}`)

	// chamando de maneira ASSINCRONA
	go assincrono.Display("Caneta Azul")

	// chamando de maneira SINCRONA
	assincrono.Display("Azul Caneta")

	done := make(chan bool)
	msg := make(chan string)

	go sendMessage(msg)
	go receiveMessage(msg, done)

	<-done

	// exemplo pratico de heranca por composicao
	truques := [...]string{"deitar", "rolar", "pular", "sentar"}

	iena := animal.Canino{}
	iena.Nome = "Saori"
	iena.Pet = true
	iena.Truques = truques

	doguinho := animal.Cachorro{}
	doguinho.Nome = "Pipoca"
	doguinho.Pet = true
	doguinho.Truques = truques

	iena.Latir()
	iena.Alimentar()

	doguinho.Latir()
	doguinho.FingirDeMorto()
	doguinho.Alimentar()

	fmt.Println(iena)
	fmt.Println(doguinho)

	lobo := animal.Lobo{}
	lobo.Nome = "Petrus"
	lobo.Pet = false
	lobo.Matilha = 4
	lobo.Mortes = 99
	lobo.PresaFavorita = "chapeuzinho vermelho"
	lobo.Alimentar()

}

func sendMessage(msg chan string) {
	fmt.Println("sending message")
	msg <- "DIIISSSSS"
}

func receiveMessage(msg chan string, done chan bool) {
	fmt.Println("receiveing message ", <-msg)
	done <- false
}
