/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package providers

import (
	"magma/orc8r/lib/go/protos"

	"github.com/golang/protobuf/ptypes/any"
)

// StreamProvider provides a streamer policy. Given a gateway hardware ID,
// return a serialized data bundle of updates to stream back to the gateway.
type StreamProvider interface {
	// GetStreamName returns the name of the stream that this provider
	// services. This name must be globally unique.
	GetStreamName() string

	// GetUpdates returns updates to stream updates back to a gateway given its hardware ID
	// If GetUpdates returns error, the stream will be closed without sending any updates
	// If GetUpdates returns error == nil, updates will be sent & the stream will be closed after that
	// If GetUpdates returns error == io.EAGAIN - the returned updates will be sent & GetUpdates will be called again
	// on the same stream
	GetUpdates(gatewayId string, extraArgs *any.Any) ([]*protos.DataUpdate, error)
}
