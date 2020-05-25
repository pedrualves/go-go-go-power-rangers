package assincrono

import (
	"fmt"
	"time"
)

func Display(str string) {
	defer fmt.Printf("chega de %s", str)
	for w := 0; w < 6; w++ {
		time.Sleep(1 * time.Second)
		fmt.Println(str)
	}
}
