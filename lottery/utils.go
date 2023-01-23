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
	"context"
	"fmt"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ThingsIXFoundation/frequency-plan/go/frequency_plan"
	"github.com/ThingsIXFoundation/mapper-lottery/draw"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func createBindings(cmd *cobra.Command) (*LotteryCaller, *ethclient.Client, error) {
	var (
		lotteryAddressStr, _ = cmd.Flags().GetString(LotterySmartContractAddressFlag)
		lotteryAddress       = common.HexToAddress(lotteryAddressStr)
		rpcEndpoint, _       = cmd.Flags().GetString(RPCEndpointFlag)
		ctx, cancel          = context.WithTimeout(context.Background(), 30*time.Second)
	)
	defer cancel()

	client, err := ethclient.DialContext(ctx, rpcEndpoint)
	if err != nil {
		return nil, nil, err
	}

	lottery, err := NewLotteryCaller(lotteryAddress, client)
	if err != nil {
		return nil, nil, err
	}
	return lottery, client, nil
}

func printTickets(lottery IMapperLotteryByDrawLotteryDetails, participants []*draw.Participant) {
	if len(participants) == 0 {
		fmt.Println("no tickets sold")
		return
	}

	var (
		table          = tablewriter.NewWriter(os.Stdout)
		hasDrawResults = lottery.DrawSecret != nil && lottery.DrawSecret.BitLen() > 0
		results        draw.DrawResults
		err            error
	)

	if hasDrawResults {
		table.SetHeader([]string{"ticket num", "buyer", "draw num", "result"})
		results, err = draw.Draw(lottery.DrawSecret, participants)
		if err != nil {
			panic(err)
		}
	} else {
		table.SetHeader([]string{"ticket num", "buyer"})
	}

	var winningTickets []uint64

	if hasDrawResults {
		for i, r := range results {
			result := "lost"
			if uint64(i) < lottery.AvailableMappers {
				result = "won"
				winningTickets = append(winningTickets, r.TicketNumber)
			}
			table.Append([]string{fmt.Sprintf("%d", r.TicketNumber), r.Address.Hex(), fmt.Sprintf("0x%x", r.DrawNumber), result})
		}
	} else {
		for _, p := range participants {
			for _, t := range p.TicketNumbers {
				table.Append([]string{fmt.Sprintf("%d", t), p.Address.Hex()})
			}
		}
	}

	table.Render()
	if len(winningTickets) > 0 {
		fmt.Printf("winning ticket numbers: [%v]\n", strings.Trim(strings.Replace(fmt.Sprint(winningTickets), " ", ",", -1), "[]"))
	}
}

func printLotteries(client *ethclient.Client, lotteries []IMapperLotteryByDrawLotteryDetails) {
	if len(lotteries) == 0 {
		fmt.Println("no running or recent lotteries")
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"lottery", "status", "start", "end", "draw random", "ticket price", "mapper", "mappers available", "tickets sold", "token"})

	for _, l := range lotteries {
		table.Append([]string{
			l.Id.String(),
			statusToString(&l),
			time.Unix(l.StartTimestamp.Int64(), 0).String(),
			time.Unix(l.EndTimestamp.Int64(), 0).String(),
			drawValueToString(&l),
			ticketPrice(client, &l),
			mapperBandToString(&l),
			fmt.Sprintf("%d", l.AvailableMappers),
			fmt.Sprintf("%d", l.TicketsSold),
			l.Token.Hex(),
		})
	}

	table.Render()
}

func mapperBandToString(lottery *IMapperLotteryByDrawLotteryDetails) string {
	band := frequency_plan.BlockchainFrequencyPlan(lottery.MapperFrequencyPlan)
	return string(frequency_plan.FromBlockchain(band))
}

func ticketPrice(client *ethclient.Client, lottery *IMapperLotteryByDrawLotteryDetails) string {
	erc20, err := NewErc20Caller(lottery.Token, client)
	if err != nil {
		return "?"
	}
	decimals, err := erc20.Decimals(nil)
	if err != nil {
		return "?"
	}
	symbol, err := erc20.Symbol(nil)
	if err != nil {
		return "?"
	}

	var (
		exp    = new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), big.NewInt(0))
		before = new(big.Int).Div(lottery.TicketPrice, exp)
		after  = new(big.Int).Mod(lottery.TicketPrice, exp)
		last   = strings.TrimRight(after.String(), "0")
	)

	if last == "" {
		last = "0"
	}

	return before.String() + "." + last + " " + symbol
}

func drawValueToString(lottery *IMapperLotteryByDrawLotteryDetails) string {
	if lottery.DrawSecret != nil && lottery.DrawSecret.BitLen() != 0 {
		return fmt.Sprintf("0x%x", lottery.DrawSecret)
	}
	return ""
}

func statusToString(lottery *IMapperLotteryByDrawLotteryDetails) string {
	switch lottery.Status {
	case 0:
		start := time.Unix(lottery.StartTimestamp.Int64(), 0)
		if start.After(time.Now()) {
			return "pending"
		}
		end := time.Unix(lottery.EndTimestamp.Int64(), 0)
		if end.Before(time.Now()) {
			return "waiting for draw"
		}
		return "open"
	case 1:
		return "draw initiated"
	case 2:
		return "draw finished"
	case 3:
		return "finished"
	default:
		return "unknown"
	}
}
