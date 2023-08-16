package model

type RegistrationData struct {
	FullName     string `json:"full_name"`
	TradeName    string `json:"trade_name"`
	Code         string `json:"code"`
	PersonType   string `json:"person_type"`
	CPForCNPJ    string `json:"cpf_or_cnpj"`
	TaxCode      string `json:"tax_code"`
	ClientSince  string `json:"client_since"`
	MunicipalReg string `json:"municipal_registration"`
	ExemptFromIE bool   `json:"exempt_from_ie"`
}

type CompanyData struct {
	RegistrationInfo RegistrationData `json:"registration_info"`
	AddressInfo      Address          `json:"address_info"`
	Peoples          []Person         `json:"people"`
	Admin            *Person          `json:"administrator,omitempty"`
}
