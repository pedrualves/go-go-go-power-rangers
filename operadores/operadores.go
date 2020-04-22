package operadores

import "log"

const (
	a       int  = 3
	b       int  = 2
	valueI  bool = true
	valueII bool = false
)

var num int = 42

func Atribuicao() {
	num += 42
	log.Println("(atribuindo e somando) num += 42 = ", num)

	num -= 42
	log.Println("(atribuindo e subtraindo) num -= 42 = ", num)

	num /= 42
	log.Println("(atribuindo e dividindo) num /= 42) = ", num)

	num *= 42
	log.Println("(atribuindo e multiplicacao) num *= 42) = ", num)

	num %= 42
	log.Println("(atribuindo e mod) num %= 42) = ", num)
}

func Aritmeticos() {
	log.Println("soma de 3 e 2", a+b)
	log.Println("subtracao de 3 e 2", a-b)
	log.Println("divisao de 3 e 2", a/b)
	log.Println("multiplicacao de 3 e 2", a*b)
	log.Println("mod de 3 e 2", a%b)
}

func Logicos() {
	log.Println("olha o AND resultI && resultII =>", valueI && valueII)
	log.Println("olha o OR resultI || resultII =>", valueI || valueII)
	log.Println("olha o NOT !(resultI && resultII) =>", !(valueI && valueII))
}

func Relacionais() {
	log.Println("comparando iguais", a == b)
	log.Println("comparando diferentes", a != b)
	log.Println("comparando maior", a > b)
	log.Println("comparando menor", a < b)
	log.Println("comparando maior igual", a >= b)
	log.Println("comparando menor igual", a <= b)
}
