// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// PostSecurityGroupsSecurityGroupIDRulesReader is a Reader for the PostSecurityGroupsSecurityGroupIDRules structure.
type PostSecurityGroupsSecurityGroupIDRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSecurityGroupsSecurityGroupIDRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostSecurityGroupsSecurityGroupIDRulesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostSecurityGroupsSecurityGroupIDRulesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostSecurityGroupsSecurityGroupIDRulesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostSecurityGroupsSecurityGroupIDRulesCreated creates a PostSecurityGroupsSecurityGroupIDRulesCreated with default headers values
func NewPostSecurityGroupsSecurityGroupIDRulesCreated() *PostSecurityGroupsSecurityGroupIDRulesCreated {
	return &PostSecurityGroupsSecurityGroupIDRulesCreated{}
}

/*PostSecurityGroupsSecurityGroupIDRulesCreated handles this case with default header values.

dummy
*/
type PostSecurityGroupsSecurityGroupIDRulesCreated struct {
	Payload *models.SecurityGroupRule
}

func (o *PostSecurityGroupsSecurityGroupIDRulesCreated) Error() string {
	return fmt.Sprintf("[POST /security_groups/{security_group_id}/rules][%d] postSecurityGroupsSecurityGroupIdRulesCreated  %+v", 201, o.Payload)
}

func (o *PostSecurityGroupsSecurityGroupIDRulesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecurityGroupRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSecurityGroupsSecurityGroupIDRulesBadRequest creates a PostSecurityGroupsSecurityGroupIDRulesBadRequest with default headers values
func NewPostSecurityGroupsSecurityGroupIDRulesBadRequest() *PostSecurityGroupsSecurityGroupIDRulesBadRequest {
	return &PostSecurityGroupsSecurityGroupIDRulesBadRequest{}
}

/*PostSecurityGroupsSecurityGroupIDRulesBadRequest handles this case with default header values.

error
*/
type PostSecurityGroupsSecurityGroupIDRulesBadRequest struct {
	Payload *models.Riaaserror
}

func (o *PostSecurityGroupsSecurityGroupIDRulesBadRequest) Error() string {
	return fmt.Sprintf("[POST /security_groups/{security_group_id}/rules][%d] postSecurityGroupsSecurityGroupIdRulesBadRequest  %+v", 400, o.Payload)
}

func (o *PostSecurityGroupsSecurityGroupIDRulesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSecurityGroupsSecurityGroupIDRulesInternalServerError creates a PostSecurityGroupsSecurityGroupIDRulesInternalServerError with default headers values
func NewPostSecurityGroupsSecurityGroupIDRulesInternalServerError() *PostSecurityGroupsSecurityGroupIDRulesInternalServerError {
	return &PostSecurityGroupsSecurityGroupIDRulesInternalServerError{}
}

/*PostSecurityGroupsSecurityGroupIDRulesInternalServerError handles this case with default header values.

error
*/
type PostSecurityGroupsSecurityGroupIDRulesInternalServerError struct {
	Payload *models.Riaaserror
}

func (o *PostSecurityGroupsSecurityGroupIDRulesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /security_groups/{security_group_id}/rules][%d] postSecurityGroupsSecurityGroupIdRulesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSecurityGroupsSecurityGroupIDRulesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostSecurityGroupsSecurityGroupIDRulesBody SecurityGroupRuleTemplate
swagger:model PostSecurityGroupsSecurityGroupIDRulesBody
*/
type PostSecurityGroupsSecurityGroupIDRulesBody struct {

	// The ICMP traffic code to allow. Valid values from 0 to 255. If unspecified, all codes are allowed. This can only be specified if type is also specified.
	Code *int64 `json:"code,omitempty"`

	// The direction of traffic to enforce (ingress, egress)
	// Required: true
	// Enum: [ingress egress]
	Direction *string `json:"direction"`

	// The IP version to enforce (ipv4, ipv6). The format of 'remote.address' or 'remote.cidr_block' must match this field, if they are used. Also, if 'remote' references another security group (ie. using remote.id, remote.name, remote.crn) then this rule will only apply to IP addresses (network interfaces) in that group which match this ip_version.
	// Enum: [ipv4 ipv6]
	IPVersion string `json:"ip_version,omitempty"`

	// The highest port in the range of ports to be matched; if unspecified, `65535` is used.
	PortMax *int64 `json:"port_max,omitempty"`

	// The lowest port in the range of ports to be matched; if unspecified, `1` is used.
	PortMin *int64 `json:"port_min,omitempty"`

	// The protocol to enforce. Must be one of (icmp, tcp, udp, all). Defaults to 'all' if omitted.
	// Required: true
	// Enum: [all icmp tcp udp]
	Protocol *string `json:"protocol"`

	// remote
	Remote *PostSecurityGroupsSecurityGroupIDRulesParamsBodyRemote `json:"remote,omitempty"`

	// The ICMP traffic type to allow. Valid values from 0 to 254. If unspecified, all types are allowed by this rule.
	Type *int64 `json:"type,omitempty"`
}

// Validate validates this post security groups security group ID rules body
func (o *PostSecurityGroupsSecurityGroupIDRulesBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIPVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRemote(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var postSecurityGroupsSecurityGroupIdRulesBodyTypeDirectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ingress","egress"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postSecurityGroupsSecurityGroupIdRulesBodyTypeDirectionPropEnum = append(postSecurityGroupsSecurityGroupIdRulesBodyTypeDirectionPropEnum, v)
	}
}

const (

	// PostSecurityGroupsSecurityGroupIDRulesBodyDirectionIngress captures enum value "ingress"
	PostSecurityGroupsSecurityGroupIDRulesBodyDirectionIngress string = "ingress"

	// PostSecurityGroupsSecurityGroupIDRulesBodyDirectionEgress captures enum value "egress"
	PostSecurityGroupsSecurityGroupIDRulesBodyDirectionEgress string = "egress"
)

// prop value enum
func (o *PostSecurityGroupsSecurityGroupIDRulesBody) validateDirectionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, postSecurityGroupsSecurityGroupIdRulesBodyTypeDirectionPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *PostSecurityGroupsSecurityGroupIDRulesBody) validateDirection(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"direction", "body", o.Direction); err != nil {
		return err
	}

	// value enum
	if err := o.validateDirectionEnum("body"+"."+"direction", "body", *o.Direction); err != nil {
		return err
	}

	return nil
}

