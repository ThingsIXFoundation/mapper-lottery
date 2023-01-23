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
package draw_test

import (
	"crypto/rand"
	"math/big"
	"testing"

	. "github.com/ThingsIXFoundation/mapper-lottery/draw"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/go-cmp/cmp"
)

func drawResultComperer() cmp.Option {
	return cmp.Comparer(func(x, y DrawResult) bool {
		return x.Address == y.Address &&
			x.TicketNumber == y.TicketNumber &&
			x.DrawNumber.Cmp(y.DrawNumber) == 0
	})
}

func genParticipants(t *testing.T, n int) []*Participant {
	var (
		participants = make([]*Participant, n)
		ticketNumber = uint64(0)
	)
	for i := range participants {
		var addr common.Address
		if _, err := rand.Read(addr[:]); err != nil {
			t.Fatalf("unable to generate participants: %v", err)
		}
		nLots, err := rand.Int(rand.Reader, big.NewInt(5))
		if err != nil {
			t.Fatal(err)
		}

		participant := &Participant{
			Address: addr,
		}

		for i := 0; i < int(nLots.Int64()+1); i++ {
			ticketNumber++
			participant.TicketNumbers = append(participant.TicketNumbers, ticketNumber)
		}

		participants[i] = participant
	}
	return participants
}

func TestDrawReproducable(t *testing.T) {
	participants := genParticipants(t, 10)

	// generate a random value that is used for the lottery draw
	var r [32]byte
	if _, err := rand.Read(r[:]); err != nil {
		t.Fatalf("unable to generate random for lottery draw: %v", err)
	}
	secret := new(big.Int).SetBytes(r[:])
	results1, err := Draw(secret, participants)
	if err != nil {
		t.Fatal(err)
	}

	// repeat draw with same rand and ensure that results are the same
	results2, err := Draw(secret, participants)
	if err != nil {
		t.Fatal(err)
	}

	// ensure that both results are equal
	if diff := cmp.Diff(results1, results2, drawResultComperer()); diff != "" {
		t.Errorf("lottery draw not stable: %v", diff)
	}

	for _, p := range results1 {
		t.Logf("%s %d", p.Address, p.TicketNumber)
	}
}

func TestDrawRand(t *testing.T) {
	participants := genParticipants(t, 32)

	var r [32]byte
	if _, err := rand.Read(r[:]); err != nil {
		t.Fatalf("unable to generate random for lottery draw: %v", err)
	}

	secret := new(big.Int).SetBytes(r[:])
	results1, err := Draw(secret, participants)
	if err != nil {
		t.Fatal(err)
	}

	// draw again with different secret and ensure that results differ
	r[0] = r[0] + 1
	secret = new(big.Int).SetBytes(r[:])
	results2, err := Draw(secret, participants)
	if err != nil {
		t.Fatal(err)
	}

	if cmp.Equal(results1, results2, drawResultComperer()) {
		t.Fatal("lottery draw yields same results with different secrets")
	}
}
