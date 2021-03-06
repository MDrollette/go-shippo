package models

// See https://goshippo.com/docs/reference#refunds
type RefundInput struct {
	Transaction string `json:"transaction"`
	Async       bool   `json:"async"`
}

// See https://goshippo.com/docs/reference#refunds
type Refund struct {
	RefundInput
	CommonOutputFields
	ObjectStatus string `json:"object_status,omitempty"`
}
