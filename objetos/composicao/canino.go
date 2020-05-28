package composicao

import "fmt"

type Canino struct {
	Animal
	Truques [4]string
}

func (c Canino) Latir() {
	fmt.Println("au au")
}
