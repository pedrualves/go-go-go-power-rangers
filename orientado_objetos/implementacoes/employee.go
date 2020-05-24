package orientado_objetos

import "log"

type Employee struct {
	Job   string
	Alias string
	Alive bool
}

func (c Employee) ComprarNoShopTrem() bool {
	log.Println(c.Alias + " vai chamar com 3 reais e vai levar 6 barras de suflair")
	return true
}
