package libs

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

// ExternalIP structure for ipinfo JSON
type ExternalIP struct {
	IP       string
	Hostname string
	City     string
	Region   string
	Country  string
	Loc      string
	Org      string
	Postal   string
	Timezone string
}

// PrintData prints ExternalIP fields with respective values
func (c *ExternalIP) PrintData() {
	fmt.Println("IP Externo:", c.IP)
	fmt.Println("Hostname:", c.Hostname)
	fmt.Println("Cidade:", c.City)
	fmt.Println("Região:", c.Region)
	fmt.Println("País:", c.Country)
	fmt.Println("Localização:", c.Loc)
	fmt.Println("Organização:", c.Org)
	fmt.Println("Código Postal:", c.Postal)
	fmt.Println("Fuso-Horário:", c.Timezone)
}

// SetJSONData sets ExternalIP fields with JSON data
func (c *ExternalIP) SetJSONData(data []byte) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	json.Unmarshal(data, c)
}