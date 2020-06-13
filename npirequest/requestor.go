package requestor

import "fmt"

// NpiRequest struct for the many different parameters.
type NpiRequest struct {
	Version             string `json:"version"`
	NpiNumber           string `json:"number"`
	EnumerationType     string `json:"enumeration_type"`
	TaxonomyDescription string `json:"taxonomy_description"`
	FirstName           string `json:"first_name"`
	LastName            string `json:"last_name"`
	OrganizationName    string `json:"organization_name"`
	AddressPurpose      string `json:"address_purpose"`
	City                string `json:"city"`
	State               string `json:"state"`
	PostalCode          string `json:"postal_code"`
	CountryCode         string `json:"country_code"`
	Limit               string `json:"limit"`
	Skip                string `json:"skip"`

	response interface{}
}

// Adresses Ignore
type Adresses struct {
	Addresses []Address
}

// Address Ignore
type Address struct {
	Address1        string `json:"address_1"`
	Address2        string `json:"address_2"`
	AddressPurpose  string `json:"address_purpose"`
	AddressType     string `json:"address_type"`
	City            string `json:"city"`
	CountryCode     string `json:"country_code"`
	CountryName     string `json:"country_name"`
	FaxNumber       string `json:"fax_number"`
	PostalCode      string `json:"postal_code"`
	State           string `json:"state"`
	TelephoneNumber string `json:"telephone_number"`
}

// Basic Ignore
type Basic struct {
	Basic BasicInfo
}

// BasicInfo Ignore
type BasicInfo struct {
	Credential      string `json:"credential"`
	EnumerationDate string `json:"enumeration_date"`
	FirstName       string `json:"first_name"`
	Gender          string `json:"gender"`
	LastName        string `json:"last_name"`
	LastUpdated     string `json:"last_updated"`
	MiddleName      string `json:"middle_name"`
	Name            string `json:"name"`
	NamePrefix      string `json:"name_prefix"`
	SoleProprietor  string `json:"sole_proprietor"`
	Status          string `json:"status"`
}

const registryURI = "https://npiregistry.cms.hhs.gov/api/"

// ToRequestString converts the NpiRequest struct to
// valid http request paramater string.
func (n *NpiRequest) ToRequestString() string {
	var reqVars = registryURI

	if len(n.Version) > 0 {
		reqVars += fmt.Sprintf("?version=%v", n.Version)
	} else {
		reqVars += fmt.Sprintf("?version=2.0")
	}
	if len(n.City) > 0 {
		reqVars += fmt.Sprintf("&city=%v", n.City)
	}
	if len(n.State) > 0 {
		reqVars += fmt.Sprintf("&state=%v", n.State)
	}
	if len(n.NpiNumber) > 0 {
		reqVars += fmt.Sprintf("&number=%v", n.NpiNumber)
	}
	if len(n.EnumerationType) > 0 {
		reqVars += fmt.Sprintf("&enumeration_type=%v", n.EnumerationType)
	}
	if len(n.TaxonomyDescription) > 0 {
		reqVars += fmt.Sprintf("&taxonomy_description=%v", n.TaxonomyDescription)
	}
	if len(n.FirstName) > 0 {
		reqVars += fmt.Sprintf("&first_name=%v", n.FirstName)
	}
	if len(n.LastName) > 0 {
		reqVars += fmt.Sprintf("&last_name=%v", n.LastName)
	}
	if len(n.OrganizationName) > 0 {
		reqVars += fmt.Sprintf("&organization_name=%v", n.OrganizationName)
	}
	if len(n.AddressPurpose) > 0 {
		reqVars += fmt.Sprintf("&address_purpose=%v", n.AddressPurpose)
	}
	if len(n.PostalCode) > 0 {
		reqVars += fmt.Sprintf("&postal_code=%v", n.PostalCode)
	}
	if len(n.CountryCode) > 0 {
		reqVars += fmt.Sprintf("&country_code=%v", n.CountryCode)
	}
	if len(n.Limit) > 0 {
		reqVars += fmt.Sprintf("&limit=%v", n.Limit)
	}
	if len(n.Skip) > 0 {
		reqVars += fmt.Sprintf("&skip=%v", n.Skip)
	}
	return reqVars
}
