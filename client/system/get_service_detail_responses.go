// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/jenkins-x/anchore-client/models"
)

// GetServiceDetailReader is a Reader for the GetServiceDetail structure.
type GetServiceDetailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServiceDetailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetServiceDetailOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetServiceDetailInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetServiceDetailOK creates a GetServiceDetailOK with default headers values
func NewGetServiceDetailOK() *GetServiceDetailOK {
	return &GetServiceDetailOK{}
}

/*GetServiceDetailOK handles this case with default header values.

Status listing
*/
type GetServiceDetailOK struct {
	Payload *models.SystemStatusResponse
}

func (o *GetServiceDetailOK) Error() string {
	return fmt.Sprintf("[GET /system][%d] getServiceDetailOK  %+v", 200, o.Payload)
}

func (o *GetServiceDetailOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SystemStatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceDetailInternalServerError creates a GetServiceDetailInternalServerError with default headers values
func NewGetServiceDetailInternalServerError() *GetServiceDetailInternalServerError {
	return &GetServiceDetailInternalServerError{}
}

/*GetServiceDetailInternalServerError handles this case with default header values.

Internal error
*/
type GetServiceDetailInternalServerError struct {
	Payload *models.APIErrorResponse
}

func (o *GetServiceDetailInternalServerError) Error() string {
	return fmt.Sprintf("[GET /system][%d] getServiceDetailInternalServerError  %+v", 500, o.Payload)
}

func (o *GetServiceDetailInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}