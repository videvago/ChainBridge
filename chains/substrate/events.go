// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
	"fmt"

	events "github.com/ChainSafe/chainbridge-substrate-events"
	"github.com/ChainSafe/chainbridge-utils/msg"
	"github.com/ChainSafe/log15"
)

type eventName string
type eventHandler func(interface{}, log15.Logger) (msg.Message, error)

const GenericTransfer eventName = "GenericTransfer"

var Subscriptions = []struct {
	name    eventName
	handler eventHandler
}{
	{GenericTransfer, genericTransferHandler},
}

func genericTransferHandler(evtI interface{}, log log15.Logger) (msg.Message, error) {
	evt, ok := evtI.(events.EventGenericTransfer)
	if !ok {
		return msg.Message{}, fmt.Errorf("failed to cast EventGenericTransfer type")
	}

	log.Info("Got generic transfer event!", "destination", evt.Destination, "resourceId", evt.ResourceId)

	return msg.NewGenericTransfer(
		0, // Unset
		msg.ChainId(evt.Destination),
		msg.Nonce(evt.DepositNonce),
		msg.Bytes32(evt.DepositKey),
		msg.Bytes32(evt.ResourceId),
		evt.Metadata,
	), nil
}
