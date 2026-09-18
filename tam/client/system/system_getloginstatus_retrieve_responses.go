package system

import (
	"encoding/json"
	stderrors "errors"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SystemGetLoginStatusRetrieveReader is a Reader for the SystemGetLoginStatusRetrieve structure.
type SystemGetLoginStatusRetrieveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SystemGetLoginStatusRetrieveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (any, error) {
	switch response.Code() {
	case 200:
		result := NewSystemGetLoginStatusRetrieveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSystemGetLoginStatusRetrieveDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSystemGetLoginStatusRetrieveOK creates a SystemGetLoginStatusRetrieveOK with default header values
func NewSystemGetLoginStatusRetrieveOK() *SystemGetLoginStatusRetrieveOK {
	return &SystemGetLoginStatusRetrieveOK{}
}

// SystemGetLoginStatusRetrieveOK describes a response with status code 200, with default header values.
//
// SystemGetLoginStatusRetrieveOK system getLoginStatus retrieves o k
type SystemGetLoginStatusRetrieveOK struct {
	Payload map[string]any
}

// IsSuccess returns true when this system getLoginStatus retrieve o k response has a 2xx status code
func (o *SystemGetLoginStatusRetrieveOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this system getLoginStatus retrieve o k response has a 3xx status code.
func (o *SystemGetLoginStatusRetrieveOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this system getLoginStatus o k response has a 4xx status code.
func (o *SystemGetLoginStatusRetrieveOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this system getLoginStatus o k response has a 5xx status code.
func (o *SystemGetLoginStatusRetrieveOK) IsServerError() bool {
	return false
}

// IsCode returns true when the system getLoginStatus o k response has a status code equal to that given
func (o *SystemGetLoginStatusRetrieveOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the system getLoginStatus retrieve o k response
func (o *SystemGetLoginStatusRetrieveOK) Code() int {
	return 200
}

func (o *SystemGetLoginStatusRetrieveOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/System/getLoginStatus][%d] systemGetLoginStatusRetrieveOK %s", 200, payload)
}

func (o *SystemGetLoginStatusRetrieveOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/System/getLoginStatus][%d] systemGetLoginStatusRetrieveOK %s", 200, payload)
}

func (o *SystemGetLoginStatusRetrieveOK) GetPayload() map[string]any {
	return o.Payload
}

func (o *SystemGetLoginStatusRetrieveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && !stderrors.Is(err, io.EOF) {
		return err
	}

	return nil
}

// NewSystemGetLoginStatusRetrieveDefault creates a SystemGetLoginStatusRetrieveDefault with default header values
func NewSystemGetLoginStatusRetrieveDefault(code int) *SystemGetLoginStatusRetrieveDefault {
	return &SystemGetLoginStatusRetrieveDefault{
		_statusCode: code,
	}
}

// SystemGetLoginStatusRetrieveDefault describes a response with status code -1, with default header values.
//
// SystemGetLoginStatusRetrieveDefault system getLoginStatus default
type SystemGetLoginStatusRetrieveDefault struct {
	_statusCode int

	Payload any
}

// IsSuccess returns true when this system getLoginStatus retrieve default response has a 2xx status code
func (o *SystemGetLoginStatusRetrieveDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this system getLoginStatus retrieve default response has a 3xx status code
func (o *SystemGetLoginStatusRetrieveDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this system getLoginStatus retrieve default response has a 4xx status code
func (o *SystemGetLoginStatusRetrieveDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this system getLoginStatus retrieve default response has a 5xx status code
func (o *SystemGetLoginStatusRetrieveDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this system getLoginStatus retrieve default response a status code equal to that given
func (o *SystemGetLoginStatusRetrieveDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the system getLoginStatus default response code
func (o *SystemGetLoginStatusRetrieveDefault) Code() int {
	return o._statusCode
}

func (o *SystemGetLoginStatusRetrieveDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/System/getLoginStatus][%d] systemGetLoginStatusRetrieveDefault %s", o._statusCode, payload)
}

func (o *SystemGetLoginStatusRetrieveDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/System/getLoginStatus][%d] systemGetLoginStatusRetrieveDefault %s", o._statusCode, payload)
}

func (o *SystemGetLoginStatusRetrieveDefault) GetPayload() any {
	return o.Payload
}

func (o *SystemGetLoginStatusRetrieveDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && !stderrors.Is(err, io.EOF) {
		return err
	}

	return nil
}
