// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lottery

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IMapperLotteryByDrawLotteryDetails is an auto generated low-level Go binding around an user-defined struct.
type IMapperLotteryByDrawLotteryDetails struct {
	Id                         *big.Int
	Status                     uint8
	Paused                     bool
	StartTimestamp             *big.Int
	EndTimestamp               *big.Int
	StartDrawSecretTimestamp   *big.Int
	AvailableMappers           uint64
	MapperFrequencyPlan        uint8
	TicketPrice                *big.Int
	Token                      common.Address
	TicketsSold                uint64
	TicketTokensReclaimingLeft uint64
	MaxTickets                 uint64
	DrawSecret                 *big.Int
	DrawFinishedBlock          *big.Int
}

// IMapperLotteryByDrawSoldTicket is an auto generated low-level Go binding around an user-defined struct.
type IMapperLotteryByDrawSoldTicket struct {
	Number  uint64
	Buyer   common.Address
	Won     bool
	Claimed bool
}

// LotteryMetaData contains all meta data concerning the Lottery contract.
var LotteryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AllTicketsSoldErr\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyBoughtTicketErr\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DrawInProgressErr\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DrawNotFinishedErr\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DrawResultsPublishedErr\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAvailableMappersErr\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNumberOfMappersToSellErr\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShippingProofErr\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTicketNumberErr\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTimeWindowErr\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LotteryInvalidStateErr\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LotteryIsPausedErr\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LotteryNotFinishedErr\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoTicketBoughtErr\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RequestDrawAlreadyInitiatedErr\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RequestDrawWindowNotStartedErr\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TicketBuyEndedErr\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TicketBuyNotStartedErr\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"ticketNumber\",\"type\":\"uint64\"}],\"name\":\"TicketTokensClaimedErr\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lotteryId\",\"type\":\"uint256\"}],\"name\":\"UnknownLotteryErr\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"ticketNumber\",\"type\":\"uint64\"}],\"name\":\"WonInLotteryErr\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"noContractAllowedErr\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"lotteryId\",\"type\":\"uint256\"}],\"name\":\"DrawInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"lotteryId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"secret\",\"type\":\"uint256\"}],\"name\":\"DrawRandomDetermined\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"lotteryId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"mappersAvailable\",\"type\":\"uint64\"}],\"name\":\"LotteryAvailableMappersChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"lotteryId\",\"type\":\"uint256\"}],\"name\":\"LotteryCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"lotteryId\",\"type\":\"uint256\"}],\"name\":\"LotteryFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"lotteryId\",\"type\":\"uint256\"}],\"name\":\"LotteryPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"lotteryId\",\"type\":\"uint256\"}],\"name\":\"LotteryUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"lotteryId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"ticketNumber\",\"type\":\"uint64\"}],\"name\":\"TicketSold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"lotteryId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"ticketNumer\",\"type\":\"uint64\"}],\"name\":\"TicketTokensClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"lotteryId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"ticketNumer\",\"type\":\"uint64\"}],\"name\":\"TicketTokensClaimedOnBehalf\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"lotteryId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"ticketNumber\",\"type\":\"uint64\"}],\"name\":\"WinningTicket\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lotteryId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"shippingProof\",\"type\":\"bytes\"}],\"name\":\"buyTicket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lotteryId\",\"type\":\"uint256\"}],\"name\":\"details\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumIMapperLotteryByDraw.LotteryStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"startTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startDrawSecretTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"availableMappers\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"mapperFrequencyPlan\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"ticketPrice\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ticketsSold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ticketTokensReclaimingLeft\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTickets\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"drawSecret\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"drawFinishedBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIMapperLotteryByDraw.LotteryDetails\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lotteryId\",\"type\":\"uint256\"}],\"name\":\"initiateDraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lotteriesCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lotteryId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"myTicket\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"number\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"won\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"claimed\",\"type\":\"bool\"}],\"internalType\":\"structIMapperLotteryByDraw.SoldTicket\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lotteryId\",\"type\":\"uint256\"}],\"name\":\"reclaimToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lotteryId\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"pageIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pageSize\",\"type\":\"uint64\"}],\"name\":\"soldTicketsPaged\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"number\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"won\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"claimed\",\"type\":\"bool\"}],\"internalType\":\"structIMapperLotteryByDraw.SoldTicket[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LotteryABI is the input ABI used to generate the binding from.
// Deprecated: Use LotteryMetaData.ABI instead.
var LotteryABI = LotteryMetaData.ABI

