// Copyright 2023 Stichting ThingsIX Foundation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0
package lottery

import (
	"math/big"
	"time"

	"github.com/ThingsIXFoundation/mapper-lottery/draw"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

const (
	LotterySmartContractAddressFlag = "lottery-contract"
	RPCEndpointFlag                 = "rpc-endpoint"
)

func ListCmd(cmd *cobra.Command, args []string) {
	contract, client, err := createBindings(cmd)
	if err != nil {
		panic(err)
	}

	count, err := contract.LotteriesCount(nil)
	if err != nil {
		panic(err)
	}

	// ignore lotteries with an end time before cutoff
	var (
		cutoff    = time.Now().Add(-31 * 24 * time.Hour)
		lotteries []IMapperLotteryByDrawLotteryDetails
	)

	for i := int64(1); i <= count.Int64(); i++ {
		lottery, err := contract.Details(nil, big.NewInt(i))
		if err != nil {
			panic(err)
		}
		endTime := time.Unix(lottery.EndTimestamp.Int64(), 0)
		if endTime.After(cutoff) {
			lotteries = append(lotteries, lottery)
		}
	}

	printLotteries(client, lotteries)
}

func TicketsCmd(cmd *cobra.Command, args []string) {
	contract, _, err := createBindings(cmd)
	if err != nil {
		panic(err)
	}

	var (
		participants           = make(map[common.Address][]uint64)
		pageSize               = uint64(50)
		lotteryID, lotteryIDOk = new(big.Int).SetString(args[0], 0)
	)

	if !lotteryIDOk {
		panic("invalid lottery id")
	}

	lottery, err := contract.Details(nil, lotteryID)
	if err != nil {
		panic(err)
	}

	for page := uint64(0); ; page++ {
		tickets, err := contract.SoldTicketsPaged(nil, lotteryID, page, pageSize)
		if err != nil {
			panic(err)
		}
		if len(tickets) == 0 {
			break
		}
		for _, t := range tickets {
			participants[t.Buyer] = append(participants[t.Buyer], t.Number)
		}
		if len(tickets) != int(pageSize) {
			break
		}
	}

	// normalize partipants map to slice
	var p []*draw.Participant
	for b, t := range participants {
		p = append(p, &draw.Participant{
			Address:       b,
			TicketNumbers: t,
		})
	}

	printTickets(lottery, p)
}
