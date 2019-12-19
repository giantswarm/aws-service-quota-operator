// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// V5ClusterDetailsResponseConditionsItems v5 cluster details response conditions items
// swagger:model v5ClusterDetailsResponseConditionsItems
type V5ClusterDetailsResponseConditionsItems struct {

	// A string describing the condition, e. g. 'Created'
	Condition string `json:"condition,omitempty"`

	// Date and time when the cluster transitioned into this condition
	LastTransitionTime string `json:"last_transition_time,omitempty"`
}

// Validate validates this v5 cluster details response conditions items
func (m *V5ClusterDetailsResponseConditionsItems) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V5ClusterDetailsResponseConditionsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V5ClusterDetailsResponseConditionsItems) UnmarshalBinary(b []byte) error {
	var res V5ClusterDetailsResponseConditionsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
