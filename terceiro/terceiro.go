package main

import "fmt"

func main() {
	fmt.Println("Olá! Qual o seu nome?")
	var name string   //Declaração da variável name como string
	fmt.Scanln(&name) //Recebe a variável name como string
	fmt.Printf("Olá %s. Esse foi meu terceiro programa em Go!\n", name)
}
