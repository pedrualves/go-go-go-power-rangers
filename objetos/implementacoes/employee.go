package objetos

import "log"

type Employee struct {
	Job   string
	Alias string
	Alive bool
}

func (c Employee) ComprarNoShopTrem() bool {
	log.Printf("e na promocao %s vai chamar com 5 reais e vai levar 2 barras de suflair\n", c.Alias)
	return true
}
