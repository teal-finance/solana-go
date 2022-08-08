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

package ws

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/teal-finance/solana-go"
	"go.uber.org/zap"
)

func Test_AccountSubscribe(t *testing.T) {
	t.Skip("Never ending test, revisit me to not depend on actual network calls, or hide between env flag")

	zlog, _ = zap.NewDevelopment()

	c := NewClient("ws://api.mainnet-beta.solana.com:80/rpc", false)
	err := c.Dial(context.Background())
	require.NoError(t, err)
	defer c.Close()

	accountID := solana.MustPublicKeyFromBase58("SqJP6vrvMad5XBQK5PCFEZjeuQSFi959sdpqtSNvnsX")
	sub, err := c.AccountSubscribe(accountID, "")
	require.NoError(t, err)

	data, err := sub.Recv(context.Background())
	if err != nil {
		fmt.Println("receive an error: ", err)
		return
	}

	fmt.Println("OpenOrders: ", data.(*AccountResult).Value.Account.Owner)
	fmt.Println("data: ", data.(*AccountResult).Value.Account.Data)
}

func Test_ProgramSubscribe(t *testing.T) {
	t.Skip("Never ending test, revisit me to not depend on actual network calls, or hide between env flag")

	zlog, _ = zap.NewDevelopment()

	c := NewClient("wss://solana-api.projectserum.com", false)
	err := c.Dial(context.Background())
	require.NoError(t, err)

	defer c.Close()

	programID := solana.MustPublicKeyFromBase58("EUqojwWA2rd19FZrzeBncJsm38Jm1hEhE3zsmX3bRc2o")
	sub, err := c.ProgramSubscribe(programID, "")
	require.NoError(t, err)

	for {
		data, err := sub.Recv(context.Background())
		if err != nil {
			fmt.Println("receive an error: ", err)
			return
		}
		fmt.Println("data received: ", data.(*ProgramResult).Value.PubKey)
	}
}

func Test_SlotSubscribe(t *testing.T) {
	t.Skip("Never ending test, revisit me to not depend on actual network calls, or hide between env flag")

	zlog, _ = zap.NewDevelopment()

	c := NewClient("ws://api.mainnet-beta.solana.com:80/rpc", false)
	err := c.Dial(context.Background())
	require.NoError(t, err)
	defer c.Close()

	sub, err := c.SlotSubscribe()
	require.NoError(t, err)

	data, err := sub.Recv(context.Background())
	if err != nil {
		fmt.Println("receive an error: ", err)
		return
	}
	fmt.Println("data received: ", data.(*SlotResult).Parent)
	return
}
