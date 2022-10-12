// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BidTATList is an auto generated low-level Go binding around an user-defined struct.
type BidTATList struct {
	Bider  common.Address
	Amount *big.Int
}

// ContractABI is the input ABI used to generate the binding from.
const ContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"bidBurn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"bidList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_bider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_type\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_result\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_status\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"bidResult\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GetInitializeData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"TATBiders\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_TAT\",\"type\":\"address\"}],\"name\":\"bidInit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"bidTAT\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getList\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"bider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structBid.TATList[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pAddr\",\"type\":\"address\"}],\"name\":\"isTATBider\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reset\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// ContractBin is the compiled bytecode used for deploying new contracts.
var ContractBin = "0x6080604052600160685534801561001557600080fd5b50611d17806100256000396000f3fe6080604052600436106100a05760003560e01c8063942b765a11610064578063942b765a1461017f578063a18af9bb146101aa578063bd2ea1f3146101e7578063d3d655f814610224578063d826f88f1461024f578063f2fde38b1461027a576100a7565b80632b231df3146100ac5780634b954a1d146100e9578063715018a6146101265780638129fc1c1461013d5780638da5cb5b14610154576100a7565b366100a757005b600080fd5b3480156100b857600080fd5b506100d360048036038101906100ce919061132c565b6102a3565b6040516100e09190611682565b60405180910390f35b3480156100f557600080fd5b50610110600480360381019061010b919061132c565b610671565b60405161011d919061161c565b60405180910390f35b34801561013257600080fd5b5061013b6106b0565b005b34801561014957600080fd5b506101526106c4565b005b34801561016057600080fd5b50610169610802565b604051610176919061161c565b60405180910390f35b34801561018b57600080fd5b5061019461082c565b6040516101a19190611660565b60405180910390f35b3480156101b657600080fd5b506101d160048036038101906101cc91906112bf565b610ad9565b6040516101de9190611682565b60405180910390f35b3480156101f357600080fd5b5061020e60048036038101906102099190611292565b610c08565b60405161021b9190611682565b60405180910390f35b34801561023057600080fd5b50610239610c5e565b604051610246919061169d565b60405180910390f35b34801561025b57600080fd5b50610264610ce8565b6040516102719190611682565b60405180910390f35b34801561028657600080fd5b506102a1600480360381019061029c9190611292565b610f23565b005b60006001606854146102ea576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102e19061179a565b60405180910390fd5b670de0b6b3a7640000821015610335576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161032c906117e8565b60405180910390fd5b81606b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231336040518263ffffffff1660e01b8152600401610391919061161c565b60206040518083038186803b1580156103a957600080fd5b505afa1580156103bd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103e19190611359565b1015610422576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104199061173a565b60405180910390fd5b606b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ff2559bb33846040518363ffffffff1660e01b815260040161047f929190611637565b602060405180830381600087803b15801561049957600080fd5b505af11580156104ad573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104d191906112ff565b610510576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610507906116da565b60405180910390fd5b61051933610c08565b6105d9576001606660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506065339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b81606760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610628919061188e565b925050819055507fb515408192f1fae274bb8d2d170a6cff6eb5c3beba687b48212bb45de26f17403383604051610660929190611637565b60405180910390a160019050919050565b6065818154811061068157600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6106b8610fa7565b6106c26000611025565b565b60008060019054906101000a900460ff161590508080156106f55750600160008054906101000a900460ff1660ff16105b806107225750610704306110eb565b1580156107215750600160008054906101000a900460ff1660ff16145b5b610761576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107589061175a565b60405180910390fd5b60016000806101000a81548160ff021916908360ff160217905550801561079e576001600060016101000a81548160ff0219169083151502179055505b6107a661110e565b80156107ff5760008060016101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860016040516107f691906116bf565b60405180910390a15b50565b6000603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6060606a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663bf5bb758336040518263ffffffff1660e01b815260040161088991906117ba565b60206040518083038186803b1580156108a157600080fd5b505afa1580156108b5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108d991906112ff565b610918576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161090f9061171a565b60405180910390fd5b6002606881905550600060658054905067ffffffffffffffff81111561094157610940611a25565b5b60405190808252806020026020018201604052801561097a57816020015b6109676111d0565b81526020019060019003908161095f5790505b50905060005b606580549050811015610ad157606581815481106109a1576109a06119f6565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168282815181106109df576109de6119f6565b5b60200260200101516000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250506067600060658381548110610a3557610a346119f6565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054828281518110610aae57610aad6119f6565b5b602002602001015160200181815250508080610ac99061197e565b915050610980565b508091505090565b60008073ffffffffffffffffffffffffffffffffffffffff16606960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415610bfd5782606960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082606a60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081606b60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060019050610c02565b600080fd5b92915050565b6000606660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b60606040516024016040516020818303038152906040527f8129fc1c000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050905090565b6000606a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663bf5bb758336040518263ffffffff1660e01b8152600401610d4591906117ba565b60206040518083038186803b158015610d5d57600080fd5b505afa158015610d71573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d9591906112ff565b610dd4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dcb9061171a565b60405180910390fd5b600160688190555060005b606580549050811015610f0d576066600060658381548110610e0457610e036119f6565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549060ff02191690556067600060658381548110610e9157610e906119f6565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600090558080610f059061197e565b915050610ddf565b5060656000610f1c9190611200565b6001905090565b610f2b610fa7565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610f9b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f92906116fa565b60405180910390fd5b610fa481611025565b50565b610faf611167565b73ffffffffffffffffffffffffffffffffffffffff16610fcd610802565b73ffffffffffffffffffffffffffffffffffffffff1614611023576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161101a9061177a565b60405180910390fd5b565b6000603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081603360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600060019054906101000a900460ff1661115d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161115490611808565b60405180910390fd5b61116561116f565b565b600033905090565b600060019054906101000a900460ff166111be576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111b590611808565b60405180910390fd5b6111ce6111c9611167565b611025565b565b6040518060400160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b508054600082559060005260206000209081019061121e9190611221565b50565b5b8082111561123a576000816000905550600101611222565b5090565b60008135905061124d81611c9c565b92915050565b60008151905061126281611cb3565b92915050565b60008135905061127781611cca565b92915050565b60008151905061128c81611cca565b92915050565b6000602082840312156112a8576112a7611a54565b5b60006112b68482850161123e565b91505092915050565b600080604083850312156112d6576112d5611a54565b5b60006112e48582860161123e565b92505060206112f58582860161123e565b9150509250929050565b60006020828403121561131557611314611a54565b5b600061132384828501611253565b91505092915050565b60006020828403121561134257611341611a54565b5b600061135084828501611268565b91505092915050565b60006020828403121561136f5761136e611a54565b5b600061137d8482850161127d565b91505092915050565b600061139283836115cf565b60408301905092915050565b6113a7816118e4565b82525050565b6113b6816118e4565b82525050565b60006113c782611838565b6113d1818561185b565b93506113dc83611828565b8060005b8381101561140d5781516113f48882611386565b97506113ff8361184e565b9250506001810190506113e0565b5085935050505092915050565b611423816118f6565b82525050565b600061143482611843565b61143e818561186c565b935061144e81856020860161194b565b61145781611a59565b840191505092915050565b61146b81611939565b82525050565b600061147e600f8361187d565b915061148982611a6a565b602082019050919050565b60006114a160268361187d565b91506114ac82611a93565b604082019050919050565b60006114c4602d8361187d565b91506114cf82611ae2565b604082019050919050565b60006114e760198361187d565b91506114f282611b31565b602082019050919050565b600061150a602e8361187d565b915061151582611b5a565b604082019050919050565b600061152d60208361187d565b915061153882611ba9565b602082019050919050565b600061155060148361187d565b915061155b82611bd2565b602082019050919050565b600061157360118361187d565b915061157e82611bfb565b602082019050919050565b600061159660158361187d565b91506115a182611c24565b602082019050919050565b60006115b9602b8361187d565b91506115c482611c4d565b604082019050919050565b6040820160008201516115e5600085018261139e565b5060208201516115f860208501826115fe565b50505050565b61160781611922565b82525050565b61161681611922565b82525050565b600060208201905061163160008301846113ad565b92915050565b600060408201905061164c60008301856113ad565b611659602083018461160d565b9392505050565b6000602082019050818103600083015261167a81846113bc565b905092915050565b6000602082019050611697600083018461141a565b92915050565b600060208201905081810360008301526116b78184611429565b905092915050565b60006020820190506116d46000830184611462565b92915050565b600060208201905081810360008301526116f381611471565b9050919050565b6000602082019050818103600083015261171381611494565b9050919050565b60006020820190508181036000830152611733816114b7565b9050919050565b60006020820190508181036000830152611753816114da565b9050919050565b60006020820190508181036000830152611773816114fd565b9050919050565b6000602082019050818103600083015261179381611520565b9050919050565b600060208201905081810360008301526117b381611543565b9050919050565b600060408201905081810360008301526117d381611566565b90506117e260208301846113ad565b92915050565b6000602082019050818103600083015261180181611589565b9050919050565b60006020820190508181036000830152611821816115ac565b9050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600061189982611922565b91506118a483611922565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156118d9576118d86119c7565b5b828201905092915050565b60006118ef82611902565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60006119448261192c565b9050919050565b60005b8381101561196957808201518184015260208101905061194e565b83811115611978576000848401525b50505050565b600061198982611922565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214156119bc576119bb6119c7565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b6000601f19601f8301169050919050565b7f62696420636f7374206661696c65640000000000000000000000000000000000600082015250565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f4f6e6c7920466f756e646174696f6e4d616e616765722063616e20757365207460008201527f6869732066756e6374696f6e2e00000000000000000000000000000000000000602082015250565b7f5441542062616c616e6365206973206e6f7420656e6f75676800000000000000600082015250565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b7f61756469746f72206973206f7065726174696f6e000000000000000000000000600082015250565b7f466f756e646174696f6e4d616e61676572000000000000000000000000000000600082015250565b7f544154206e6f7420656e6f75676820746f206269640000000000000000000000600082015250565b7f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960008201527f6e697469616c697a696e67000000000000000000000000000000000000000000602082015250565b611ca5816118e4565b8114611cb057600080fd5b50565b611cbc816118f6565b8114611cc757600080fd5b50565b611cd381611922565b8114611cde57600080fd5b5056fea264697066735822122050dd223457b7e02c592443dd0326bca8f9abf6d8ff3bd0ae4e62919aa07fd7ad64736f6c63430008070033"

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// GetInitializeData is a free data retrieval call binding the contract method 0xd3d655f8.
//
// Solidity: function GetInitializeData() pure returns(bytes)
func (_Contract *ContractCaller) GetInitializeData(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "GetInitializeData")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetInitializeData is a free data retrieval call binding the contract method 0xd3d655f8.
//
// Solidity: function GetInitializeData() pure returns(bytes)
func (_Contract *ContractSession) GetInitializeData() ([]byte, error) {
	return _Contract.Contract.GetInitializeData(&_Contract.CallOpts)
}

