package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// URL da API que será chamada
	url := "https://jsonplaceholder.typicode.com/posts/1"

	// Cria uma requisição GET
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	// Cria um cliente HTTP e executa a requisição
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Lê o corpo da resposta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Imprime o corpo da resposta
	fmt.Println(string(body))
}
