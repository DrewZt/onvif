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

// Call_SetRemoteDiscoveryMode forwards the call to dev.CallMethod() then parses the payload of the reply as a SetRemoteDiscoveryModeResponse.
func Call_SetRemoteDiscoveryMode(ctx context.Context, dev *onvif.Device, request device.SetRemoteDiscoveryMode) (device.SetRemoteDiscoveryModeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetRemoteDiscoveryModeResponse device.SetRemoteDiscoveryModeResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetRemoteDiscoveryModeResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetRemoteDiscoveryMode")
		return reply.Body.SetRemoteDiscoveryModeResponse, errors.Annotate(err, "reply")
	}
}
