// Code generated by go-swagger; DO NOT EDIT.

package v_p_naa_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// PostVpnGatewaysVpnGatewayIDConnectionsReader is a Reader for the PostVpnGatewaysVpnGatewayIDConnections structure.
type PostVpnGatewaysVpnGatewayIDConnectionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostVpnGatewaysVpnGatewayIDConnectionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostVpnGatewaysVpnGatewayIDConnectionsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostVpnGatewaysVpnGatewayIDConnectionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostVpnGatewaysVpnGatewayIDConnectionsCreated creates a PostVpnGatewaysVpnGatewayIDConnectionsCreated with default headers values
func NewPostVpnGatewaysVpnGatewayIDConnectionsCreated() *PostVpnGatewaysVpnGatewayIDConnectionsCreated {
	return &PostVpnGatewaysVpnGatewayIDConnectionsCreated{}
}

/*PostVpnGatewaysVpnGatewayIDConnectionsCreated handles this case with default header values.

The VPN connection was created successfully.
*/
type PostVpnGatewaysVpnGatewayIDConnectionsCreated struct {
	Payload *models.VPNGatewayConnection
}

func (o *PostVpnGatewaysVpnGatewayIDConnectionsCreated) Error() string {
	return fmt.Sprintf("[POST /vpn_gateways/{vpn_gateway_id}/connections][%d] postVpnGatewaysVpnGatewayIdConnectionsCreated  %+v", 201, o.Payload)
}

func (o *PostVpnGatewaysVpnGatewayIDConnectionsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VPNGatewayConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVpnGatewaysVpnGatewayIDConnectionsBadRequest creates a PostVpnGatewaysVpnGatewayIDConnectionsBadRequest with default headers values
func NewPostVpnGatewaysVpnGatewayIDConnectionsBadRequest() *PostVpnGatewaysVpnGatewayIDConnectionsBadRequest {
	return &PostVpnGatewaysVpnGatewayIDConnectionsBadRequest{}
}

/*PostVpnGatewaysVpnGatewayIDConnectionsBadRequest handles this case with default header values.

An invalid VPN connection template was provided.
*/
type PostVpnGatewaysVpnGatewayIDConnectionsBadRequest struct {
	Payload *models.Riaaserror
}

func (o *PostVpnGatewaysVpnGatewayIDConnectionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /vpn_gateways/{vpn_gateway_id}/connections][%d] postVpnGatewaysVpnGatewayIdConnectionsBadRequest  %+v", 400, o.Payload)
}

func (o *PostVpnGatewaysVpnGatewayIDConnectionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
