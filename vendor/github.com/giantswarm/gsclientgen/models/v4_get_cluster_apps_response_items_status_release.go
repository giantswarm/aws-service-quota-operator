// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// V4GetClusterAppsResponseItemsStatusRelease v4 get cluster apps response items status release
// swagger:model v4GetClusterAppsResponseItemsStatusRelease
type V4GetClusterAppsResponseItemsStatusRelease struct {

	// Date and time that the app was last last deployed
	LastDeployed string `json:"last_deployed,omitempty"`

	// A string representing the status of the app. (Can be: empty, DEPLOYED or FAILED)
	Status string `json:"status,omitempty"`
}

// Validate validates this v4 get cluster apps response items status release
func (m *V4GetClusterAppsResponseItemsStatusRelease) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V4GetClusterAppsResponseItemsStatusRelease) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V4GetClusterAppsResponseItemsStatusRelease) UnmarshalBinary(b []byte) error {
	var res V4GetClusterAppsResponseItemsStatusRelease
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
