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

// Call_SetIPAddressFilter forwards the call to dev.CallMethod() then parses the payload of the reply as a SetIPAddressFilterResponse.
func Call_SetIPAddressFilter(ctx context.Context, dev *onvif.Device, request device.SetIPAddressFilter) (device.SetIPAddressFilterResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetIPAddressFilterResponse device.SetIPAddressFilterResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetIPAddressFilterResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetIPAddressFilter")
		return reply.Body.SetIPAddressFilterResponse, errors.Annotate(err, "reply")
	}
}
