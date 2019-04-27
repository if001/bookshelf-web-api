// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// OkResponse ok response
// swagger:model OkResponse
type OkResponse struct {

	// message
	Message string `json:"message,omitempty"`

	// result code
	ResultCode int32 `json:"resultCode,omitempty"`
}

// Validate validates this ok response
func (m *OkResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OkResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OkResponse) UnmarshalBinary(b []byte) error {
	var res OkResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
