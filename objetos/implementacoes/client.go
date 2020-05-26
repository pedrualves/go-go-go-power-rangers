package objetos

import "log"

type Client struct {
	ID     int
	Name   string
	Active bool
}

func (c Client) ComprarNoShopTrem() bool {
	log.Printf("o patrão %s vai chamar com 10 reais e vai levar 4 barras de suflair\n", c.Name)
	return false
}
