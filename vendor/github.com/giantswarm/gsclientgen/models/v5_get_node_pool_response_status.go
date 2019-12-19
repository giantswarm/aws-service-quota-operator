// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// V5GetNodePoolResponseStatus Information on the current size and status of the node pool
// swagger:model v5GetNodePoolResponseStatus
type V5GetNodePoolResponseStatus struct {

	// Desired number of nodes in the pool
	Nodes int64 `json:"nodes,omitempty"`

	// Number of nodes in state NodeReady
	NodesReady int64 `json:"nodes_ready,omitempty"`
}

// Validate validates this v5 get node pool response status
func (m *V5GetNodePoolResponseStatus) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V5GetNodePoolResponseStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V5GetNodePoolResponseStatus) UnmarshalBinary(b []byte) error {
	var res V5GetNodePoolResponseStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
