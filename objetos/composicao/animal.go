package composicao

import "fmt"

// no pacote composicao temos exemplos de heranca por composicao

type Animal struct {
	Nome string
	Pet  bool
}

func (a Animal) Alimentar() {
	fmt.Println("yami yami")
}
