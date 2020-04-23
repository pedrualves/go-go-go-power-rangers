package main

import (
	"go-go-go-power-rangers/estruturas"
	"go-go-go-power-rangers/funcoes"
	"go-go-go-power-rangers/operadores"
	"go-go-go-power-rangers/variaveis"
	"log"
)

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
}