var postSecurityGroupsSecurityGroupIdRulesBodyTypeIPVersionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ipv4","ipv6"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postSecurityGroupsSecurityGroupIdRulesBodyTypeIPVersionPropEnum = append(postSecurityGroupsSecurityGroupIdRulesBodyTypeIPVersionPropEnum, v)
	}
}

const (

	// PostSecurityGroupsSecurityGroupIDRulesBodyIPVersionIPV4 captures enum value "ipv4"
	PostSecurityGroupsSecurityGroupIDRulesBodyIPVersionIPV4 string = "ipv4"

	// PostSecurityGroupsSecurityGroupIDRulesBodyIPVersionIPV6 captures enum value "ipv6"
	PostSecurityGroupsSecurityGroupIDRulesBodyIPVersionIPV6 string = "ipv6"
)

// prop value enum
func (o *PostSecurityGroupsSecurityGroupIDRulesBody) validateIPVersionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, postSecurityGroupsSecurityGroupIdRulesBodyTypeIPVersionPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *PostSecurityGroupsSecurityGroupIDRulesBody) validateIPVersion(formats strfmt.Registry) error {

	if swag.IsZero(o.IPVersion) { // not required
		return nil
	}

	// value enum
	if err := o.validateIPVersionEnum("body"+"."+"ip_version", "body", o.IPVersion); err != nil {
		return err
	}

	return nil
}

var postSecurityGroupsSecurityGroupIdRulesBodyTypeProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["all","icmp","tcp","udp"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postSecurityGroupsSecurityGroupIdRulesBodyTypeProtocolPropEnum = append(postSecurityGroupsSecurityGroupIdRulesBodyTypeProtocolPropEnum, v)
	}
}

const (

	// PostSecurityGroupsSecurityGroupIDRulesBodyProtocolAll captures enum value "all"
	PostSecurityGroupsSecurityGroupIDRulesBodyProtocolAll string = "all"

	// PostSecurityGroupsSecurityGroupIDRulesBodyProtocolIcmp captures enum value "icmp"
	PostSecurityGroupsSecurityGroupIDRulesBodyProtocolIcmp string = "icmp"

	// PostSecurityGroupsSecurityGroupIDRulesBodyProtocolTCP captures enum value "tcp"
	PostSecurityGroupsSecurityGroupIDRulesBodyProtocolTCP string = "tcp"

	// PostSecurityGroupsSecurityGroupIDRulesBodyProtocolUDP captures enum value "udp"
	PostSecurityGroupsSecurityGroupIDRulesBodyProtocolUDP string = "udp"
)

// prop value enum
func (o *PostSecurityGroupsSecurityGroupIDRulesBody) validateProtocolEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, postSecurityGroupsSecurityGroupIdRulesBodyTypeProtocolPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *PostSecurityGroupsSecurityGroupIDRulesBody) validateProtocol(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"protocol", "body", o.Protocol); err != nil {
		return err
	}

	// value enum
	if err := o.validateProtocolEnum("body"+"."+"protocol", "body", *o.Protocol); err != nil {
		return err
	}

	return nil
}

func (o *PostSecurityGroupsSecurityGroupIDRulesBody) validateRemote(formats strfmt.Registry) error {

	if swag.IsZero(o.Remote) { // not required
		return nil
	}

	if o.Remote != nil {
		if err := o.Remote.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "remote")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostSecurityGroupsSecurityGroupIDRulesBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostSecurityGroupsSecurityGroupIDRulesBody) UnmarshalBinary(b []byte) error {
	var res PostSecurityGroupsSecurityGroupIDRulesBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostSecurityGroupsSecurityGroupIDRulesParamsBodyRemote SecurityGroupIdentityByName
//
// Uniquely identifies a security group using any one of ID, CRN, or name.
swagger:model PostSecurityGroupsSecurityGroupIDRulesParamsBodyRemote
*/
type PostSecurityGroupsSecurityGroupIDRulesParamsBodyRemote struct {

	// A single IPv4 or IPv6 address.
	Address string `json:"address,omitempty"`

	// A range of IPv4 or IPv6 addresses, in CIDR format.
	CidrBlock string `json:"cidr_block,omitempty"`

	// The security group's CRN
	Crn string `json:"crn,omitempty"`

	// The security group's unique identifier.
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`
}

// Validate validates this post security groups security group ID rules params body remote
func (o *PostSecurityGroupsSecurityGroupIDRulesParamsBodyRemote) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostSecurityGroupsSecurityGroupIDRulesParamsBodyRemote) validateID(formats strfmt.Registry) error {

	if swag.IsZero(o.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("body"+"."+"remote"+"."+"id", "body", "uuid", o.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostSecurityGroupsSecurityGroupIDRulesParamsBodyRemote) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostSecurityGroupsSecurityGroupIDRulesParamsBodyRemote) UnmarshalBinary(b []byte) error {
	var res PostSecurityGroupsSecurityGroupIDRulesParamsBodyRemote
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
