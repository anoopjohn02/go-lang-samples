package models

type DeviceProfile struct {
	UserId         string   `json:"userId"`
	UserName       string   `json:"userName"`
	ClientId       string   `json:"clientId"`
	CompanyId      string   `json:"companyId"`
	CompanyUser    bool     `json:"companyUser"`
	Roles          []string `json:"roles"`
	ManufacturerId string   `json:"manufacturerId"`
	Token          string
}
