package main

import (
	"net/http" // Importa o pacote net/http para criação de servidor http
)

func main() {
	http.HandleFunc("/", home)        // Quando for acessado o localhost:8080/ chamará a função home
	http.ListenAndServe(":8080", nil) // Cria o servidor http na porta 8080
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World\n")) // A função home retorna um array de bytes "Hello World\n"
}
