// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package e2e

import (
	"fmt"
	"testing"

	eth "github.com/ChainSafe/ChainBridge/e2e/ethereum"
	sub "github.com/ChainSafe/ChainBridge/e2e/substrate"
	ethtest "github.com/ChainSafe/ChainBridge/shared/ethereum/testing"
	subtest "github.com/ChainSafe/ChainBridge/shared/substrate/testing"
	msg "github.com/ChainSafe/chainbridge-utils/msg"
	log "github.com/ChainSafe/log15"
)

func testSubstrateHashToGenericHandler(t *testing.T, ctx *testContext) {
	numberOfTxs := 5
	//nonce := subtest.GetDepositNonce(t, ctx.subClient) + 1
	for i := 1; i <= numberOfTxs; i++ {
		i := i // for scope
		ok := t.Run(fmt.Sprintf("Transfer %d", i), func(t *testing.T) {
			// Execute transfer
			hash := sub.HashInt(i)
			depositKey := msg.Bytes32(subtest.InitiateHashTransfer(t, ctx.subClient, hash, EthAChainId))

			// Wait for event
			eth.WaitForProposalActive(t, ctx.ethA.Client, ctx.ethA.BaseContracts.BridgeAddress, depositKey)
			eth.WaitForProposalExecutedEvent(t, ctx.ethA.Client, ctx.ethA.BaseContracts.BridgeAddress, depositKey)
			//nonce++

			// Verify hash is available
			ethtest.AssertHashExistence(t, ctx.ethA.Client, hash, ctx.ethA.TestContracts.AssetStoreSub)
		})
		if !ok {
			return
		}
	}
}

func testEthereumHashToGenericHandler(t *testing.T, ctx *testContext) {
	numberOfTxs := 5
	//nonce := ethtest.GetDepositNonce(t, ctx.ethA.Client, ctx.ethA.BaseContracts.BridgeAddress) + 1
	for i := 1; i <= numberOfTxs; i++ {
		i := i // for scope
		ok := t.Run(fmt.Sprintf("Transfer %d", i), func(t *testing.T) {
			// Execute transfer
			hash := sub.HashInt(i)
			log.Info("Submitting transaction", "number", i, "hash", hash.Hex(), "resourceId", ctx.EthGenericResourceId.Hex(), "from", ctx.ethA.Client.Opts.From, "handler", ctx.ethA.BaseContracts.GenericHandlerAddress)
			depositKey := eth.CreateGenericDeposit(t, ctx.ethA.Client, EthBChainId, eth.AliceKp.CommonAddress().Bytes(), hash[:], ctx.ethB.BaseContracts, ctx.EthGenericResourceId)

			// Wait for event
			eth.WaitForProposalActive(t, ctx.ethB.Client, ctx.ethB.BaseContracts.BridgeAddress, depositKey)
			eth.WaitForProposalExecutedEvent(t, ctx.ethB.Client, ctx.ethB.BaseContracts.BridgeAddress, depositKey)
			//nonce++

			// Verify hash is available
			ethtest.AssertHashExistence(t, ctx.ethB.Client, hash, ctx.ethB.TestContracts.AssetStoreEth)
		})
		if !ok {
			return
		}
	}
}
