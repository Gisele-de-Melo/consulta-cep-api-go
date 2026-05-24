package internal

import (
	"consulta-cep-api/internal/providers"
	"errors"
)

type CEPService struct {
	providers []providers.Provider
}

func NewCEPService(p ...providers.Provider) *CEPService {
	return &CEPService{providers: p}
}

func (s *CEPService) FindCEP(cep string) (*providers.CEPResponse, error) {
	for _, provider := range s.providers {
		result, err := provider.GetCEP(cep)
		if err == nil {
			return result, nil
		}
	}
	return nil, errors.New("nenhuma API retornou resultado válido")
}