// GetInitializeData is a free data retrieval call binding the contract method 0xd3d655f8.
//
// Solidity: function GetInitializeData() pure returns(bytes)
func (_Contract *ContractCallerSession) GetInitializeData() ([]byte, error) {
	return _Contract.Contract.GetInitializeData(&_Contract.CallOpts)
}

// TATBiders is a free data retrieval call binding the contract method 0x4b954a1d.
//
// Solidity: function TATBiders(uint256 ) view returns(address)
func (_Contract *ContractCaller) TATBiders(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "TATBiders", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TATBiders is a free data retrieval call binding the contract method 0x4b954a1d.
//
// Solidity: function TATBiders(uint256 ) view returns(address)
func (_Contract *ContractSession) TATBiders(arg0 *big.Int) (common.Address, error) {
	return _Contract.Contract.TATBiders(&_Contract.CallOpts, arg0)
}

// TATBiders is a free data retrieval call binding the contract method 0x4b954a1d.
//
// Solidity: function TATBiders(uint256 ) view returns(address)
func (_Contract *ContractCallerSession) TATBiders(arg0 *big.Int) (common.Address, error) {
	return _Contract.Contract.TATBiders(&_Contract.CallOpts, arg0)
}

// IsTATBider is a free data retrieval call binding the contract method 0xbd2ea1f3.
//
// Solidity: function isTATBider(address pAddr) view returns(bool result)
func (_Contract *ContractCaller) IsTATBider(opts *bind.CallOpts, pAddr common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isTATBider", pAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTATBider is a free data retrieval call binding the contract method 0xbd2ea1f3.
//
// Solidity: function isTATBider(address pAddr) view returns(bool result)
func (_Contract *ContractSession) IsTATBider(pAddr common.Address) (bool, error) {
	return _Contract.Contract.IsTATBider(&_Contract.CallOpts, pAddr)
}

// IsTATBider is a free data retrieval call binding the contract method 0xbd2ea1f3.
//
// Solidity: function isTATBider(address pAddr) view returns(bool result)
func (_Contract *ContractCallerSession) IsTATBider(pAddr common.Address) (bool, error) {
	return _Contract.Contract.IsTATBider(&_Contract.CallOpts, pAddr)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// BidInit is a paid mutator transaction binding the contract method 0xa18af9bb.
//
// Solidity: function bidInit(address _governance, address _TAT) returns(bool)
func (_Contract *ContractTransactor) BidInit(opts *bind.TransactOpts, _governance common.Address, _TAT common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "bidInit", _governance, _TAT)
}

// BidInit is a paid mutator transaction binding the contract method 0xa18af9bb.
//
// Solidity: function bidInit(address _governance, address _TAT) returns(bool)
func (_Contract *ContractSession) BidInit(_governance common.Address, _TAT common.Address) (*types.Transaction, error) {
	return _Contract.Contract.BidInit(&_Contract.TransactOpts, _governance, _TAT)
}

// BidInit is a paid mutator transaction binding the contract method 0xa18af9bb.
//
// Solidity: function bidInit(address _governance, address _TAT) returns(bool)
func (_Contract *ContractTransactorSession) BidInit(_governance common.Address, _TAT common.Address) (*types.Transaction, error) {
	return _Contract.Contract.BidInit(&_Contract.TransactOpts, _governance, _TAT)
}

// BidTAT is a paid mutator transaction binding the contract method 0x2b231df3.
//
// Solidity: function bidTAT(uint256 amount) returns(bool result)
func (_Contract *ContractTransactor) BidTAT(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "bidTAT", amount)
}

// BidTAT is a paid mutator transaction binding the contract method 0x2b231df3.
//
// Solidity: function bidTAT(uint256 amount) returns(bool result)
func (_Contract *ContractSession) BidTAT(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.BidTAT(&_Contract.TransactOpts, amount)
}

// BidTAT is a paid mutator transaction binding the contract method 0x2b231df3.
//
// Solidity: function bidTAT(uint256 amount) returns(bool result)
func (_Contract *ContractTransactorSession) BidTAT(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.BidTAT(&_Contract.TransactOpts, amount)
}

// GetList is a paid mutator transaction binding the contract method 0x942b765a.
//
// Solidity: function getList() returns((address,uint256)[])
func (_Contract *ContractTransactor) GetList(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "getList")
}

