package main

import (
	"encoding/json"
	"fmt"
	"io"
	"multithreading-challenge/internal/dto"
	"net/http"
	"os"
	"time"
)

func main() {
	//Cria canal para receber resultado da API mais rapida
	resultChan := make(chan dto.ViaCEP)
	resultChan2 := make(chan dto.BrasilApi)

	//	CEP a ser buscado
	cep := "09210210"

	//	Chamar API brasilApi
	go chamaViaCep(cep, resultChan)

	//	Chamar API viaCEP
	go chamaBrasilApi(cep, resultChan2)

	//	Controla o fluxo de respostas via canal
	select {
	case response := <-resultChan:
		fmt.Println("API mais rapida (ViaCEP) \nResposta:\n ", response)
		break
	case response := <-resultChan2:
		fmt.Println("Api mais rapida (BrasilAPI) \nResposta:\n", response)
		break
	case <-time.After(1 * time.Second):
		fmt.Println("Erro de timeout: Não foi possível obter a resposta dentro do tempo limite.")
	}
}

func chamaViaCep(cep string, resultChan chan<- dto.ViaCEP) {
	req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		fmt.Fprint(os.Stderr, "Erro ao fazer requisicao: %v\n", err)
		return
	}
	res, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
		return
	}
	var data dto.ViaCEP
	err = json.Unmarshal(res, &data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer o parse da resposta %v\n", err)
		return
	}
	resultChan <- data
}

func chamaBrasilApi(cep string, resultChan chan<- dto.BrasilApi) {
	req, err := http.Get("https://brasilapi.com.br/api/cep/v1/" + cep)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer requisicao: %v\n", err)
		return
	}
	res, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
		return
	}
	var data dto.BrasilApi
	err = json.Unmarshal(res, &data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer o parse da resposta %v\n", err)
		return
	}
	resultChan <- data
}
