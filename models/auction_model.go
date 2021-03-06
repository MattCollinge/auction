package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// AuctionModel this is the format of a Auction object
// swagger:model AuctionModel
type AuctionModel struct {

	// auction Id
	// Required: true
	AuctionID *string `json:"auctionId"`

	// auction item
	// Required: true
	AuctionItem *string `json:"auctionItem"`
}

// Validate validates this auction model
func (m *AuctionModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuctionID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAuctionItem(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuctionModel) validateAuctionID(formats strfmt.Registry) error {

	if err := validate.Required("auctionId", "body", m.AuctionID); err != nil {
		return err
	}

	return nil
}

func (m *AuctionModel) validateAuctionItem(formats strfmt.Registry) error {

	if err := validate.Required("auctionItem", "body", m.AuctionItem); err != nil {
		return err
	}

	return nil
}
