package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Addr struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Por favor informe UM cep como parametro!")
		return
	}

	cep := os.Args[1]

	req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v", err)
	}
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v", err)
	}

	var addr Addr
	err = json.Unmarshal(res, &addr)
	if err != err {
		fmt.Fprintf(os.Stderr, "Erro ao fazer parse do json: %v", err)
	}

	file, err := os.Create("address.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao criar arquivo: %v", err)
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "UF: %s, Localidade: %s, Bairro: %s, Logradouro: %s\n", addr.Uf, addr.Localidade, addr.Bairro, addr.Logradouro)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao escrever endereço no arquivo: %v", err)
	}

	fmt.Printf("UF: %s, Localidade: %s, Bairro: %s, Logradouro: %s\n", addr.Uf, addr.Localidade, addr.Bairro, addr.Logradouro)
}
