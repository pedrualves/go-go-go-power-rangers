package orientado_objetos

import "log"

type Client struct {
	ID     int
	Name   string
	Active bool
}

func (c Client) ComprarNoShopTrem() bool {
	log.Println(c.Name + " vai chamar com 3 reais e vai levar 6 barras de suflair")
	return false
}
