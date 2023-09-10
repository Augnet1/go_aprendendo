package main

import (
	"net/http" // Importa o pacote net/http para criação de servidor http
)

func main() {
	http.ListenAndServe(":8080", nil) // Cria o servidor http na porta 8080
}
