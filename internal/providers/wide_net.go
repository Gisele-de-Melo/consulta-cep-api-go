package providers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

type WideNetProvider struct{}

func formatCEPForWideNet(cep string) string {
	// Remove tudo que não for número
	re := regexp.MustCompile(`\D`)
	digits := re.ReplaceAllString(cep, "")

	// Se não tiver 8 dígitos, retorna como está
	if len(digits) != 8 {
		return cep
	}

	// Insere o hífen: 12345678 → 12345-678
	return digits[:5] + "-" + digits[5:]
}

func (w WideNetProvider) GetCEP(cep string) (*CEPResponse, error) {
	cepf := formatCEPForWideNet(cep)
	url := fmt.Sprintf("https://cdn.apicep.com/file/apicep/%s.json", cepf)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("erro na WideNet: status %d", resp.StatusCode)
	}

	var data struct {
		Status   int    `json:"status"`
		Code     string `json:"code"`
		State    string `json:"state"`
		City     string `json:"city"`
		District string `json:"district"`
		Address  string `json:"address"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	if data.Code == "" {
		return nil, fmt.Errorf("cep não encontrado na WideNet")
	}

	return &CEPResponse{
		CEP:        data.Code,
		Logradouro: data.Address,
		Bairro:     data.District,
		Cidade:     data.City,
		Estado:     data.State,
	}, nil
}
