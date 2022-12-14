// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/peeush-agarwal/learn-golang/product-api-client/models"
)

// UpdateProductReader is a Reader for the UpdateProduct structure.
type UpdateProductReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateProductReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateProductOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateProductOK creates a UpdateProductOK with default headers values
func NewUpdateProductOK() *UpdateProductOK {
	return &UpdateProductOK{}
}

/*
UpdateProductOK describes a response with status code 200, with default header values.

A single product returned in the response
*/
type UpdateProductOK struct {
	Payload *models.Product
}

// IsSuccess returns true when this update product o k response has a 2xx status code
func (o *UpdateProductOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update product o k response has a 3xx status code
func (o *UpdateProductOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update product o k response has a 4xx status code
func (o *UpdateProductOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update product o k response has a 5xx status code
func (o *UpdateProductOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update product o k response a status code equal to that given
func (o *UpdateProductOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateProductOK) Error() string {
	return fmt.Sprintf("[PUT /products/{id}][%d] updateProductOK  %+v", 200, o.Payload)
}

func (o *UpdateProductOK) String() string {
	return fmt.Sprintf("[PUT /products/{id}][%d] updateProductOK  %+v", 200, o.Payload)
}

func (o *UpdateProductOK) GetPayload() *models.Product {
	return o.Payload
}

func (o *UpdateProductOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Product)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
