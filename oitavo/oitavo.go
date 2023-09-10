package main

import (
	"fmt"
	"time"
)

func counter() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

// GO routines
// Thread T0
func main() {
	// counter() // Imprime os valores em 10 segundos
	// counter() // Imprime os valores em +10 segundos - Total 20 segundos
	// counter() // Imprime os valores em +10 segundos - Total 30 segundos
	go counter() // Thread T1 - Cria uma Thread e imprime os valores em 10 segundos
	go counter() // Thread T2 - Cria uma Thread e imprime os valores em 10 segundos
	counter()    // Thread T0 - Imprime os valores em 10 segundos - Total 10 segundos
}
