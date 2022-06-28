/***************************************************************************
 *
 * Copyright (c) 2017 Baidu.com, Inc. All Rights Reserved
 * @author duanbing(duanbing@baidu.com)
 *
 **************************************************************************/

/**
 * @filename main.go
 * @desc
 * @create time 2018-04-19 15:49:26
**/
package main

import (
	"crypto/rand"
	"fmt"
	ec "github.com/duanbing/go-evm/core"
	"github.com/duanbing/go-evm/state"
	"github.com/duanbing/go-evm/types"
	"github.com/duanbing/go-evm/vm"
	"reflect"

	// "github.com/edgelesssys/ego/enclave"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	//_ "github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/params"
	//Ethclient, connect with geth
	//See https://www.quicknode.com/guides/web3-sdks/how-to-connect-to-ethereum-network-using-go
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"time"
)

var (
	//testHash    = common.StringToHash("duanbing")
	testAddress = common.StringToAddress("ubiplab")
	toAddress   = common.StringToAddress("andone")
	amount      = big.NewInt(1)
	nonce       = uint64(0)
	gasLimit    = big.NewInt(100000)
	coinbase    = common.HexToAddress("0x0000000000000000000000000000000000000000")
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}
func loadBin(filename string) []byte {
	code, err := ioutil.ReadFile(filename)
	must(err)
	return hexutil.MustDecode("0x" + string(code))
}
func loadAbi(filename string) abi.ABI {
	abiFile, err := os.Open(filename)
	must(err)
	defer abiFile.Close()
	abiObj, err := abi.JSON(abiFile)
	must(err)
	return abiObj
}

type evm_func struct {
	statedb *state.StateDB
	abiOBJ abi.ABI
	evm *vm.EVM
	contractRef vm.AccountRef
	testBalance *big.Int
}
var ef evm_func//Global var

func main() {


	binFileName := "examples/ECC/RStructTest.bin"
	abiFileName := "examples/ECC/RStructTest.abi"
	data := loadBin(binFileName)
	msg := ec.NewMessage(testAddress, &toAddress, nonce, amount, gasLimit, big.NewInt(1), data, false)
	header := types.Header{
		// ParentHash: common.Hash{},
		// UncleHash:  common.Hash{},
		Coinbase: coinbase,
		//	Root:        common.Hash{},
		//	TxHash:      common.Hash{},
		//	ReceiptHash: common.Hash{},
		//	Bloom:      types.BytesToBloom([]byte("duanbing")),
		Difficulty: big.NewInt(1),
		Number:     big.NewInt(1),
		GasLimit:   100000000,
		GasUsed:    0,
		Time:       big.NewInt(time.Now().Unix()),
		Extra:      nil,
		//MixDigest:  testHash,
		//Nonce:      types.EncodeNonce(1),
	}
	cc := ChainContext{}
	ctx := ec.NewEVMContext(msg, &header, cc, &testAddress)
	//mdb, err := ethdb.NewMemDatabase()
	dataPath := "examples/ECC/tmp/a.txt"
	os.Remove(dataPath)
	mdb, err := ethdb.NewLDBDatabase(dataPath, 100, 100)
	must(err)
	db := state.NewDatabase(mdb)
	statedb, err := state.New(common.Hash{}, db)
	//set balance
	statedb.GetOrNewStateObject(testAddress)
	statedb.GetOrNewStateObject(toAddress)
	statedb.AddBalance(testAddress, big.NewInt(1e18))
	testBalance := statedb.GetBalance(testAddress)
	ef.testBalance=testBalance
	fmt.Println("init testBalance =", testBalance)
	must(err)

	//	config := params.TestnetChainConfig
	config := params.TestnetChainConfig
	logConfig := vm.LogConfig{}
	structLogger := vm.NewStructLogger(&logConfig)
	vmConfig := vm.Config{Debug: true, Tracer: structLogger/*, JumpTable: vm.NewByzantiumInstructionSet()*/}

	evm := vm.NewEVM(ctx, statedb, config, vmConfig)
	contractRef := vm.AccountRef(testAddress)
	contractCode, _, gasLeftover, vmerr := evm.Create(contractRef, data, statedb.GetBalance(testAddress).Uint64(), big.NewInt(0))
	must(vmerr)
	statedb.SetBalance(testAddress, big.NewInt(0).SetUint64(gasLeftover))
	testBalance = statedb.GetBalance(testAddress)
	fmt.Println("after create contract, testBalance =", testBalance)
	// set input ,  formatted accocding to https://solidity.readthedocs.io/en/develop/abi-spec.html
	//find methods := "multiply(uint)"

	/*
	Contract Construction Finished
	*/
	abiObj := loadAbi(abiFileName)

	ef.statedb=statedb
	ef.abiOBJ=abiObj
	ef.evm=evm
	ef.contractRef=contractRef
	//Start HTTP Server here
	//Specify method and execute

	//method := abiObj.Methods["derivePubKey"]
	//make params := "0xa"
	//concat method and params
	//inputstr := hexutil.Encode(method.Id()) + pm[2:]
	fmt.Println("begin to exec contract")
	statedb.SetCode(testAddress, contractCode)


	//1.注册一个给定模式的处理器函数到DefaultServeMux
	http.HandleFunc("/mul", sayHello)

	//2.设置监听的TCP地址并启动服务
	//参数1：TCP地址(IP+Port)
	//参数2：当设置为nil时表示使用DefaultServeMux
	err = http.ListenAndServe("127.0.0.1:8080", nil)
	log.Fatal(err)

}

