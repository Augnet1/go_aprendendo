package main

import (
	"fmt"
	"net/http" // Importa o pacote net/http para criação de servidor http
)

type Curso struct { // Cria uma struct
	Nome      string
	Descricao string
	Preco     int
}

func (c Curso) GetFullInfo() string { // Cria um método para retornar todos os dados da struct
	return fmt.Sprintf("Nome: %s, Descrição: %s, Preço: %d", c.Nome, c.Descricao, c.Preco)
}

func main() {
	curso := Curso{ // Cria um curso do tipo struct com os dados seguintes
		Nome:      "Golang",
		Descricao: "Curso Golang",
		Preco:     100,
	}

	fmt.Println(curso.GetFullInfo())  //Chama a função para retornar todos os dados de curso
	http.HandleFunc("/", home)        // Quando for acessado o localhost:8080/ chamará a função home
	http.ListenAndServe(":8080", nil) // Cria o servidor http na porta 8080
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World\n"))
}
