// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common/math"
)

// constructErc20Data constructs the data field to be passed into a n erc20 deposit call
func ConstructErc20DepositData(destRecipient []byte, amount *big.Int) []byte {
	var data = []byte{}
	data = append(data, math.PaddedBigBytes(amount, 32)...)
	data = append(data, math.PaddedBigBytes(big.NewInt(int64(len(destRecipient))), 32)...)
	data = append(data, padBytes32(destRecipient)...)
	return data
}

// constructErc20Data constructs the data field to be passed into an erc721 deposit call
func ConstructErc721DepositData(tokenId *big.Int, destRecipient []byte) []byte {
	var data = []byte{}
	data = append(data, math.PaddedBigBytes(tokenId, 32)...)
	data = append(data, math.PaddedBigBytes(big.NewInt(int64(len(destRecipient))), 32)...)
	data = append(data, padBytes32(destRecipient)...)

	return data
}

func ConstructGenericDepositData(metadata []byte) []byte {
	var data = []byte{}
	data = append(data, math.PaddedBigBytes(big.NewInt(int64(len(metadata))), 32)...)
	data = append(data, metadata...)

	return data
}

func padBytes32(input []byte) []byte {
	var data = []byte{}
	data = append(data, input...)

	for (len(data) & 31) != 0 {
		data = append(data, byte(0))
	}

	return data
}
