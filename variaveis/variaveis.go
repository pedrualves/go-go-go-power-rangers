package variaveis

import "log"

func Explicacao() {
	/* para numeros inteiros eh possivel explicitar o tamanho do inteiro a partir dos tipos:
	   int int8 int16 int32 int64 uint uint8 uint16 uint32 uint64 uintptr
	   qualquer tipo inteiro que comece com u significa que independente do sistema operacional
	   sera da quantidade definida */
	var num1 int = 42 // declarando uma variavel explicitando o tipo inteiro
	var num2 = 42     // declarando uma variavel implicitando do tipo inteiro
	num3 := 42        // declarando uma variavel por inferencia de tipo

	// para numeros flutuantes podemos especifiar entre 32bits e 64, respectivamente float32 e float64
	const f1 float32 = 3.1415 // declarando uma constante explicitando o tipo flutuante
	const f2 = 3.1415         // declarando uma constante implicitando o tipo flutuante
	f3 := 3.1415              // declarando uma variavel por inferencia de tipo

	// declarando blocos de variaveis explicitando o tipo string
	var (
		a string = " sou a letra a"
		b string = " sou a letra b"
	)

	// declarando uma constante com implicitando o tipo string
	const (
		c = " sou a letra c"
		d = " diiissss"
	)

	// tudo junto e misturado
	var (
		e    string = " sou a letra E"
		num4        = 42
	)

	ehVddEsseBilete := false // tao simples quanto isso para declarar booleanos
	eleNao := true

	log.Print(num1, num2, num3, num4, f1, f2, f3, a, b, c, d, e, ehVddEsseBilete, eleNao)
}
