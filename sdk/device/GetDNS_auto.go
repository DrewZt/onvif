// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"
	"github.com/DrewZt/onvif"
	"github.com/DrewZt/onvif/device"
	"github.com/DrewZt/onvif/sdk"
	"github.com/juju/errors"
)

// Call_GetDNS forwards the call to dev.CallMethod() then parses the payload of the reply as a GetDNSResponse.
func Call_GetDNS(ctx context.Context, dev *onvif.Device, request device.GetDNS) (device.GetDNSResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetDNSResponse device.GetDNSResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetDNSResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetDNS")
		return reply.Body.GetDNSResponse, errors.Annotate(err, "reply")
	}
}
