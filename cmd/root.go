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
package cmd

import (
	"os"

	"github.com/ThingsIXFoundation/mapper-lottery/lottery"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mapper-lottery",
	Short: "ThingsIX mapper lottery cli",
	Long:  `Command line utility to interact with the ThingsIX lottery smart contract.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().String(lottery.LotterySmartContractAddressFlag, "", "ThingsIX lottery smart contrat address")
	rootCmd.PersistentFlags().String(lottery.RPCEndpointFlag, "", "Blockchain node RPC endpoint")
}
