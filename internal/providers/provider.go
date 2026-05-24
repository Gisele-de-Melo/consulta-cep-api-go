package providers

type CEPResponse struct {
	CEP        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Cidade     string `json:"cidade"`
	Estado     string `json:"estado"`
}

type Provider interface {
	GetCEP(cep string) (*CEPResponse, error)
}
