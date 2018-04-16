// Code generated by go-swagger; DO NOT EDIT.

package catalog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/jenkins-x/anchore-client/models"
)

// AddRepositoryReader is a Reader for the AddRepository structure.
type AddRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddRepositoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddRepositoryOK creates a AddRepositoryOK with default headers values
func NewAddRepositoryOK() *AddRepositoryOK {
	return &AddRepositoryOK{}
}

/*AddRepositoryOK handles this case with default header values.

Repository and discovered tags added
*/
type AddRepositoryOK struct {
	Payload models.RepositoryTagList
}

func (o *AddRepositoryOK) Error() string {
	return fmt.Sprintf("[POST /repositories][%d] addRepositoryOK  %+v", 200, o.Payload)
}

func (o *AddRepositoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
