// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// PostImagesReader is a Reader for the PostImages structure.
type PostImagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostImagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostImagesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostImagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostImagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostImagesCreated creates a PostImagesCreated with default headers values
func NewPostImagesCreated() *PostImagesCreated {
	return &PostImagesCreated{}
}

/*PostImagesCreated handles this case with default header values.

dummy
*/
type PostImagesCreated struct {
	/*Upload URL for image file
	 */
	Location string

	Payload *models.Image
}

func (o *PostImagesCreated) Error() string {
	return fmt.Sprintf("[POST /images][%d] postImagesCreated  %+v", 201, o.Payload)
}

func (o *PostImagesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	o.Payload = new(models.Image)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostImagesBadRequest creates a PostImagesBadRequest with default headers values
func NewPostImagesBadRequest() *PostImagesBadRequest {
	return &PostImagesBadRequest{}
}

/*PostImagesBadRequest handles this case with default header values.

error
*/
type PostImagesBadRequest struct {
	Payload *models.Riaaserror
}

func (o *PostImagesBadRequest) Error() string {
	return fmt.Sprintf("[POST /images][%d] postImagesBadRequest  %+v", 400, o.Payload)
}

func (o *PostImagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostImagesInternalServerError creates a PostImagesInternalServerError with default headers values
func NewPostImagesInternalServerError() *PostImagesInternalServerError {
	return &PostImagesInternalServerError{}
}

/*PostImagesInternalServerError handles this case with default header values.

error
*/
type PostImagesInternalServerError struct {
	Payload *models.Riaaserror
}

func (o *PostImagesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /images][%d] postImagesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostImagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