// Lottery is an auto generated Go binding around an Ethereum contract.
type Lottery struct {
	LotteryCaller     // Read-only binding to the contract
}

// LotteryCaller is an auto generated read-only Go binding around an Ethereum contract.
type LotteryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LotteryRaw is an auto generated low-level Go binding around an Ethereum contract.
type LotteryRaw struct {
	Contract *Lottery // Generic contract binding to access the raw methods on
}

// LotteryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LotteryCallerRaw struct {
	Contract *LotteryCaller // Generic read-only contract binding to access the raw methods on
}

// NewLotteryCaller creates a new read-only instance of Lottery, bound to a specific deployed contract.
func NewLotteryCaller(address common.Address, caller bind.ContractCaller) (*LotteryCaller, error) {
	contract, err := bindLottery(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LotteryCaller{contract: contract}, nil
}

// bindLottery binds a generic wrapper to an already deployed contract.
func bindLottery(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LotteryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lottery *LotteryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lottery.Contract.LotteryCaller.contract.Call(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lottery *LotteryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lottery.Contract.contract.Call(opts, result, method, params...)
}

// Details is a free data retrieval call binding the contract method 0xa005ec7a.
//
// Solidity: function details(uint256 lotteryId) view returns((uint256,uint8,bool,uint256,uint256,uint256,uint64,uint8,uint256,address,uint64,uint64,uint64,uint256,uint256))
func (_Lottery *LotteryCaller) Details(opts *bind.CallOpts, lotteryId *big.Int) (IMapperLotteryByDrawLotteryDetails, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "details", lotteryId)

	if err != nil {
		return *new(IMapperLotteryByDrawLotteryDetails), err
	}

	out0 := *abi.ConvertType(out[0], new(IMapperLotteryByDrawLotteryDetails)).(*IMapperLotteryByDrawLotteryDetails)

	return out0, err
}

// LotteriesCount is a free data retrieval call binding the contract method 0x50d9533e.
//
// Solidity: function lotteriesCount() view returns(uint256)
func (_Lottery *LotteryCaller) LotteriesCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "lotteriesCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MyTicket is a free data retrieval call binding the contract method 0x30b51f40.
//
// Solidity: function myTicket(uint256 lotteryId, address buyer) view returns((uint64,address,bool,bool))
func (_Lottery *LotteryCaller) MyTicket(opts *bind.CallOpts, lotteryId *big.Int, buyer common.Address) (IMapperLotteryByDrawSoldTicket, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "myTicket", lotteryId, buyer)

	if err != nil {
		return *new(IMapperLotteryByDrawSoldTicket), err
	}

	out0 := *abi.ConvertType(out[0], new(IMapperLotteryByDrawSoldTicket)).(*IMapperLotteryByDrawSoldTicket)

	return out0, err

}

// SoldTicketsPaged is a free data retrieval call binding the contract method 0xc354005b.
//
// Solidity: function soldTicketsPaged(uint256 lotteryId, uint64 pageIndex, uint64 pageSize) view returns((uint64,address,bool,bool)[])
func (_Lottery *LotteryCaller) SoldTicketsPaged(opts *bind.CallOpts, lotteryId *big.Int, pageIndex uint64, pageSize uint64) ([]IMapperLotteryByDrawSoldTicket, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "soldTicketsPaged", lotteryId, pageIndex, pageSize)

	if err != nil {
		return *new([]IMapperLotteryByDrawSoldTicket), err
	}

	out0 := *abi.ConvertType(out[0], new([]IMapperLotteryByDrawSoldTicket)).(*[]IMapperLotteryByDrawSoldTicket)

	return out0, err
}