// GetList is a paid mutator transaction binding the contract method 0x942b765a.
//
// Solidity: function getList() returns((address,uint256)[])
func (_Contract *ContractSession) GetList() (*types.Transaction, error) {
	return _Contract.Contract.GetList(&_Contract.TransactOpts)
}

// GetList is a paid mutator transaction binding the contract method 0x942b765a.
//
// Solidity: function getList() returns((address,uint256)[])
func (_Contract *ContractTransactorSession) GetList() (*types.Transaction, error) {
	return _Contract.Contract.GetList(&_Contract.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Contract *ContractTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Contract *ContractSession) Initialize() (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Contract *ContractTransactorSession) Initialize() (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// Reset is a paid mutator transaction binding the contract method 0xd826f88f.
//
// Solidity: function reset() returns(bool)
func (_Contract *ContractTransactor) Reset(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "reset")
}

// Reset is a paid mutator transaction binding the contract method 0xd826f88f.
//
// Solidity: function reset() returns(bool)
func (_Contract *ContractSession) Reset() (*types.Transaction, error) {
	return _Contract.Contract.Reset(&_Contract.TransactOpts)
}

// Reset is a paid mutator transaction binding the contract method 0xd826f88f.
//
// Solidity: function reset() returns(bool)
func (_Contract *ContractTransactorSession) Reset() (*types.Transaction, error) {
	return _Contract.Contract.Reset(&_Contract.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contract *ContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contract *ContractSession) Receive() (*types.Transaction, error) {
	return _Contract.Contract.Receive(&_Contract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contract *ContractTransactorSession) Receive() (*types.Transaction, error) {
	return _Contract.Contract.Receive(&_Contract.TransactOpts)
}

// ContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Contract contract.
type ContractInitializedIterator struct {
	Event *ContractInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractInitialized represents a Initialized event raised by the Contract contract.
type ContractInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Contract *ContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractInitializedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractInitializedIterator{contract: _Contract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Contract *ContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractInitialized) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractInitialized)
				if err := _Contract.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Contract *ContractFilterer) ParseInitialized(log types.Log) (*ContractInitialized, error) {
	event := new(ContractInitialized)
	if err := _Contract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contract contract.
type ContractOwnershipTransferredIterator struct {
	Event *ContractOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOwnershipTransferred represents a OwnershipTransferred event raised by the Contract contract.
type ContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOwnershipTransferredIterator{contract: _Contract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOwnershipTransferred)
				if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOwnershipTransferred, error) {
	event := new(ContractOwnershipTransferred)
	if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractBidBurnIterator is returned from FilterBidBurn and is used to iterate over the raw logs and unpacked data for BidBurn events raised by the Contract contract.
type ContractBidBurnIterator struct {
	Event *ContractBidBurn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractBidBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBidBurn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractBidBurn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractBidBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractBidBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractBidBurn represents a BidBurn event raised by the Contract contract.
type ContractBidBurn struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBidBurn is a free log retrieval operation binding the contract event 0xb515408192f1fae274bb8d2d170a6cff6eb5c3beba687b48212bb45de26f1740.
//
// Solidity: event bidBurn(address account, uint256 amount)
func (_Contract *ContractFilterer) FilterBidBurn(opts *bind.FilterOpts) (*ContractBidBurnIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "bidBurn")
	if err != nil {
		return nil, err
	}
	return &ContractBidBurnIterator{contract: _Contract.contract, event: "bidBurn", logs: logs, sub: sub}, nil
}

// WatchBidBurn is a free log subscription operation binding the contract event 0xb515408192f1fae274bb8d2d170a6cff6eb5c3beba687b48212bb45de26f1740.
//
// Solidity: event bidBurn(address account, uint256 amount)
func (_Contract *ContractFilterer) WatchBidBurn(opts *bind.WatchOpts, sink chan<- *ContractBidBurn) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "bidBurn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractBidBurn)
				if err := _Contract.contract.UnpackLog(event, "bidBurn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBidBurn is a log parse operation binding the contract event 0xb515408192f1fae274bb8d2d170a6cff6eb5c3beba687b48212bb45de26f1740.
//
// Solidity: event bidBurn(address account, uint256 amount)
func (_Contract *ContractFilterer) ParseBidBurn(log types.Log) (*ContractBidBurn, error) {
	event := new(ContractBidBurn)
	if err := _Contract.contract.UnpackLog(event, "bidBurn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractBidListIterator is returned from FilterBidList and is used to iterate over the raw logs and unpacked data for BidList events raised by the Contract contract.
type ContractBidListIterator struct {
	Event *ContractBidList // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractBidListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBidList)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractBidList)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractBidListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractBidListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractBidList represents a BidList event raised by the Contract contract.
type ContractBidList struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBidList is a free log retrieval operation binding the contract event 0x55d50d3afe11324fa885d8ec1031ac66ca67531d3d9ae6aeaee3a1e25ef24823.
//
// Solidity: event bidList(address account, uint256 amount)
func (_Contract *ContractFilterer) FilterBidList(opts *bind.FilterOpts) (*ContractBidListIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "bidList")
	if err != nil {
		return nil, err
	}
	return &ContractBidListIterator{contract: _Contract.contract, event: "bidList", logs: logs, sub: sub}, nil
}

// WatchBidList is a free log subscription operation binding the contract event 0x55d50d3afe11324fa885d8ec1031ac66ca67531d3d9ae6aeaee3a1e25ef24823.
//
// Solidity: event bidList(address account, uint256 amount)
func (_Contract *ContractFilterer) WatchBidList(opts *bind.WatchOpts, sink chan<- *ContractBidList) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "bidList")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractBidList)
				if err := _Contract.contract.UnpackLog(event, "bidList", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBidList is a log parse operation binding the contract event 0x55d50d3afe11324fa885d8ec1031ac66ca67531d3d9ae6aeaee3a1e25ef24823.
//
// Solidity: event bidList(address account, uint256 amount)
func (_Contract *ContractFilterer) ParseBidList(log types.Log) (*ContractBidList, error) {
	event := new(ContractBidList)
	if err := _Contract.contract.UnpackLog(event, "bidList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractBidResultIterator is returned from FilterBidResult and is used to iterate over the raw logs and unpacked data for BidResult events raised by the Contract contract.
type ContractBidResultIterator struct {
	Event *ContractBidResult // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractBidResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBidResult)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractBidResult)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractBidResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractBidResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractBidResult represents a BidResult event raised by the Contract contract.
type ContractBidResult struct {
	Bider  common.Address
	Type   string
	Result string
	Status *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBidResult is a free log retrieval operation binding the contract event 0xc3a61fe96123abb10b574ab10422029d2e04745f9c87315946801a27f36609fe.
//
// Solidity: event bidResult(address _bider, string _type, string _result, uint256 _status, uint256 _amount)
func (_Contract *ContractFilterer) FilterBidResult(opts *bind.FilterOpts) (*ContractBidResultIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "bidResult")
	if err != nil {
		return nil, err
	}
	return &ContractBidResultIterator{contract: _Contract.contract, event: "bidResult", logs: logs, sub: sub}, nil
}

// WatchBidResult is a free log subscription operation binding the contract event 0xc3a61fe96123abb10b574ab10422029d2e04745f9c87315946801a27f36609fe.
//
// Solidity: event bidResult(address _bider, string _type, string _result, uint256 _status, uint256 _amount)
func (_Contract *ContractFilterer) WatchBidResult(opts *bind.WatchOpts, sink chan<- *ContractBidResult) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "bidResult")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractBidResult)
				if err := _Contract.contract.UnpackLog(event, "bidResult", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBidResult is a log parse operation binding the contract event 0xc3a61fe96123abb10b574ab10422029d2e04745f9c87315946801a27f36609fe.
//
// Solidity: event bidResult(address _bider, string _type, string _result, uint256 _status, uint256 _amount)
func (_Contract *ContractFilterer) ParseBidResult(log types.Log) (*ContractBidResult, error) {
	event := new(ContractBidResult)
	if err := _Contract.contract.UnpackLog(event, "bidResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
