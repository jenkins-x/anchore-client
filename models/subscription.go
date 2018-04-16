// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Subscription Subscription entry
// swagger:model Subscription
type Subscription struct {

	// Is the subscription currently active
	Active bool `json:"active,omitempty"`

	// the unique id for this subscription record
	SubscriptionID string `json:"subscription_id,omitempty"`

	// The key value that the subscription references. E.g. a tag value or a repo name.
	SubscriptionKey string `json:"subscription_key,omitempty"`

	// The type of the subscription
	SubscriptionType string `json:"subscription_type,omitempty"`

	// The value of the subscription target
	SubscriptionValue string `json:"subscription_value,omitempty"`

	// The userId of the subscribed user
	UserID string `json:"userId,omitempty"`
}

// Validate validates this subscription
func (m *Subscription) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSubscriptionType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var subscriptionTypeSubscriptionTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["policy_eval","tag_update","vuln_update","repo_update"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		subscriptionTypeSubscriptionTypePropEnum = append(subscriptionTypeSubscriptionTypePropEnum, v)
	}
}

const (
	// SubscriptionSubscriptionTypePolicyEval captures enum value "policy_eval"
	SubscriptionSubscriptionTypePolicyEval string = "policy_eval"
	// SubscriptionSubscriptionTypeTagUpdate captures enum value "tag_update"
	SubscriptionSubscriptionTypeTagUpdate string = "tag_update"
	// SubscriptionSubscriptionTypeVulnUpdate captures enum value "vuln_update"
	SubscriptionSubscriptionTypeVulnUpdate string = "vuln_update"
	// SubscriptionSubscriptionTypeRepoUpdate captures enum value "repo_update"
	SubscriptionSubscriptionTypeRepoUpdate string = "repo_update"
)

// prop value enum
func (m *Subscription) validateSubscriptionTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, subscriptionTypeSubscriptionTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Subscription) validateSubscriptionType(formats strfmt.Registry) error {

	if swag.IsZero(m.SubscriptionType) { // not required
		return nil
	}

	// value enum
	if err := m.validateSubscriptionTypeEnum("subscription_type", "body", m.SubscriptionType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Subscription) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Subscription) UnmarshalBinary(b []byte) error {
	var res Subscription
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