type ChainContext struct{}

func (cc ChainContext) GetHeader(hash common.Hash, number uint64) *types.Header {
	fmt.Println("(cc ChainContext) GetHeader(hash common.Hash, number uint64)")
	return nil
	//return &header
}

func sayHello(w http.ResponseWriter, r *http.Request) {

	for i:=1;i<=1;i++ {
		//input, err = abiObj.Pack("mint", sender, big.NewInt(1000000))
		//pm := abi.U256(big.NewInt(int64(100 * i)))
		//input := append(method.Id(), pm...)
		r,_:=rand.Prime(rand.Reader,5)
		input, err := ef.abiOBJ.Pack("test",r)
		//
		//fmt.Println(input)
		outputs, gasLeftover, vmerr := ef.evm.Call(ef.contractRef, testAddress, input, ef.statedb.GetBalance(testAddress).Uint64(), big.NewInt(0))
		must(vmerr)

		ef.statedb.SetBalance(testAddress, big.NewInt(0).SetUint64(gasLeftover))
		ef.testBalance = ef.statedb.GetBalance(testAddress)
		fmt.Println("after call contract, testBalance =", ef.testBalance)
		/*Get return data from output bytes
		First way:
		*/
		data,err:=ef.abiOBJ.Methods["test"].Outputs.UnpackValues(outputs)
		//Convertion
		/*Get return data from output bytes
		Second way:
		*/
		//data := []interface{}{new(int64), new(int64)}
		fmt.Println(data[0],reflect.TypeOf(data[0]))
		fmt.Println(data[1],reflect.TypeOf(data[1]))
		fmt.Println(data[2],reflect.TypeOf(data[2]))
		//返回结果为interface形式的byte32数组(hashs)时，使用断言转换为[][32]uint8类型，对[32]uint8使用[:]进行切片并转为hex字符串存储
		hashs:=data[2]
		value, ok := hashs.([][32]uint8)
		if !ok {
			fmt.Println("It's not ok for type string")
		}
		fmt.Println(common.Bytes2Hex(value[0][:]))
		//fmt.Println(common.Bytes2Hex(data[2]))

		if err != nil {
			log.Fatal(err)
		}
		//for _, op := range method.Outputs {
		//	switch op.Type.String() {
		//	case "uint256":
		//		tabi:=abiObj.Unpack(struct{
		//			Intleft *big.Int
		//			Intright *big.Int
		//		}{},"left",outputs)
		//
		//		fmt.Printf("Output name=%s, value=%d\n", op.Name, outputs)
		//
		//	default:
		//		fmt.Println(op.Name, op.Type.String())
		//	}
		//}

	}
}