// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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
)

// DoSCMetaData contains all meta data concerning the DoSC contract.
var DoSCMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[{\"name\":\"pkcoix\",\"type\":\"uint256\"},{\"name\":\"pkcoiy\",\"type\":\"uint256\"},{\"name\":\"hclj\",\"type\":\"bytes32\"},{\"name\":\"dt\",\"type\":\"uint256\"}],\"name\":\"Store\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"aaa\",\"type\":\"uint256\"}],\"name\":\"printLenOfStore\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"comx\",\"type\":\"uint256\"},{\"name\":\"comy\",\"type\":\"uint256\"},{\"name\":\"tokx\",\"type\":\"uint256\"},{\"name\":\"toky\",\"type\":\"uint256\"},{\"name\":\"hres\",\"type\":\"bytes32\"},{\"name\":\"hcl\",\"type\":\"bytes32\"},{\"name\":\"dt\",\"type\":\"uint256\"}],\"name\":\"Donate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pkcoix\",\"type\":\"uint256\"},{\"name\":\"pkcoiy\",\"type\":\"uint256\"},{\"name\":\"hclj\",\"type\":\"bytes32\"},{\"name\":\"dp\",\"type\":\"uint256\"},{\"name\":\"dt\",\"type\":\"uint256\"}],\"name\":\"Deliver2\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pkcoix\",\"type\":\"uint256\"},{\"name\":\"pkcoiy\",\"type\":\"uint256\"},{\"name\":\"dt\",\"type\":\"uint256\"},{\"name\":\"hcl\",\"type\":\"bytes32\"}],\"name\":\"Refund\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hfbi\",\"type\":\"bytes32\"},{\"name\":\"dt\",\"type\":\"uint256\"},{\"name\":\"hcl\",\"type\":\"bytes32\"}],\"name\":\"Receive\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"len1\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pkcoix\",\"type\":\"uint256\"},{\"name\":\"pkcoiy\",\"type\":\"uint256\"},{\"name\":\"hclj\",\"type\":\"bytes32\"},{\"name\":\"cmqx\",\"type\":\"uint256\"},{\"name\":\"cmqy\",\"type\":\"uint256\"},{\"name\":\"tokqx\",\"type\":\"uint256\"},{\"name\":\"tokqy\",\"type\":\"uint256\"}],\"name\":\"Distribute\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pkcoix\",\"type\":\"uint256\"},{\"name\":\"pkcoiy\",\"type\":\"uint256\"},{\"name\":\"hclj\",\"type\":\"bytes32\"},{\"name\":\"dp\",\"type\":\"uint256\"}],\"name\":\"Ship\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"txx\",\"type\":\"string\"},{\"name\":\"hcl\",\"type\":\"bytes32\"},{\"name\":\"dp\",\"type\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\"}],\"name\":\"Transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"a\",\"type\":\"string\"},{\"name\":\"b\",\"type\":\"string\"}],\"name\":\"compareStrings\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"len0\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"},{\"name\":\"k\",\"type\":\"uint256\"}],\"name\":\"test_add_dp\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"aaa\",\"type\":\"uint256\"}],\"name\":\"printLenOfDis\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pkcoix\",\"type\":\"uint256\"},{\"name\":\"pkcoiy\",\"type\":\"uint256\"},{\"name\":\"hclj\",\"type\":\"bytes32\"},{\"name\":\"dp\",\"type\":\"uint256\"},{\"name\":\"dt\",\"type\":\"uint256\"}],\"name\":\"Deliver1\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6060604052600680549050600755600980549050600a55341561002157600080fd5b612597806100306000396000f3006060604052600436106100db576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063073f21a9146100e05780630862c570146101225780631463f015146101595780632c5ecd73146101ba57806347b788f214610205578063746cac8d1461024757806379d3202e1461028457806384d298ff146102ad57806395e481451461030a5780639875a1e21461034c578063bed34bba146103de578063c3a97b8414610496578063ce09d470146104bf578063cf93815714610501578063e5aa1c3614610538575b600080fd5b34156100eb57600080fd5b610120600480803590602001909190803590602001909190803560001916906020019091908035906020019091905050610583565b005b341561012d57600080fd5b6101436004808035906020019091905050610657565b6040518082815260200191505060405180910390f35b341561016457600080fd5b6101b86004808035906020019091908035906020019091908035906020019091908035906020019091908035600019169060200190919080356000191690602001909190803590602001909190505061074b565b005b34156101c557600080fd5b61020360048080359060200190919080359060200190919080356000191690602001909190803590602001909190803590602001909190505061085c565b005b341561021057600080fd5b61024560048080359060200190919080359060200190919080359060200190919080356000191690602001909190505061099c565b005b341561025257600080fd5b61028260048080356000191690602001909190803590602001909190803560001916906020019091905050610d3d565b005b341561028f57600080fd5b610297610dc8565b6040518082815260200191505060405180910390f35b34156102b857600080fd5b610308600480803590602001909190803590602001909190803560001916906020019091908035906020019091908035906020019091908035906020019091908035906020019091905050610dce565b005b341561031557600080fd5b61034a600480803590602001909190803590602001909190803560001916906020019091908035906020019091905050610ef8565b005b341561035757600080fd5b6103dc600480803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919080356000191690602001909190803590602001909190803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050611027565b005b34156103e957600080fd5b61047c600480803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091905050611b78565b604051808215151515815260200191505060405180910390f35b34156104a157600080fd5b6104a9611c51565b6040518082815260200191505060405180910390f35b34156104ca57600080fd5b6104ff600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091908035906020019091905050611c57565b005b341561050c57600080fd5b6105226004808035906020019091905050611c9f565b6040518082815260200191505060405180910390f35b341561054357600080fd5b610581600480803590602001909190803590602001909190803560001916906020019091908035906020019091908035906020019091905050611d9c565b005b61058b611edc565b608060405190810160405280868152602001858152602001846000191681526020018381525090506105f56040805190810160405280600581526020017f53746f726500000000000000000000000000000000000000000000000000000081525084600033611027565b600680548060010182816106099190611f08565b91600052602060002090600402016000839091909150600082015181600001556020820151816001015560408201518160020190600019169055606082015181600301555050505050505050565b6000806000806000610667611edc565b608060405190810160405280868152602001858152602001846000191681526020018381525090506106d16040805190810160405280600581526020017f53746f726500000000000000000000000000000000000000000000000000000081525084600033611027565b600680548060010182816106e59190611f08565b91600052602060002090600402016000839091909150600082015181600001556020820151816001015560408201518160020190600019169055606082015181600301555050506001965060068054905060078190555060075495505050505050919050565b610753611f3a565b60e06040519081016040528089815260200188815260200187815260200186815260200185600019168152602001846000191681526020018381525090506107d36040805190810160405280600681526020017f446f6e617465000000000000000000000000000000000000000000000000000081525084600033611027565b600480548060010182816107e79190611f7e565b91600052602060002090600702016000839091909150600082015181600001556020820151816001015560408201518160020155606082015181600301556080820151816004019060001916905560a0820151816005019060001916905560c082015181600601555050505050505050505050565b610864611fb0565b610873336402540be400611c57565b60a0604051908101604052808781526020018681526020018560001916815260200184815260200183815250905082600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254039250508190555061092f6040805190810160405280600881526020017f44656c6976657232000000000000000000000000000000000000000000000000815250858533611027565b600b80548060010182816109439190611fe3565b916000526020600020906005020160008390919091506000820151816000015560208201518160010155604082015181600201906000191690556060820151816003015560808201518160040155505050505050505050565b6109a4612015565b6109ac612048565b60008060008060008760001916600019168152602001908152602001600020606060405190810160405290816000820160009054906101000a900460ff1660058111156109f557fe5b6005811115610a0057fe5b81526020016001820154815260200160028201805480602002602001604051908101604052809291908181526020018280548015610a9357602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610a49575b50505050508152505094506080604051908101604052808a815260200189815260200188815260200187600019168152509350600d8054806001018281610ada9190612074565b9160005260206000209060040201600086909190915060008201518160000155602082015181600101556040820151816002015560608201518160030190600019169055505050600060016000886000191660001916815260200190815260200160002060009054906101000a900460ff1660ff16141515610d325760046005811115610b6357fe5b85600001516005811115610b7357fe5b1415610c9557600092505b60038360ff161015610c945784604001518360ff16815181101515610b9f57fe5b906020019060200201519150600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905080600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254019250508190555080600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825403925050819055508280600101935050610b7e565b5b600585600001906005811115610ca757fe5b90816005811115610cb457fe5b81525050600085602001818152505084600080886000191660001916815260200190815260200160002060008201518160000160006101000a81548160ff02191690836005811115610d0257fe5b0217905550602082015181600101556040820151816002019080519060200190610d2d9291906120a6565b509050505b505050505050505050565b610d45612130565b6060604051908101604052808560001916815260200184815260200183600019168152509050600c8054806001018281610d7f9190612158565b9160005260206000209060030201600083909190915060008201518160000190600019169055602082015181600101556040820151816002019060001916905550505050505050565b600a5481565b6000610dd861218a565b429150610100604051908101604052808a815260200189815260200188600019168152602001878152602001868152602001858152602001848152602001838152509050610e5e6040805190810160405280600a81526020017f446973747269627574650000000000000000000000000000000000000000000081525088600033611027565b60098054806001018281610e7291906121d3565b91600052602060002090600802016000839091909150600082015181600001556020820151816001015560408201518160020190600019169055606082015181600301556080820151816004015560a0820151816005015560c0820151816006015560e08201518160070155505050600980549050600a81905550505050505050505050565b610f00612205565b610f0f336402540be400611c57565b6080604051908101604052808681526020018581526020018460001916815260200183815250905081600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282540392505081905550610fc56040805190810160405280600581526020017f53746f7265000000000000000000000000000000000000000000000000000000815250848433611027565b60088054806001018281610fd99190612231565b91600052602060002090600402016000839091909150600082015181600001556020820151816001015560408201518160020190600019169055606082015181600301555050505050505050565b61102f612263565b611037612015565b61103f612015565b611047612015565b61104f612015565b611057612015565b6110968a6040805190810160405280600681526020017f446f6e6174650000000000000000000000000000000000000000000000000000815250611b78565b156111c4576000600160008b6000191660001916815260200190815260200160002060009054906101000a900460ff1660ff1614156111bf5760036040518059106110de5750595b908082528060200260200182016040525095506060604051908101604052806000600581111561110a57fe5b815260200160008152602001878152509450846000808b6000191660001916815260200190815260200160002060008201518160000160006101000a81548160ff0219169083600581111561115b57fe5b02179055506020820151816001015560408201518160020190805190602001906111869291906120a6565b5090505060018060008b6000191660001916815260200190815260200160002060006101000a81548160ff021916908360ff1602179055505b611b6c565b6112038a6040805190810160405280600881526020017f44656c6976657231000000000000000000000000000000000000000000000000815250611b78565b15611453576000808a60001916600019168152602001908152602001600020606060405190810160405290816000820160009054906101000a900460ff16600581111561124c57fe5b600581111561125757fe5b815260200160018201548152602001600282018054806020026020016040519081016040528092919081815260200182805480156112ea57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116112a0575b50505050508152505093506000600581111561130257fe5b8460000151600581111561131257fe5b141561144e5760018460000190600581111561132a57fe5b9081600581111561133757fe5b8152505087846020015101846020018181525050868460400151600081518110151561135f57fe5b9060200190602002019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505087600360008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550836000808b6000191660001916815260200190815260200160002060008201518160000160006101000a81548160ff0219169083600581111561141e57fe5b02179055506020820151816001015560408201518160020190805190602001906114499291906120a6565b509050505b611b6b565b6114928a6040805190810160405280600581526020017f53746f7265000000000000000000000000000000000000000000000000000000815250611b78565b1561164e576000808a60001916600019168152602001908152602001600020606060405190810160405290816000820160009054906101000a900460ff1660058111156114db57fe5b60058111156114e657fe5b8152602001600182015481526020016002820180548060200260200160405190810160405280929190818152602001828054801561157957602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831161152f575b50505050508152505092506001600581111561159157fe5b836000015160058111156115a157fe5b1415611649576002836000019060058111156115b957fe5b908160058111156115c657fe5b8152505087836020015101836020018181525050826000808b6000191660001916815260200190815260200160002060008201518160000160006101000a81548160ff0219169083600581111561161957fe5b02179055506020820151816001015560408201518160020190805190602001906116449291906120a6565b509050505b611b6a565b61168d8a6040805190810160405280600481526020017f5368697000000000000000000000000000000000000000000000000000000000815250611b78565b156118dd576000808a60001916600019168152602001908152602001600020606060405190810160405290816000820160009054906101000a900460ff1660058111156116d657fe5b60058111156116e157fe5b8152602001600182015481526020016002820180548060200260200160405190810160405280929190818152602001828054801561177457602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831161172a575b50505050508152505091506002600581111561178c57fe5b8260000151600581111561179c57fe5b14156118d8576003826000019060058111156117b457fe5b908160058111156117c157fe5b815250508782602001510182602001818152505086826040015160018151811015156117e957fe5b9060200190602002019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505087600360008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816000808b6000191660001916815260200190815260200160002060008201518160000160006101000a81548160ff021916908360058111156118a857fe5b02179055506020820151816001015560408201518160020190805190602001906118d39291906120a6565b509050505b611b69565b61191c8a6040805190810160405280600881526020017f44656c6976657232000000000000000000000000000000000000000000000000815250611b78565b15611b68576000808a60001916600019168152602001908152602001600020606060405190810160405290816000820160009054906101000a900460ff16600581111561196557fe5b600581111561197057fe5b81526020016001820154815260200160028201805480602002602001604051908101604052809291908181526020018280548015611a0357602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116119b9575b505050505081525050905060036005811115611a1b57fe5b81600001516005811115611a2b57fe5b1415611b6757600481600001906005811115611a4357fe5b90816005811115611a5057fe5b81525050878160200151018160200181815250508681604001516002815181101515611a7857fe5b9060200190602002019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505087600360008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550806000808b6000191660001916815260200190815260200160002060008201518160000160006101000a81548160ff02191690836005811115611b3757fe5b0217905550602082015181600101556040820151816002019080519060200190611b629291906120a6565b509050505b5b5b5b5b5b50505050505050505050565b6000816040518082805190602001908083835b602083101515611bb05780518252602082019150602081019050602083039250611b8b565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902060001916836040518082805190602001908083835b602083101515611c175780518252602082019150602081019050602083039250611bf2565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390206000191614905092915050565b60075481565b80600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050565b6000806000806000806000806000611cb561218a565b610100604051908101604052808a81526020018981526020018860001916815260200186815260200185815260200184815260200183815260200187815250905060098054806001018281611d0a91906121d3565b91600052602060002090600802016000839091909150600082015181600001556020820151816001015560408201518160020190600019169055606082015181600301556080820151816004015560a0820151816005015560c0820151816006015560e0820151816007015550505060019a50600980549050600a81905550600a549950505050505050505050919050565b611da4612277565b611db3336402540be400611c57565b60a0604051908101604052808781526020018681526020018560001916815260200184815260200183815250905082600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282540392505081905550611e6f6040805190810160405280600881526020017f44656c6976657231000000000000000000000000000000000000000000000000815250858533611027565b60058054806001018281611e8391906122aa565b916000526020600020906005020160008390919091506000820151816000015560208201518160010155604082015181600201906000191690556060820151816003015560808201518160040155505050505050505050565b608060405190810160405280600081526020016000815260200160008019168152602001600081525090565b815481835581811511611f3557600402816004028360005260206000209182019101611f3491906122dc565b5b505050565b60e060405190810160405280600081526020016000815260200160008152602001600081526020016000801916815260200160008019168152602001600081525090565b815481835581811511611fab57600702816007028360005260206000209182019101611faa919061231b565b5b505050565b60a06040519081016040528060008152602001600081526020016000801916815260200160008152602001600081525090565b8154818355818115116120105760050281600502836000526020600020918201910161200f9190612372565b5b505050565b6060604051908101604052806000600581111561202e57fe5b8152602001600081526020016120426123b9565b81525090565b608060405190810160405280600081526020016000815260200160008152602001600080191681525090565b8154818355818115116120a1576004028160040283600052602060002091820191016120a091906123cd565b5b505050565b82805482825590600052602060002090810192821561211f579160200282015b8281111561211e5782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550916020019190600101906120c6565b5b50905061212c919061240c565b5090565b6060604051908101604052806000801916815260200160008152602001600080191681525090565b81548183558181151161218557600302816003028360005260206000209182019101612184919061244f565b5b505050565b6101006040519081016040528060008152602001600081526020016000801916815260200160008152602001600081526020016000815260200160008152602001600081525090565b815481835581811511612200576008028160080283600052602060002091820191016121ff9190612486565b5b505050565b608060405190810160405280600081526020016000815260200160008019168152602001600081525090565b81548183558181151161225e5760040281600402836000526020600020918201910161225d91906124e5565b5b505050565b602060405190810160405280600081525090565b60a06040519081016040528060008152602001600081526020016000801916815260200160008152602001600081525090565b8154818355818115116122d7576005028160050283600052602060002091820191016122d69190612524565b5b505050565b61231891905b8082111561231457600080820160009055600182016000905560028201600090556003820160009055506004016122e2565b5090565b90565b61236f91905b8082111561236b5760008082016000905560018201600090556002820160009055600382016000905560048201600090556005820160009055600682016000905550600701612321565b5090565b90565b6123b691905b808211156123b257600080820160009055600182016000905560028201600090556003820160009055600482016000905550600501612378565b5090565b90565b602060405190810160405280600081525090565b61240991905b8082111561240557600080820160009055600182016000905560028201600090556003820160009055506004016123d3565b5090565b90565b61244c91905b8082111561244857600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550600101612412565b5090565b90565b61248391905b8082111561247f576000808201600090556001820160009055600282016000905550600301612455565b5090565b90565b6124e291905b808211156124de5760008082016000905560018201600090556002820160009055600382016000905560048201600090556005820160009055600682016000905560078201600090555060080161248c565b5090565b90565b61252191905b8082111561251d57600080820160009055600182016000905560028201600090556003820160009055506004016124eb565b5090565b90565b61256891905b808211156125645760008082016000905560018201600090556002820160009055600382016000905560048201600090555060050161252a565b5090565b905600a165627a7a72305820f78aa7ea5c83e8bce2ebbac97af2d7fc3deec0954fe0164be50d38ac8d95167b0029",
}

// DoSCABI is the input ABI used to generate the binding from.
// Deprecated: Use DoSCMetaData.ABI instead.
var DoSCABI = DoSCMetaData.ABI

// DoSCBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DoSCMetaData.Bin instead.
var DoSCBin = DoSCMetaData.Bin

// DeployDoSC deploys a new Ethereum contract, binding an instance of DoSC to it.
func DeployDoSC(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DoSC, error) {
	parsed, err := DoSCMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DoSCBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DoSC{DoSCCaller: DoSCCaller{contract: contract}, DoSCTransactor: DoSCTransactor{contract: contract}, DoSCFilterer: DoSCFilterer{contract: contract}}, nil
}

// DoSC is an auto generated Go binding around an Ethereum contract.
type DoSC struct {
	DoSCCaller     // Read-only binding to the contract
	DoSCTransactor // Write-only binding to the contract
	DoSCFilterer   // Log filterer for contract events
}

// DoSCCaller is an auto generated read-only Go binding around an Ethereum contract.
type DoSCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DoSCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DoSCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DoSCFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DoSCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DoSCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DoSCSession struct {
	Contract     *DoSC             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DoSCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DoSCCallerSession struct {
	Contract *DoSCCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DoSCTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DoSCTransactorSession struct {
	Contract     *DoSCTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DoSCRaw is an auto generated low-level Go binding around an Ethereum contract.
type DoSCRaw struct {
	Contract *DoSC // Generic contract binding to access the raw methods on
}

// DoSCCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DoSCCallerRaw struct {
	Contract *DoSCCaller // Generic read-only contract binding to access the raw methods on
}

// DoSCTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DoSCTransactorRaw struct {
	Contract *DoSCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDoSC creates a new instance of DoSC, bound to a specific deployed contract.
func NewDoSC(address common.Address, backend bind.ContractBackend) (*DoSC, error) {
	contract, err := bindDoSC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DoSC{DoSCCaller: DoSCCaller{contract: contract}, DoSCTransactor: DoSCTransactor{contract: contract}, DoSCFilterer: DoSCFilterer{contract: contract}}, nil
}

// NewDoSCCaller creates a new read-only instance of DoSC, bound to a specific deployed contract.
func NewDoSCCaller(address common.Address, caller bind.ContractCaller) (*DoSCCaller, error) {
	contract, err := bindDoSC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DoSCCaller{contract: contract}, nil
}

// NewDoSCTransactor creates a new write-only instance of DoSC, bound to a specific deployed contract.
func NewDoSCTransactor(address common.Address, transactor bind.ContractTransactor) (*DoSCTransactor, error) {
	contract, err := bindDoSC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DoSCTransactor{contract: contract}, nil
}

// NewDoSCFilterer creates a new log filterer instance of DoSC, bound to a specific deployed contract.
func NewDoSCFilterer(address common.Address, filterer bind.ContractFilterer) (*DoSCFilterer, error) {
	contract, err := bindDoSC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DoSCFilterer{contract: contract}, nil
}

// bindDoSC binds a generic wrapper to an already deployed contract.
func bindDoSC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DoSCABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DoSC *DoSCRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DoSC.Contract.DoSCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DoSC *DoSCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DoSC.Contract.DoSCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DoSC *DoSCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DoSC.Contract.DoSCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DoSC *DoSCCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DoSC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DoSC *DoSCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DoSC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DoSC *DoSCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DoSC.Contract.contract.Transact(opts, method, params...)
}

// CompareStrings is a free data retrieval call binding the contract method 0xbed34bba.
//
// Solidity: function compareStrings(string a, string b) view returns(bool)
func (_DoSC *DoSCCaller) CompareStrings(opts *bind.CallOpts, a string, b string) (bool, error) {
	var out []interface{}
	err := _DoSC.contract.Call(opts, &out, "compareStrings", a, b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CompareStrings is a free data retrieval call binding the contract method 0xbed34bba.
//
// Solidity: function compareStrings(string a, string b) view returns(bool)
func (_DoSC *DoSCSession) CompareStrings(a string, b string) (bool, error) {
	return _DoSC.Contract.CompareStrings(&_DoSC.CallOpts, a, b)
}

// CompareStrings is a free data retrieval call binding the contract method 0xbed34bba.
//
// Solidity: function compareStrings(string a, string b) view returns(bool)
func (_DoSC *DoSCCallerSession) CompareStrings(a string, b string) (bool, error) {
	return _DoSC.Contract.CompareStrings(&_DoSC.CallOpts, a, b)
}

// Len0 is a free data retrieval call binding the contract method 0xc3a97b84.
//
// Solidity: function len0() view returns(uint256)
func (_DoSC *DoSCCaller) Len0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DoSC.contract.Call(opts, &out, "len0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Len0 is a free data retrieval call binding the contract method 0xc3a97b84.
//
// Solidity: function len0() view returns(uint256)
func (_DoSC *DoSCSession) Len0() (*big.Int, error) {
	return _DoSC.Contract.Len0(&_DoSC.CallOpts)
}

// Len0 is a free data retrieval call binding the contract method 0xc3a97b84.
//
// Solidity: function len0() view returns(uint256)
func (_DoSC *DoSCCallerSession) Len0() (*big.Int, error) {
	return _DoSC.Contract.Len0(&_DoSC.CallOpts)
}

// Len1 is a free data retrieval call binding the contract method 0x79d3202e.
//
// Solidity: function len1() view returns(uint256)
func (_DoSC *DoSCCaller) Len1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DoSC.contract.Call(opts, &out, "len1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Len1 is a free data retrieval call binding the contract method 0x79d3202e.
//
// Solidity: function len1() view returns(uint256)
func (_DoSC *DoSCSession) Len1() (*big.Int, error) {
	return _DoSC.Contract.Len1(&_DoSC.CallOpts)
}

// Len1 is a free data retrieval call binding the contract method 0x79d3202e.
//
// Solidity: function len1() view returns(uint256)
func (_DoSC *DoSCCallerSession) Len1() (*big.Int, error) {
	return _DoSC.Contract.Len1(&_DoSC.CallOpts)
}

// Deliver1 is a paid mutator transaction binding the contract method 0xe5aa1c36.
//
// Solidity: function Deliver1(uint256 pkcoix, uint256 pkcoiy, bytes32 hclj, uint256 dp, uint256 dt) returns()
func (_DoSC *DoSCTransactor) Deliver1(opts *bind.TransactOpts, pkcoix *big.Int, pkcoiy *big.Int, hclj [32]byte, dp *big.Int, dt *big.Int) (*types.Transaction, error) {
	return _DoSC.contract.Transact(opts, "Deliver1", pkcoix, pkcoiy, hclj, dp, dt)
}

// Deliver1 is a paid mutator transaction binding the contract method 0xe5aa1c36.
//
// Solidity: function Deliver1(uint256 pkcoix, uint256 pkcoiy, bytes32 hclj, uint256 dp, uint256 dt) returns()
func (_DoSC *DoSCSession) Deliver1(pkcoix *big.Int, pkcoiy *big.Int, hclj [32]byte, dp *big.Int, dt *big.Int) (*types.Transaction, error) {
	return _DoSC.Contract.Deliver1(&_DoSC.TransactOpts, pkcoix, pkcoiy, hclj, dp, dt)
}

// Deliver1 is a paid mutator transaction binding the contract method 0xe5aa1c36.
//
// Solidity: function Deliver1(uint256 pkcoix, uint256 pkcoiy, bytes32 hclj, uint256 dp, uint256 dt) returns()
func (_DoSC *DoSCTransactorSession) Deliver1(pkcoix *big.Int, pkcoiy *big.Int, hclj [32]byte, dp *big.Int, dt *big.Int) (*types.Transaction, error) {
	return _DoSC.Contract.Deliver1(&_DoSC.TransactOpts, pkcoix, pkcoiy, hclj, dp, dt)
}

// Deliver2 is a paid mutator transaction binding the contract method 0x2c5ecd73.
//
// Solidity: function Deliver2(uint256 pkcoix, uint256 pkcoiy, bytes32 hclj, uint256 dp, uint256 dt) returns()
func (_DoSC *DoSCTransactor) Deliver2(opts *bind.TransactOpts, pkcoix *big.Int, pkcoiy *big.Int, hclj [32]byte, dp *big.Int, dt *big.Int) (*types.Transaction, error) {
	return _DoSC.contract.Transact(opts, "Deliver2", pkcoix, pkcoiy, hclj, dp, dt)
}

// Deliver2 is a paid mutator transaction binding the contract method 0x2c5ecd73.
//
// Solidity: function Deliver2(uint256 pkcoix, uint256 pkcoiy, bytes32 hclj, uint256 dp, uint256 dt) returns()
func (_DoSC *DoSCSession) Deliver2(pkcoix *big.Int, pkcoiy *big.Int, hclj [32]byte, dp *big.Int, dt *big.Int) (*types.Transaction, error) {
	return _DoSC.Contract.Deliver2(&_DoSC.TransactOpts, pkcoix, pkcoiy, hclj, dp, dt)
}

// Deliver2 is a paid mutator transaction binding the contract method 0x2c5ecd73.
//
// Solidity: function Deliver2(uint256 pkcoix, uint256 pkcoiy, bytes32 hclj, uint256 dp, uint256 dt) returns()
func (_DoSC *DoSCTransactorSession) Deliver2(pkcoix *big.Int, pkcoiy *big.Int, hclj [32]byte, dp *big.Int, dt *big.Int) (*types.Transaction, error) {
	return _DoSC.Contract.Deliver2(&_DoSC.TransactOpts, pkcoix, pkcoiy, hclj, dp, dt)
}

// Distribute is a paid mutator transaction binding the contract method 0x84d298ff.
//
// Solidity: function Distribute(uint256 pkcoix, uint256 pkcoiy, bytes32 hclj, uint256 cmqx, uint256 cmqy, uint256 tokqx, uint256 tokqy) returns()
func (_DoSC *DoSCTransactor) Distribute(opts *bind.TransactOpts, pkcoix *big.Int, pkcoiy *big.Int, hclj [32]byte, cmqx *big.Int, cmqy *big.Int, tokqx *big.Int, tokqy *big.Int) (*types.Transaction, error) {
	return _DoSC.contract.Transact(opts, "Distribute", pkcoix, pkcoiy, hclj, cmqx, cmqy, tokqx, tokqy)
}

// Distribute is a paid mutator transaction binding the contract method 0x84d298ff.
//
// Solidity: function Distribute(uint256 pkcoix, uint256 pkcoiy, bytes32 hclj, uint256 cmqx, uint256 cmqy, uint256 tokqx, uint256 tokqy) returns()
func (_DoSC *DoSCSession) Distribute(pkcoix *big.Int, pkcoiy *big.Int, hclj [32]byte, cmqx *big.Int, cmqy *big.Int, tokqx *big.Int, tokqy *big.Int) (*types.Transaction, error) {
	return _DoSC.Contract.Distribute(&_DoSC.TransactOpts, pkcoix, pkcoiy, hclj, cmqx, cmqy, tokqx, tokqy)
}

// Distribute is a paid mutator transaction binding the contract method 0x84d298ff.
//
// Solidity: function Distribute(uint256 pkcoix, uint256 pkcoiy, bytes32 hclj, uint256 cmqx, uint256 cmqy, uint256 tokqx, uint256 tokqy) returns()
func (_DoSC *DoSCTransactorSession) Distribute(pkcoix *big.Int, pkcoiy *big.Int, hclj [32]byte, cmqx *big.Int, cmqy *big.Int, tokqx *big.Int, tokqy *big.Int) (*types.Transaction, error) {
	return _DoSC.Contract.Distribute(&_DoSC.TransactOpts, pkcoix, pkcoiy, hclj, cmqx, cmqy, tokqx, tokqy)
}

// Donate is a paid mutator transaction binding the contract method 0x1463f015.
//
// Solidity: function Donate(uint256 comx, uint256 comy, uint256 tokx, uint256 toky, bytes32 hres, bytes32 hcl, uint256 dt) returns()
func (_DoSC *DoSCTransactor) Donate(opts *bind.TransactOpts, comx *big.Int, comy *big.Int, tokx *big.Int, toky *big.Int, hres [32]byte, hcl [32]byte, dt *big.Int) (*types.Transaction, error) {
	return _DoSC.contract.Transact(opts, "Donate", comx, comy, tokx, toky, hres, hcl, dt)
}

// Donate is a paid mutator transaction binding the contract method 0x1463f015.
//
// Solidity: function Donate(uint256 comx, uint256 comy, uint256 tokx, uint256 toky, bytes32 hres, bytes32 hcl, uint256 dt) returns()
func (_DoSC *DoSCSession) Donate(comx *big.Int, comy *big.Int, tokx *big.Int, toky *big.Int, hres [32]byte, hcl [32]byte, dt *big.Int) (*types.Transaction, error) {
	return _DoSC.Contract.Donate(&_DoSC.TransactOpts, comx, comy, tokx, toky, hres, hcl, dt)
}

// Donate is a paid mutator transaction binding the contract method 0x1463f015.
//
// Solidity: function Donate(uint256 comx, uint256 comy, uint256 tokx, uint256 toky, bytes32 hres, bytes32 hcl, uint256 dt) returns()
func (_DoSC *DoSCTransactorSession) Donate(comx *big.Int, comy *big.Int, tokx *big.Int, toky *big.Int, hres [32]byte, hcl [32]byte, dt *big.Int) (*types.Transaction, error) {
	return _DoSC.Contract.Donate(&_DoSC.TransactOpts, comx, comy, tokx, toky, hres, hcl, dt)
}

// Receive is a paid mutator transaction binding the contract method 0x746cac8d.
//
// Solidity: function Receive(bytes32 hfbi, uint256 dt, bytes32 hcl) returns()
func (_DoSC *DoSCTransactor) Receive(opts *bind.TransactOpts, hfbi [32]byte, dt *big.Int, hcl [32]byte) (*types.Transaction, error) {
	return _DoSC.contract.Transact(opts, "Receive", hfbi, dt, hcl)
}

// Receive is a paid mutator transaction binding the contract method 0x746cac8d.
//
// Solidity: function Receive(bytes32 hfbi, uint256 dt, bytes32 hcl) returns()
func (_DoSC *DoSCSession) Receive(hfbi [32]byte, dt *big.Int, hcl [32]byte) (*types.Transaction, error) {
	return _DoSC.Contract.Receive(&_DoSC.TransactOpts, hfbi, dt, hcl)
}

// Receive is a paid mutator transaction binding the contract method 0x746cac8d.
//
// Solidity: function Receive(bytes32 hfbi, uint256 dt, bytes32 hcl) returns()
func (_DoSC *DoSCTransactorSession) Receive(hfbi [32]byte, dt *big.Int, hcl [32]byte) (*types.Transaction, error) {
	return _DoSC.Contract.Receive(&_DoSC.TransactOpts, hfbi, dt, hcl)
}

// Refund is a paid mutator transaction binding the contract method 0x47b788f2.
//
// Solidity: function Refund(uint256 pkcoix, uint256 pkcoiy, uint256 dt, bytes32 hcl) returns()
func (_DoSC *DoSCTransactor) Refund(opts *bind.TransactOpts, pkcoix *big.Int, pkcoiy *big.Int, dt *big.Int, hcl [32]byte) (*types.Transaction, error) {
	return _DoSC.contract.Transact(opts, "Refund", pkcoix, pkcoiy, dt, hcl)
}

// Refund is a paid mutator transaction binding the contract method 0x47b788f2.
//
// Solidity: function Refund(uint256 pkcoix, uint256 pkcoiy, uint256 dt, bytes32 hcl) returns()
func (_DoSC *DoSCSession) Refund(pkcoix *big.Int, pkcoiy *big.Int, dt *big.Int, hcl [32]byte) (*types.Transaction, error) {
	return _DoSC.Contract.Refund(&_DoSC.TransactOpts, pkcoix, pkcoiy, dt, hcl)
}

// Refund is a paid mutator transaction binding the contract method 0x47b788f2.
//
// Solidity: function Refund(uint256 pkcoix, uint256 pkcoiy, uint256 dt, bytes32 hcl) returns()
func (_DoSC *DoSCTransactorSession) Refund(pkcoix *big.Int, pkcoiy *big.Int, dt *big.Int, hcl [32]byte) (*types.Transaction, error) {
	return _DoSC.Contract.Refund(&_DoSC.TransactOpts, pkcoix, pkcoiy, dt, hcl)
}

// Ship is a paid mutator transaction binding the contract method 0x95e48145.
//
// Solidity: function Ship(uint256 pkcoix, uint256 pkcoiy, bytes32 hclj, uint256 dp) returns()
func (_DoSC *DoSCTransactor) Ship(opts *bind.TransactOpts, pkcoix *big.Int, pkcoiy *big.Int, hclj [32]byte, dp *big.Int) (*types.Transaction, error) {
	return _DoSC.contract.Transact(opts, "Ship", pkcoix, pkcoiy, hclj, dp)
}

// Ship is a paid mutator transaction binding the contract method 0x95e48145.
//
// Solidity: function Ship(uint256 pkcoix, uint256 pkcoiy, bytes32 hclj, uint256 dp) returns()
func (_DoSC *DoSCSession) Ship(pkcoix *big.Int, pkcoiy *big.Int, hclj [32]byte, dp *big.Int) (*types.Transaction, error) {
	return _DoSC.Contract.Ship(&_DoSC.TransactOpts, pkcoix, pkcoiy, hclj, dp)
}

// Ship is a paid mutator transaction binding the contract method 0x95e48145.
//
// Solidity: function Ship(uint256 pkcoix, uint256 pkcoiy, bytes32 hclj, uint256 dp) returns()
func (_DoSC *DoSCTransactorSession) Ship(pkcoix *big.Int, pkcoiy *big.Int, hclj [32]byte, dp *big.Int) (*types.Transaction, error) {
	return _DoSC.Contract.Ship(&_DoSC.TransactOpts, pkcoix, pkcoiy, hclj, dp)
}

// Store is a paid mutator transaction binding the contract method 0x073f21a9.
//
// Solidity: function Store(uint256 pkcoix, uint256 pkcoiy, bytes32 hclj, uint256 dt) returns()
func (_DoSC *DoSCTransactor) Store(opts *bind.TransactOpts, pkcoix *big.Int, pkcoiy *big.Int, hclj [32]byte, dt *big.Int) (*types.Transaction, error) {
	return _DoSC.contract.Transact(opts, "Store", pkcoix, pkcoiy, hclj, dt)
}

// Store is a paid mutator transaction binding the contract method 0x073f21a9.
//
// Solidity: function Store(uint256 pkcoix, uint256 pkcoiy, bytes32 hclj, uint256 dt) returns()
func (_DoSC *DoSCSession) Store(pkcoix *big.Int, pkcoiy *big.Int, hclj [32]byte, dt *big.Int) (*types.Transaction, error) {
	return _DoSC.Contract.Store(&_DoSC.TransactOpts, pkcoix, pkcoiy, hclj, dt)
}

// Store is a paid mutator transaction binding the contract method 0x073f21a9.
//
// Solidity: function Store(uint256 pkcoix, uint256 pkcoiy, bytes32 hclj, uint256 dt) returns()
func (_DoSC *DoSCTransactorSession) Store(pkcoix *big.Int, pkcoiy *big.Int, hclj [32]byte, dt *big.Int) (*types.Transaction, error) {
	return _DoSC.Contract.Store(&_DoSC.TransactOpts, pkcoix, pkcoiy, hclj, dt)
}

// Transfer is a paid mutator transaction binding the contract method 0x9875a1e2.
//
// Solidity: function Transfer(string txx, bytes32 hcl, uint256 dp, address from) returns()
func (_DoSC *DoSCTransactor) Transfer(opts *bind.TransactOpts, txx string, hcl [32]byte, dp *big.Int, from common.Address) (*types.Transaction, error) {
	return _DoSC.contract.Transact(opts, "Transfer", txx, hcl, dp, from)
}

// Transfer is a paid mutator transaction binding the contract method 0x9875a1e2.
//
// Solidity: function Transfer(string txx, bytes32 hcl, uint256 dp, address from) returns()
func (_DoSC *DoSCSession) Transfer(txx string, hcl [32]byte, dp *big.Int, from common.Address) (*types.Transaction, error) {
	return _DoSC.Contract.Transfer(&_DoSC.TransactOpts, txx, hcl, dp, from)
}

// Transfer is a paid mutator transaction binding the contract method 0x9875a1e2.
//
// Solidity: function Transfer(string txx, bytes32 hcl, uint256 dp, address from) returns()
func (_DoSC *DoSCTransactorSession) Transfer(txx string, hcl [32]byte, dp *big.Int, from common.Address) (*types.Transaction, error) {
	return _DoSC.Contract.Transfer(&_DoSC.TransactOpts, txx, hcl, dp, from)
}

// PrintLenOfDis is a paid mutator transaction binding the contract method 0xcf938157.
//
// Solidity: function printLenOfDis(uint256 aaa) returns(uint256)
func (_DoSC *DoSCTransactor) PrintLenOfDis(opts *bind.TransactOpts, aaa *big.Int) (*types.Transaction, error) {
	return _DoSC.contract.Transact(opts, "printLenOfDis", aaa)
}

// PrintLenOfDis is a paid mutator transaction binding the contract method 0xcf938157.
//
// Solidity: function printLenOfDis(uint256 aaa) returns(uint256)
func (_DoSC *DoSCSession) PrintLenOfDis(aaa *big.Int) (*types.Transaction, error) {
	return _DoSC.Contract.PrintLenOfDis(&_DoSC.TransactOpts, aaa)
}

// PrintLenOfDis is a paid mutator transaction binding the contract method 0xcf938157.
//
// Solidity: function printLenOfDis(uint256 aaa) returns(uint256)
func (_DoSC *DoSCTransactorSession) PrintLenOfDis(aaa *big.Int) (*types.Transaction, error) {
	return _DoSC.Contract.PrintLenOfDis(&_DoSC.TransactOpts, aaa)
}

// PrintLenOfStore is a paid mutator transaction binding the contract method 0x0862c570.
//
// Solidity: function printLenOfStore(uint256 aaa) returns(uint256)
func (_DoSC *DoSCTransactor) PrintLenOfStore(opts *bind.TransactOpts, aaa *big.Int) (*types.Transaction, error) {
	return _DoSC.contract.Transact(opts, "printLenOfStore", aaa)
}

// PrintLenOfStore is a paid mutator transaction binding the contract method 0x0862c570.
//
// Solidity: function printLenOfStore(uint256 aaa) returns(uint256)
func (_DoSC *DoSCSession) PrintLenOfStore(aaa *big.Int) (*types.Transaction, error) {
	return _DoSC.Contract.PrintLenOfStore(&_DoSC.TransactOpts, aaa)
}

// PrintLenOfStore is a paid mutator transaction binding the contract method 0x0862c570.
//
// Solidity: function printLenOfStore(uint256 aaa) returns(uint256)
func (_DoSC *DoSCTransactorSession) PrintLenOfStore(aaa *big.Int) (*types.Transaction, error) {
	return _DoSC.Contract.PrintLenOfStore(&_DoSC.TransactOpts, aaa)
}

// TestAddDp is a paid mutator transaction binding the contract method 0xce09d470.
//
// Solidity: function test_add_dp(address a, uint256 k) returns()
func (_DoSC *DoSCTransactor) TestAddDp(opts *bind.TransactOpts, a common.Address, k *big.Int) (*types.Transaction, error) {
	return _DoSC.contract.Transact(opts, "test_add_dp", a, k)
}

// TestAddDp is a paid mutator transaction binding the contract method 0xce09d470.
//
// Solidity: function test_add_dp(address a, uint256 k) returns()
func (_DoSC *DoSCSession) TestAddDp(a common.Address, k *big.Int) (*types.Transaction, error) {
	return _DoSC.Contract.TestAddDp(&_DoSC.TransactOpts, a, k)
}

// TestAddDp is a paid mutator transaction binding the contract method 0xce09d470.
//
// Solidity: function test_add_dp(address a, uint256 k) returns()
func (_DoSC *DoSCTransactorSession) TestAddDp(a common.Address, k *big.Int) (*types.Transaction, error) {
	return _DoSC.Contract.TestAddDp(&_DoSC.TransactOpts, a, k)
}
