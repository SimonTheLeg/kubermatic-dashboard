// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/dashboard/v2/pkg/test/e2e/utils/apiclient/models"
)

// GetNodeDeploymentReader is a Reader for the GetNodeDeployment structure.
type GetNodeDeploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNodeDeploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNodeDeploymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetNodeDeploymentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetNodeDeploymentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetNodeDeploymentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNodeDeploymentOK creates a GetNodeDeploymentOK with default headers values
func NewGetNodeDeploymentOK() *GetNodeDeploymentOK {
	return &GetNodeDeploymentOK{}
}

/*
GetNodeDeploymentOK describes a response with status code 200, with default header values.

NodeDeployment
*/
type GetNodeDeploymentOK struct {
	Payload *models.NodeDeployment
}

// IsSuccess returns true when this get node deployment o k response has a 2xx status code
func (o *GetNodeDeploymentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get node deployment o k response has a 3xx status code
func (o *GetNodeDeploymentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get node deployment o k response has a 4xx status code
func (o *GetNodeDeploymentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get node deployment o k response has a 5xx status code
func (o *GetNodeDeploymentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get node deployment o k response a status code equal to that given
func (o *GetNodeDeploymentOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetNodeDeploymentOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments/{nodedeployment_id}][%d] getNodeDeploymentOK  %+v", 200, o.Payload)
}

func (o *GetNodeDeploymentOK) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments/{nodedeployment_id}][%d] getNodeDeploymentOK  %+v", 200, o.Payload)
}

func (o *GetNodeDeploymentOK) GetPayload() *models.NodeDeployment {
	return o.Payload
}

func (o *GetNodeDeploymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NodeDeployment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNodeDeploymentUnauthorized creates a GetNodeDeploymentUnauthorized with default headers values
func NewGetNodeDeploymentUnauthorized() *GetNodeDeploymentUnauthorized {
	return &GetNodeDeploymentUnauthorized{}
}

/*
GetNodeDeploymentUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetNodeDeploymentUnauthorized struct {
}

// IsSuccess returns true when this get node deployment unauthorized response has a 2xx status code
func (o *GetNodeDeploymentUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get node deployment unauthorized response has a 3xx status code
func (o *GetNodeDeploymentUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get node deployment unauthorized response has a 4xx status code
func (o *GetNodeDeploymentUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get node deployment unauthorized response has a 5xx status code
func (o *GetNodeDeploymentUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get node deployment unauthorized response a status code equal to that given
func (o *GetNodeDeploymentUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetNodeDeploymentUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments/{nodedeployment_id}][%d] getNodeDeploymentUnauthorized ", 401)
}

func (o *GetNodeDeploymentUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments/{nodedeployment_id}][%d] getNodeDeploymentUnauthorized ", 401)
}

func (o *GetNodeDeploymentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetNodeDeploymentForbidden creates a GetNodeDeploymentForbidden with default headers values
func NewGetNodeDeploymentForbidden() *GetNodeDeploymentForbidden {
	return &GetNodeDeploymentForbidden{}
}

/*
GetNodeDeploymentForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type GetNodeDeploymentForbidden struct {
}

// IsSuccess returns true when this get node deployment forbidden response has a 2xx status code
func (o *GetNodeDeploymentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get node deployment forbidden response has a 3xx status code
func (o *GetNodeDeploymentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get node deployment forbidden response has a 4xx status code
func (o *GetNodeDeploymentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get node deployment forbidden response has a 5xx status code
func (o *GetNodeDeploymentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get node deployment forbidden response a status code equal to that given
func (o *GetNodeDeploymentForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetNodeDeploymentForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments/{nodedeployment_id}][%d] getNodeDeploymentForbidden ", 403)
}

func (o *GetNodeDeploymentForbidden) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments/{nodedeployment_id}][%d] getNodeDeploymentForbidden ", 403)
}

func (o *GetNodeDeploymentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetNodeDeploymentDefault creates a GetNodeDeploymentDefault with default headers values
func NewGetNodeDeploymentDefault(code int) *GetNodeDeploymentDefault {
	return &GetNodeDeploymentDefault{
		_statusCode: code,
	}
}

/*
GetNodeDeploymentDefault describes a response with status code -1, with default header values.

errorResponse
*/
type GetNodeDeploymentDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get node deployment default response
func (o *GetNodeDeploymentDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get node deployment default response has a 2xx status code
func (o *GetNodeDeploymentDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get node deployment default response has a 3xx status code
func (o *GetNodeDeploymentDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get node deployment default response has a 4xx status code
func (o *GetNodeDeploymentDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get node deployment default response has a 5xx status code
func (o *GetNodeDeploymentDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get node deployment default response a status code equal to that given
func (o *GetNodeDeploymentDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetNodeDeploymentDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments/{nodedeployment_id}][%d] getNodeDeployment default  %+v", o._statusCode, o.Payload)
}

func (o *GetNodeDeploymentDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments/{nodedeployment_id}][%d] getNodeDeployment default  %+v", o._statusCode, o.Payload)
}

func (o *GetNodeDeploymentDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetNodeDeploymentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}