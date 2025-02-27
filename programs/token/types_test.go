// Copyright 2020 dfuse Platform Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package token

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"

	"github.com/mr-tron/base58"
	bin "github.com/streamingfast/binary"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/teal-finance/solana-go"
	"github.com/teal-finance/solana-go/rpc"
)

func TestAccount(t *testing.T) {
	b58data := "SqtzmJArwV2556pK7AdHbHNPVP2L2WaR6zfcFeot94TzGRUyUMEWew558UxnYEGrmm9b9VZY7MS6TCHT5wqtzaA5Vy8ghoFyGmbRNC58CttRf5GzH9wfjCkncyrmKjfevyjrJ2W9XKLgYGth46ctFWzJJXCeHsYwDx1d"
	data, _ := base58.Decode(b58data)

	// fmt.Println("HEX:", hex.EncodeToString(data))
	// ba71eb12868584549b86f75620e7bb3ac5ef49df3fef0d48ad08e48dfa0fc786  // mint
	// d7a1d0a56e355f17cedd5733e36a0cc9e2caf7a435e3256e4c9bff755f682b5a  // owner
	// 5ece000000000000   // amount
	// 00000000    // is delegate set
	// 0000000000000000000000000000000000000000000000000000000000000000  // delegate
	// 01000000    // is initialized, is native + padding
	// 0000000000000000    // delegate amount
	var out Account
	err := bin.NewDecoder(data).Decode(&out)
	require.NoError(t, err)

	expect := Account{
		Mint:          solana.MustPublicKeyFromBase58("DYoajiN32pjK8zMAa67ScNn2E7EmXrZ6doABRqfSZ63F"),
		Owner:         solana.MustPublicKeyFromBase58("FWjmNcjufwC3QFdcHrAK1yAQkCwJSUAxvVFFgvQ1nAJM"),
		Amount:        bin.Uint64(52830),
		IsInitialized: true,
	}
	expectJSON, err := json.MarshalIndent(expect, "", "  ")
	require.NoError(t, err)

	outJSON, err := json.MarshalIndent(out, "", "  ")
	require.NoError(t, err)

	assert.JSONEq(t, string(expectJSON), string(outJSON))

	buf := &bytes.Buffer{}
	assert.NoError(t, bin.NewEncoder(buf).Encode(out))

	assert.Equal(t, b58data, base58.Encode(buf.Bytes()))
}

func TestMint(t *testing.T) {
	t.Skip("skipping network call")
	addr := solana.MustPublicKeyFromBase58("EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v")
	cli := rpc.NewClient("https://api.mainnet-beta.solana.com")

	var m Mint
	err := cli.GetAccountDataIn(addr, &m)
	// handle `err`
	require.NoError(t, err)

	json.NewEncoder(os.Stdout).Encode(m)
	// {"OwnerOption":1,
	//  "Owner":"2wmVCSfPxGPjrnMMn7rchp4uaeoTqN39mXFC2zhPdri9",
	//  "Decimals":128,
	//  "IsInitialized":true}
}

func TestRawMint(t *testing.T) {
	t.Skip("skipping network call")
	addr := solana.MustPublicKeyFromBase58("EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v")
	cli := rpc.NewClient("https://api.mainnet-beta.solana.com")

	resp, err := cli.GetAccountInfo(addr)
	// handle `err`
	require.NoError(t, err)

	json.NewEncoder(os.Stdout).Encode(resp)
	// {"OwnerOption":1,
	//  "Owner":"2wmVCSfPxGPjrnMMn7rchp4uaeoTqN39mXFC2zhPdri9",
	//  "Decimals":128,
	//  "IsInitialized":true}
}
