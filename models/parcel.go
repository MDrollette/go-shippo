package models

import "encoding/json"

// See https://goshippo.com/docs/reference#parcels
type ParcelInput struct {
	ObjectID     string       `json:"-"`
	Length       string       `json:"length"`
	Width        string       `json:"width"`
	Height       string       `json:"height"`
	DistanceUnit string       `json:"distance_unit"`
	Weight       string       `json:"weight"`
	MassUnit     string       `json:"mass_unit"`
	Template     string       `json:"template,omitempty"` // https://goshippo.com/docs/reference#parceltemplates
	Metadata     string       `json:"metadata,omitempty"`
	Extra        *ParcelExtra `json:"extra,omitempty"`
}

func (p *ParcelInput) MarshalJSON() ([]byte, error) {
	if p.ObjectID != "" {
		return json.Marshal(p.ObjectID)
	}

	return json.Marshal(*p)
}

func (p *ParcelInput) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &p.ObjectID); nil == err {
		return nil
	}

	return json.Unmarshal(data, *p)
}

type ParcelExtra struct {
	COD       *COD             `json:"COD"`
	Insurance *ParcelInsurance `json:"insurance"`
}

type ParcelInsurance struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
	Provider string `json:"provider"`
	Content  string `json:"content"`
}

// See https://goshippo.com/docs/reference#parcels
type Parcel struct {
	ParcelInput
	CommonOutputFields
}
