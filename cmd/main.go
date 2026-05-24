package main

import (
	"consulta-cep-api/internal"
	"consulta-cep-api/internal/providers"
	"fmt"
)

func main() {
	service := internal.NewCEPService(
		providers.ViaCEPProvider{},
		providers.BrasilAPIProvider{},
		providers.WideNetProvider{},
	)
	//cep := "01001000"
	var cep string
	fmt.Print("Digite o CEP:")
	fmt.Scanln(&cep)

	result, err := service.FindCEP(cep)
	if err != nil {
		fmt.Println("Erro:", cep)
		return
	}

	fmt.Printf("CEP encontrado: %+v\n", result)
}
