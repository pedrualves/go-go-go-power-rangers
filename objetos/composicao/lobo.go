package composicao

import "fmt"

type Lobo struct {
	Animal
	Predador
}

func (l Lobo) Uivar() {
	fmt.Println("AUUUUUUUUUUUUUUUUUUU")
}
