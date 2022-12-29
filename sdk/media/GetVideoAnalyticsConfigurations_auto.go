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

// Call_GetVideoAnalyticsConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetVideoAnalyticsConfigurationsResponse.
func Call_GetVideoAnalyticsConfigurations(ctx context.Context, dev *onvif.Device, request media.GetVideoAnalyticsConfigurations) (media.GetVideoAnalyticsConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetVideoAnalyticsConfigurationsResponse media.GetVideoAnalyticsConfigurationsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetVideoAnalyticsConfigurationsResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetVideoAnalyticsConfigurations")
		return reply.Body.GetVideoAnalyticsConfigurationsResponse, errors.Annotate(err, "reply")
	}
}
