package providers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

type BrasilAPIProvider struct{}

func formatCEPForBrasilAPI(cep string) string {
	// Remove tudo que não for número
	re := regexp.MustCompile(`\D`)
	digits := re.ReplaceAllString(cep, "")

	return digits
}

func (b BrasilAPIProvider) GetCEP(cep string) (*CEPResponse, error) {
	cepf := formatCEPForBrasilAPI(cep)
	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cepf)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("erro na BrasilAPI: status %d", resp.StatusCode)
	}

	var data struct {
		CEP          string `json:"cep"`
		State        string `json:"state"`
		City         string `json:"city"`
		Neighborhood string `json:"neighborhood"`
		Street       string `json:"street"`
		Service      string `json:"service"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	if data.CEP == "" {
		return nil, fmt.Errorf("cep não encontrado na BrasilAPI")
	}

	return &CEPResponse{
		CEP:        data.CEP,
		Logradouro: data.Street,
		Bairro:     data.Neighborhood,
		Cidade:     data.City,
		Estado:     data.State,
	}, nil

}
