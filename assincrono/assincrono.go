package assincrono

import (
	"fmt"
	"time"
)

func Display(str string) {
	defer fmt.Printf("chega de %s\n", str)
	for i := 0; i < 6; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(str, i)
	}
}
