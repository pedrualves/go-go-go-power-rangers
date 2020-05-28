package composicao

import "fmt"

type Cachorro Canino

func (c Cachorro) Latir() {
	fmt.Println("au au bem mansinho")
}

func (c Cachorro) FingirDeMorto() {
	fmt.Println("e morreu, de brincadeira")
}
