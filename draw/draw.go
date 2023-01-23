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
package draw

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/binary"
	"math/big"
	"sort"

	"github.com/ethereum/go-ethereum/common"
)

type Participant struct {
	Address       common.Address
	TicketNumbers []uint64
}

type DrawResult struct {
	Address      common.Address
	TicketNumber uint64
	DrawNumber   *big.Int
}

type DrawResults []*DrawResult

// Draw orders the given list of participants deterministic for the given random.
// This first lottery.MappersAvailable in the DrawResults won.
func Draw(drawRandom *big.Int, participants []*Participant) (DrawResults, error) {
	var (
		draw    = make([]*big.Int, 0, len(participants))
		mapping = make(map[*big.Int]*DrawResult)
		result  DrawResults
	)

	// derive a number for each ticket and create a mapper from that
	// number -> Participant/Lot
	for _, p := range participants {
		for _, ticket := range p.TicketNumbers {
			// prepare buf with lotnumber || participant.address
			var buf [8 + 20]byte
			binary.BigEndian.PutUint64(buf[:], ticket)
			copy(buf[8:], p.Address[:])

			// determine a draw number for the ticket that is derived from
			// buf with the lottery draw random value
			mac := hmac.New(sha256.New, drawRandom.Bytes())
			mac.Write(buf[:])
			drawNumber := new(big.Int).SetBytes(mac.Sum(nil))

			// mapping keys are ptrs, but that is fine in this case
			mapping[drawNumber] = &DrawResult{
				Address:      p.Address,
				TicketNumber: ticket,
				DrawNumber:   drawNumber,
			}
			draw = append(draw, drawNumber)
		}
	}

	// order tickets by their draw number and account.
	// this yields a sorted list with Participant/Lot combinations.
	sort.Slice(draw, func(i, j int) bool {
		switch draw[i].Cmp(draw[j]) {
		case -1:
			return true
		case 0:
			return bytes.Compare(
				mapping[draw[i]].Address.Bytes(),
				mapping[draw[j]].Address.Bytes()) < 0
		default:
			return false
		}
	})

	for _, num := range draw {
		result = append(result, mapping[num])
	}

	return result, nil
}
