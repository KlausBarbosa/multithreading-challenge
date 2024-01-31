package dto

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"-"`
	Gia         string `json:"-"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"-"`
}

type BrasilApi struct {
	Cep     string `json:"cep"`
	Estado  string `json:"state"`
	Cidade  string `json:"city"`
	Bairro  string `json:"neighborhood"`
	Rua     string `json:"street"`
	Service string `json:"-"`
}
