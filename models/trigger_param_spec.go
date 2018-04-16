// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TriggerParamSpec trigger param spec
// swagger:model TriggerParamSpec
type TriggerParamSpec struct {

	// description
	Description string `json:"description,omitempty"`

	// Parameter name as it appears in policy document
	Name string `json:"name,omitempty"`

	// Is this a required parameter or optional
	Required bool `json:"required,omitempty"`

	// If present, a definition for validation of input. Typically a jsonschema object that can be used to validate an input against.
	Validator interface{} `json:"validator,omitempty"`
}

// Validate validates this trigger param spec
func (m *TriggerParamSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *TriggerParamSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TriggerParamSpec) UnmarshalBinary(b []byte) error {
	var res TriggerParamSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}