package providers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

type ViaCEPProvider struct{}

func formatCEPForViaCEP(cep string) string {
	// Remove tudo que não for número
	re := regexp.MustCompile(`\D`)
	digits := re.ReplaceAllString(cep, "")

	return digits
}

func (v ViaCEPProvider) GetCEP(cep string) (*CEPResponse, error) {
	cepf := formatCEPForViaCEP(cep)
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cepf)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("erro na API ViaCEP: status %d ", resp.StatusCode)
	}

	var data struct {
		CEP        string `json:"cep"`
		Logradouro string `json:"logradouro"`
		Bairro     string `json:"bairro"`
		Localidade string `json:"localidade"`
		UF         string `json:"uf"`
		Erro       bool   `json:"erro"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	if data.Erro {
		return nil, fmt.Errorf("cep não encontrado no ViaCEP")
	}

	return &CEPResponse{
		CEP:        data.CEP,
		Logradouro: data.Logradouro,
		Bairro:     data.Bairro,
		Cidade:     data.Localidade,
		Estado:     data.UF,
	}, nil
}
