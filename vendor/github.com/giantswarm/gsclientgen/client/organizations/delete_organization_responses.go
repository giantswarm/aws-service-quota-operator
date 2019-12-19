// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/giantswarm/gsclientgen/models"
)

// DeleteOrganizationReader is a Reader for the DeleteOrganization structure.
type DeleteOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeleteOrganizationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteOrganizationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewDeleteOrganizationConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteOrganizationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteOrganizationOK creates a DeleteOrganizationOK with default headers values
func NewDeleteOrganizationOK() *DeleteOrganizationOK {
	return &DeleteOrganizationOK{}
}

/*DeleteOrganizationOK handles this case with default header values.

Organization deleted
*/
type DeleteOrganizationOK struct {
	Payload *models.V4GenericResponse
}

func (o *DeleteOrganizationOK) Error() string {
	return fmt.Sprintf("[DELETE /v4/organizations/{organization_id}/][%d] deleteOrganizationOK  %+v", 200, o.Payload)
}

func (o *DeleteOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrganizationUnauthorized creates a DeleteOrganizationUnauthorized with default headers values
func NewDeleteOrganizationUnauthorized() *DeleteOrganizationUnauthorized {
	return &DeleteOrganizationUnauthorized{}
}

/*DeleteOrganizationUnauthorized handles this case with default header values.

Permission denied
*/
type DeleteOrganizationUnauthorized struct {
	Payload *models.V4GenericResponse
}

func (o *DeleteOrganizationUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v4/organizations/{organization_id}/][%d] deleteOrganizationUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteOrganizationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrganizationNotFound creates a DeleteOrganizationNotFound with default headers values
func NewDeleteOrganizationNotFound() *DeleteOrganizationNotFound {
	return &DeleteOrganizationNotFound{}
}

/*DeleteOrganizationNotFound handles this case with default header values.

Organization not found
*/
type DeleteOrganizationNotFound struct {
	Payload *models.V4GenericResponse
}

func (o *DeleteOrganizationNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v4/organizations/{organization_id}/][%d] deleteOrganizationNotFound  %+v", 404, o.Payload)
}

func (o *DeleteOrganizationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrganizationConflict creates a DeleteOrganizationConflict with default headers values
func NewDeleteOrganizationConflict() *DeleteOrganizationConflict {
	return &DeleteOrganizationConflict{}
}

/*DeleteOrganizationConflict handles this case with default header values.

Organization has credentials
*/
type DeleteOrganizationConflict struct {
	Payload *models.V4GenericResponse
}

func (o *DeleteOrganizationConflict) Error() string {
	return fmt.Sprintf("[DELETE /v4/organizations/{organization_id}/][%d] deleteOrganizationConflict  %+v", 409, o.Payload)
}

func (o *DeleteOrganizationConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrganizationDefault creates a DeleteOrganizationDefault with default headers values
func NewDeleteOrganizationDefault(code int) *DeleteOrganizationDefault {
	return &DeleteOrganizationDefault{
		_statusCode: code,
	}
}

/*DeleteOrganizationDefault handles this case with default header values.

Error
*/
type DeleteOrganizationDefault struct {
	_statusCode int

	Payload *models.V4GenericResponse
}

// Code gets the status code for the delete organization default response
func (o *DeleteOrganizationDefault) Code() int {
	return o._statusCode
}

func (o *DeleteOrganizationDefault) Error() string {
	return fmt.Sprintf("[DELETE /v4/organizations/{organization_id}/][%d] deleteOrganization default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteOrganizationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}