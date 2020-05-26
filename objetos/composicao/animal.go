package composicao

import "fmt"

type Animal struct {
	Nome string
	Pet  bool
}

func (a Animal) alimentar() {
	fmt.Println("yami yami")
}

type Canino struct {
	Animal
	Truques string
}

func (c Canino) latir() {
	fmt.Println("au au")
}

type Felino struct {
	Animal
	AutoLimpante   bool
	ApelidoFofinho string
}

func (f Felino) miar() {
	fmt.Println("miau miau")
}

type Cachorro Canino

type Lobo struct {
	Animal
	Matilha int
}

type Gato Felino

type Tigre struct {
	Animal
	Matar bool
}
