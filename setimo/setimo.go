package main

import (
	"encoding/json"
	"fmt"
	"net/http" // Importa o pacote net/http para criação de servidor http
)

type Curso struct { // Cria uma struct
	Nome      string `json:"curso"`     //TAG para o jason retornar curso ao invés de Nome
	Descricao string `json:"descrição"` //TAG para o jason retornar descrição ao invés de Descricao
	Preco     int    `json:"preço"`     //TAG para o jason retornar preço ao invés de Preco
}

func (c Curso) GetFullInfo() string { // Cria um método para retornar todos os dados da struct
	return fmt.Sprintf("Nome: %s, Descrição: %s, Preço: %d", c.Nome, c.Descricao, c.Preco)
}

func main() {
	http.HandleFunc("/", home)        // Quando for acessado o localhost:8080/ chamará a função home
	http.ListenAndServe(":8080", nil) // Cria o servidor http na porta 8080
}

func home(w http.ResponseWriter, r *http.Request) {
	curso := Curso{ // Cria um curso do tipo struct com os dados seguintes
		Nome:      "Golang",
		Descricao: "Curso Golang",
		Preco:     100,
	}
	json.NewEncoder(w).Encode(curso)
}
