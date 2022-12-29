// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package media

import (
	"context"
	"github.com/DrewZt/onvif"
	"github.com/DrewZt/onvif/media"
	"github.com/DrewZt/onvif/sdk"
	"github.com/juju/errors"
)

// Call_CreateProfile forwards the call to dev.CallMethod() then parses the payload of the reply as a CreateProfileResponse.
func Call_CreateProfile(ctx context.Context, dev *onvif.Device, request media.CreateProfile) (media.CreateProfileResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			CreateProfileResponse media.CreateProfileResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.CreateProfileResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "CreateProfile")
		return reply.Body.CreateProfileResponse, errors.Annotate(err, "reply")
	}
}
