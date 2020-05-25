package main

import (
	"go-go-go-power-rangers/conversor"
	"go-go-go-power-rangers/estruturas"
	"go-go-go-power-rangers/funcoes"
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

	conversor.ConverterUserFromJSON(`{"username":"abelardo","active":true}`)
}
