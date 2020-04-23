package estruturas

import (
	"fmt"
	"log"
)

// essa funcao publica eh responsavel por invocar as funcoes privadas
func ExemplosDeEstruturasDeRepeticao() {
	const num int = 10

	forTradicional(num)

	whileCondicional(num)

	forEach("diiissssss")
}

// exemplo do classico for
func forTradicional(num int) {
	for i := 0; i < num; i++ {
		fmt.Println(i, " carneirinho(s)")
	}
}

// curiosamente nao temos a palavra reservada while,
// mas a ideia eh a mesma para esse tipo de iteracao e sem a condicional seria um while infinito
func whileCondicional(num int) {
	stop := 0
	for stop > num {
		log.Println("nao para, nao para")
	}
}

// exemplo do classico de um elegante for each, interessante que podemos manipular o indice e valor
func forEach(palavra string) {
	for index, value := range palavra {
		// aqui tem um pulo do gato pra exibir o caracter
		// caso contrario sera exibido o numero do caracter
		log.Printf("%d %c", index, value)
	}
}
