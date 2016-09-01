package client

import (
	"net/http"

	"github.com/d5/go-shippo/models"
)

// GetTrackingUpdate requests the tracking status of a shipment.
func (c *Client) GetTrackingUpdate(carrier, trackingNumber string) (*models.TrackingStatusOutput, error) {
	output := &models.TrackingStatusOutput{}
	err := c.do(http.MethodGet, "/trakcs/"+carrier+"/"+trackingNumber, nil, output)
	return output, err
}

// RegisterTrackingWebhook registers a tracking webhook.
// TODO: documentation on this API endpoint is not clear.
// https://goshippo.com/docs/reference#track-create
func (c *Client) RegisterTrackingWebhook(carrier, trackingNumber string) (*models.TrackingStatusOutput, error) {
	output := &models.TrackingStatusOutput{}
	err := c.do(http.MethodPost, "/trakcs/"+carrier+"/"+trackingNumber, nil, output)
	return output, err
}