// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// GetNetworkAclsReader is a Reader for the GetNetworkAcls structure.
type GetNetworkAclsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkAclsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNetworkAclsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetNetworkAclsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNetworkAclsOK creates a GetNetworkAclsOK with default headers values
func NewGetNetworkAclsOK() *GetNetworkAclsOK {
	return &GetNetworkAclsOK{}
}

/*GetNetworkAclsOK handles this case with default header values.

dummy
*/
type GetNetworkAclsOK struct {
	Payload *GetNetworkAclsOKBody
}

func (o *GetNetworkAclsOK) Error() string {
	return fmt.Sprintf("[GET /network_acls][%d] getNetworkAclsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkAclsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetNetworkAclsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworkAclsInternalServerError creates a GetNetworkAclsInternalServerError with default headers values
func NewGetNetworkAclsInternalServerError() *GetNetworkAclsInternalServerError {
	return &GetNetworkAclsInternalServerError{}
}

/*GetNetworkAclsInternalServerError handles this case with default header values.

error
*/
type GetNetworkAclsInternalServerError struct {
	Payload *models.Riaaserror
}

func (o *GetNetworkAclsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /network_acls][%d] getNetworkAclsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetNetworkAclsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetNetworkAclsOKBody NetworkACLCollection
swagger:model GetNetworkAclsOKBody
*/
type GetNetworkAclsOKBody struct {

	// first
	First *GetNetworkAclsOKBodyFirst `json:"first,omitempty"`

	// The maximum number of resources can be returned by the request
	// Maximum: 100
	// Minimum: 1
	Limit int64 `json:"limit,omitempty"`

	// Collection of network ACLs
	// Required: true
	NetworkAcls []*models.NetworkACL `json:"network_acls"`

	// next
	Next *GetNetworkAclsOKBodyNext `json:"next,omitempty"`

	// The total number of resources across all pages
	// Minimum: 0
	TotalCount *int64 `json:"total_count,omitempty"`
}

// Validate validates this get network acls o k body
func (o *GetNetworkAclsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFirst(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNetworkAcls(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTotalCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkAclsOKBody) validateFirst(formats strfmt.Registry) error {

	if swag.IsZero(o.First) { // not required
		return nil
	}

	if o.First != nil {
		if err := o.First.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getNetworkAclsOK" + "." + "first")
			}
			return err
		}
	}

	return nil
}

func (o *GetNetworkAclsOKBody) validateLimit(formats strfmt.Registry) error {

	if swag.IsZero(o.Limit) { // not required
		return nil
	}

	if err := validate.MinimumInt("getNetworkAclsOK"+"."+"limit", "body", int64(o.Limit), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("getNetworkAclsOK"+"."+"limit", "body", int64(o.Limit), 100, false); err != nil {
		return err
	}

	return nil
}

func (o *GetNetworkAclsOKBody) validateNetworkAcls(formats strfmt.Registry) error {

	if err := validate.Required("getNetworkAclsOK"+"."+"network_acls", "body", o.NetworkAcls); err != nil {
		return err
	}

	for i := 0; i < len(o.NetworkAcls); i++ {
		if swag.IsZero(o.NetworkAcls[i]) { // not required
			continue
		}

		if o.NetworkAcls[i] != nil {
			if err := o.NetworkAcls[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getNetworkAclsOK" + "." + "network_acls" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetNetworkAclsOKBody) validateNext(formats strfmt.Registry) error {

	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if o.Next != nil {
		if err := o.Next.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getNetworkAclsOK" + "." + "next")
			}
			return err
		}
	}

	return nil
}

func (o *GetNetworkAclsOKBody) validateTotalCount(formats strfmt.Registry) error {

	if swag.IsZero(o.TotalCount) { // not required
		return nil
	}

	if err := validate.MinimumInt("getNetworkAclsOK"+"."+"total_count", "body", int64(*o.TotalCount), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkAclsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkAclsOKBody) UnmarshalBinary(b []byte) error {
	var res GetNetworkAclsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetNetworkAclsOKBodyFirst A reference to the first page of resources
swagger:model GetNetworkAclsOKBodyFirst
*/
type GetNetworkAclsOKBodyFirst struct {

	// The URL for the first page of resources
	// Required: true
	// Pattern: ^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$
	Href *string `json:"href"`
}

// Validate validates this get network acls o k body first
func (o *GetNetworkAclsOKBodyFirst) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkAclsOKBodyFirst) validateHref(formats strfmt.Registry) error {

	if err := validate.Required("getNetworkAclsOK"+"."+"first"+"."+"href", "body", o.Href); err != nil {
		return err
	}

	if err := validate.Pattern("getNetworkAclsOK"+"."+"first"+"."+"href", "body", string(*o.Href), `^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkAclsOKBodyFirst) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkAclsOKBodyFirst) UnmarshalBinary(b []byte) error {
	var res GetNetworkAclsOKBodyFirst
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetNetworkAclsOKBodyNext A reference to the next page of resources; this reference is included for all pages except the last page
swagger:model GetNetworkAclsOKBodyNext
*/
type GetNetworkAclsOKBodyNext struct {

	// The URL for the next page of resources
	// Required: true
	// Pattern: ^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$
	Href *string `json:"href"`
}

// Validate validates this get network acls o k body next
func (o *GetNetworkAclsOKBodyNext) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkAclsOKBodyNext) validateHref(formats strfmt.Registry) error {

	if err := validate.Required("getNetworkAclsOK"+"."+"next"+"."+"href", "body", o.Href); err != nil {
		return err
	}

	if err := validate.Pattern("getNetworkAclsOK"+"."+"next"+"."+"href", "body", string(*o.Href), `^http(s)?:\/\/([^\/?#]*)([^?#]*)(\?([^#]*))?(#(.*))?$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkAclsOKBodyNext) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkAclsOKBodyNext) UnmarshalBinary(b []byte) error {
	var res GetNetworkAclsOKBodyNext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}