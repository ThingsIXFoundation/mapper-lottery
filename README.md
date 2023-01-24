# Mapper lottery
Utility to interact with the ThingsIX mapper lotteries contract.

# Build
Install golang from https://go.dev and build the software with:

```bash
$ git clone https://github.com/ThingsIXFoundation/mapper-lottery
$ cd mapper-lottery
$ go build
```

## Usage
the `mapper-lottery` has various subcommands that all require the user to supply
2 flags:

- `lottery-contract`, lottery smart contract address, can be found [here](https://docs.thingsix.com/background/smart-contracts).
- `rpc-endpoint`, RPC address of the Polygon node. There are various parties
offering free Polygon RPC endpoints. See the Polygon docs [here](https://wiki.polygon.technology/docs/develop/metamask/config-polygon-on-metamask#add-the-polygon-network-manually) for an example.

# Retrieve lotteries
Retrieve an overview of recently finished, pending or lotteries in progress with
the list subcommand.

```bash
$ ./mapper-lottery \
    --lottery-contract <lottery-smart-contract-address> \
    --rpc-endpoint <rpc-endpoint> \
    list

+---------+------------------+-------------------------------+-------------------------------+--------------------------------------------------------------------+--------------+--------+-------------------+--------------+--------------------------------------------+
| LOTTERY |      STATUS      |             START             |              END              |                            DRAW RANDOM                             | TICKET PRICE | MAPPER | MAPPERS AVAILABLE | TICKETS SOLD |                   TOKEN                    |
+---------+------------------+-------------------------------+-------------------------------+--------------------------------------------------------------------+--------------+--------+-------------------+--------------+--------------------------------------------+
|       1 | waiting for draw | 2023-01-19 12:50:26 +0100 CET | 2023-01-20 12:50:26 +0100 CET |                                                                    | 1.5 THIX     | EU868  |                 2 |            0 | 0x0a7cc20FE1E48663AAF319aCd476E64EC35ec97A |
|       2 | waiting for draw | 2023-01-20 09:21:25 +0100 CET | 2023-01-21 15:21:25 +0100 CET |                                                                    | 1.5 DERC20   | EU868  |                 2 |            0 | 0xfe4F5145f6e09952a5ba9e956ED0C25e3Fa4c7F1 |
|       3 | finished         | 2023-01-20 09:22:25 +0100 CET | 2023-01-21 15:22:25 +0100 CET | 0x117ca29309bd477e7fcab7ba26031e7c2d7113d1d36bb51210f4a91db8545a41 | 0.25 DERC20  | EU868  |                 2 |            3 | 0xfe4F5145f6e09952a5ba9e956ED0C25e3Fa4c7F1 |
+---------+------------------+-------------------------------+-------------------------------+--------------------------------------------------------------------+--------------+--------+-------------------+--------------+--------------------------------------------+
```

Status can have the following values:
- `pending`, lottery hasn't started
- `open`, lottery tickets can be bought
- `waiting for draw`, ticker buy period closed
- `draw initiated`, request for draw random initiated, waiting for random value
- `draw finished`, draw random value received, draw can be executed offline
- `finished`, winners are determined and loosers can claim their tokens back

# Retrieve tickets
Retrieve all sold tickets for lottery with:

```bash
$ ./mapper-lottery \
    --lottery-contract <lottery-smart-contract-address> \
    --rpc-endpoint <rpc-endpoint> \
    tickets <lottery-id> --verify

+------------+--------------------------------------------+--------------------------------------------------------------------+--------+
| TICKET NUM |                   BUYER                    |                              DRAW NUM                              | RESULT |
+------------+--------------------------------------------+--------------------------------------------------------------------+--------+
|          1 | 0x058d00Ed01fc0339B86eDfcaa1bDBD371AcDf4d5 | 0x1f5c56bcd6c1f9a4ea2bb5255f3f509dc0b6d66d1fffe49ef2fecb825121f8bc | won    |
|          3 | 0x782123189312Aa15c2C50A87F7Fe737DE38f3569 | 0x514e9938ac10c88058574e90f40e4de704fc19514737d73d9048468c5d5f9c3f | won    |
|          2 | 0xE10A9A4263eE02062f1248Ff79090cAF48176E01 | 0xe052f63e88037c26fd3d87cba49222e7982e21878efbac1334d757e00702c8a1 | lost   |
+------------+--------------------------------------------+--------------------------------------------------------------------+--------+
winning ticket numbers: [1,3]
```

This command accepts an optional verify flag. If the lottery results are
available it will try to verify for ~20% of the sold tickets if their results
in the lottery contract match with the locally calculated results.

This subcommand expectes the lottery id as input. The output depends on the
status of the lottery. If the lottery draw value is known it will include the
draw number and an indication if the ticket has won. Otherwise it will only 
print the ticker number and buyer address.