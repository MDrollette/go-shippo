package models

import "encoding/json"

const (
	AddressObjectSourceFullyEntered     = "FULLY_ENTERED"
	AddressObjectSourcePartiallyEntered = "PARTIALLY_ENTERED"
	AddressObjectSourceValidator        = "VALIDATOR"
)

// See https://goshippo.com/docs/reference#addresses
type AddressInput struct {
	ObjectID      string `json:"-"`
	ObjectPurpose string `json:"object_purpose"`
	Name          string `json:"name,omitempty"`
	Company       string `json:"company,omitempty"`
	Street1       string `json:"street1,omitempty"`
	Street2       string `json:"street2,omitempty"`
	City          string `json:"city,omitempty"`
	State         string `json:"state,omtiempty"`
	Zip           string `json:"zip,omitempty"`
	Country       string `json:"country"`
	Phone         string `json:"phone,omitempty"`
	Email         string `json:"email,omitempty"`
	IsResidential bool   `json:"is_residential"`
	Validate      bool   `json:"validate"`
	Metadata      string `json:"metadata,omitempty"`
}

func (a *AddressInput) MarshalJSON() ([]byte, error) {
	if a.ObjectID != "" {
		return json.Marshal(a.ObjectID)
	}

	return json.Marshal(*a)
}

func (a *AddressInput) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &a.ObjectID); nil == err {
		return nil
	}

	return json.Unmarshal(data, *a)
}

// See https://goshippo.com/docs/reference#addresses
type Address struct {
	AddressInput
	CommonOutputFields
	ObjectSource string           `json:"object_source"`
	Messages     []*OutputMessage `json:"messages"`
}
