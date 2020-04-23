package estruturas

import (
	"log"
	"time"
)

// essa funcao publica eh responsavel por invocar as funcoes privadas
func ExemplosDeEstruturasDeCondicional() {
	log.Println(exemploIfElse("tava_bom"),
		exemploIfElse("disseram_que_ia_melhorar"),
		exemploIfElse("mas_parece_que_piorou"))

	exemploSwitchCase()
}

// exemplo de if/else complexo
func exemploIfElse(status string) string {
	if status == "tava_bom" {
		status = "disseram_que_ia_melhorar"
	} else if status == "disseram_que_ia_melhorar" {
		status = "mas_parece_que_piorou"
	} else {
		status = "nao_sei_de_mais_nada"
	}

	return status
}

// exemplo de switch case
func exemploSwitchCase() {
	switch time.Now().Weekday() { //dentro switch utilizei a lib time nativa do golang pra recuperar o dia da semana
	case time.Saturday: // case simples
		log.Println("melhor impossivel!")
		break // o break interrompera e nenhum outro caso sera avaliado
	case time.Sunday: // case simples
		log.Println("bom também!")
	case time.Tuesday, time.Wednesday, time.Thursday: //case list, onde se uma opcao for true esse case sera executado
		log.Println("dia de trabalhar!")
	case time.Monday: // mais um case simples
		log.Println("ninguem quer saber de segunda, pula pra sexta")
		fallthrough // com fallthrough o caso sera executado e em seguida obrigatoriamente passara para o proximo caso
	default: // caso padrao, se nenhum dos anteriores for satisfatorio a condicao
		log.Println("aff mais um segunda feira")
	}
}
