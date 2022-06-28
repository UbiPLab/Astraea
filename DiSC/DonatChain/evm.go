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
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"time"

	//"zy/jxr/go-evm-master/examples/event/bwt"

	ec "github.com/duanbing/go-evm/core"
	"github.com/duanbing/go-evm/state"
	"github.com/duanbing/go-evm/types"
	"github.com/duanbing/go-evm/vm"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/params"
)

var (
	testHash    = common.StringToHash("duanbing")
	fromAddress = common.StringToAddress("duanbing") //交易发送方
	toAddress   = common.StringToAddress("andone")   //交易接收方
	amount      = big.NewInt(0)                      //交易数额
	nonce       = uint64(0)                          //选用随机数
	gasLimit    = big.NewInt(100000)                 //燃料上限
	coinbase    = fromAddress                        //矿工

	abiFileName    = "./test/Test_sol_Test.abi"
	binFileName    = "./test/Test_sol_Test.bin"
	dataPath       = "./test/a.txt"
	abiObj         abi.ABI
	statedb        *state.StateDB
	db             state.Database
	mdb            *ethdb.LDBDatabase
	evm            *vm.EVM //作为全局变量后方便拆分调用，但是在并发执行时可能出问题
	contractRef    vm.AccountRef
	contractCode   []byte
	contractAddr   common.Address
	response_count = 0
	response_MAX   = 5
)

type parameters struct { // 32 bytes per big.Int
	ReqRi     []byte
	ReqCi     []byte
	ReqDni    []byte
	ReqPkcoiX []byte
	ReqPkcoiY []byte
	tp1       []byte
	lc        []byte
	idj1      []byte
	clj       []byte
	dnj       []byte
	rj        []byte
	cmx       []byte
	cmy       []byte
	tokx      []byte
	toky      []byte
	dt        []byte
}

type Audition struct {
	AggX     *[]byte
	AggY     *[]byte
	deltaX   *[]byte
	deltaY   *[]byte
	H_R_SK_X *[]byte
	H_R_SK_Y *[]byte
	H_R_X    *[]byte
	H_R_Y    *[]byte
	AA       *[]byte
	dt       *[]byte
	//Signature *[]byte
}

type Distribution struct {
	pubKeyX *[]byte
	pubKeyY *[]byte
	hash    *[]byte
	cmX     *[]byte
	cmY     *[]byte
	tokX    *[]byte
	tokY    *[]byte
	dt      *[]byte
	//Signature *[]byte
}

type SendData struct {
	listLength int64
	dis        *[]byte
	aud1       *[]byte
	aud2       *[]byte
	clm        *[][]byte
}

func must(err error) { //有错误就退出
	if err != nil {
		panic(err)
	}
}
func loadBin(filename string) []byte { //加载智能合约二进制文件
	code, err := ioutil.ReadFile(filename)
	must(err)
	return hexutil.MustDecode("0x" + string(code))
	//return []byte("0x" + string(code))
}
func loadAbi(filename string) abi.ABI { //加载其二进制接口
	abiFile, err := os.Open(filename)
	must(err)
	defer abiFile.Close()            //返回前自动关闭文件
	abiObj, err := abi.JSON(abiFile) //通过json解析成abi对象
	must(err)
	return abiObj
}

func evmInit() {
	reinit()
	//虚拟机初始化，获得一个虚拟机实例，创建智能合约对象，智能合约内初始化（H）
	fmt.Println("初始化EVM")
	fmt.Println("加载合约文件")
	data := loadBin(binFileName) //载入智能合约二进制文件
	abiObj = loadAbi(abiFileName)
	//ec是evm core，一笔交易称作消息
	msg := ec.NewMessage(fromAddress, &toAddress, nonce, amount, gasLimit, big.NewInt(0), data, false)
	cc := ChainContext{}                                                      //区块链上下文环境，自定义的结构体
	ctx := ec.NewEVMContext(msg, cc.GetHeader(testHash, 0), cc, &fromAddress) //给evm创建新的执行环境
	//交易输出数据存放路径
	os.Remove(dataPath)
	var err error                                       //先清除之前的内容?那怎么保存结果呢
	mdb, err = ethdb.NewLDBDatabase(dataPath, 100, 100) //再创建新的数据库
	must(err)
	fmt.Println("获得数据库")
	db = state.NewDatabase(mdb) //虚拟机状态中添加这个数据库

	root := common.Hash{}              //树根的hash
	statedb, err = state.New(root, db) //数据库添加树根
	must(err)
	statedb.GetOrNewStateObject(fromAddress)
	statedb.GetOrNewStateObject(toAddress)
	_aa, _ := new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFF", 16)
	statedb.AddBalance(fromAddress, _aa)

	//	config := params.TestnetChainConfig
	config := params.MainnetChainConfig            //主链配置
	logConfig := vm.LogConfig{}                    //虚拟机日志配置
	structLogger := vm.NewStructLogger(&logConfig) //创建日志实例
	vmConfig := vm.Config{Debug: true, Tracer: structLogger /*, JumpTable: vm.NewByzantiumInstructionSet()*/}
	//这个函数主要是根据当前的区块号以及相关配置，设置EVM的解释器
	// evm执行上下文，数据库，所在链上的配置，虚拟机日志配置？
	evm = vm.NewEVM(ctx, statedb, config, vmConfig)
	fmt.Println("获得EVM实例")
	//先获得智能合约对象
	contractRef = vm.AccountRef(fromAddress) //合约地址设为调用方地址
	//Create()方法首先对发送者地址（caller.Address）和账户的nonce进行keccak256计算得到合约地址，然后将合约地址传入create()方法，后者是合约创建的真正函数
	_, contractAddr, _, _ = evm.Create(contractRef, data, statedb.GetBalance(fromAddress).Uint64(), big.NewInt(0))
	//must(vmerr)
	//statedb.SetBalance(fromAddress, big.NewInt(0).SetUint64(gasLeftover))
	fmt.Println("部署合约到私链") /*
		input, err := abiObj.Pack("AddOne", big.NewInt(3))
		must(err)
		fmt.Println("Pack")
		//value是转账额度
		outputs, _, vmerr := evm.Call(contractRef, contractAddr, input, statedb.GetBalance(fromAddress).Uint64(), big.NewInt(0))
		must(vmerr)
		fmt.Println("AddOne:\t", outputs)*/
	//初始化H
	fmt.Println("初始化合约内变量")
	input, err := abiObj.Pack("getNewGenerator", HX, HY, bigr1, bigr2, bigr, pkX, pkY)
	must(err)

	//value是转账额度
	_, _, vmerr := evm.Call(contractRef, contractAddr, input, statedb.GetBalance(fromAddress).Uint64(), big.NewInt(0))
	must(vmerr)
	//fmt.Println("outputs:\t", outputs)
}

func Evm(response_ch chan []byte, sd_ch chan *SendData) {
	fmt.Println("EVM守护协程")
	// 假设虚拟机已经初始化好了
	for {
		// 不断从管道获取response，调用智能合约
		tmp := <-response_ch
		//带点偏移复制过去
		//fmt.Println("解密后:\t", tmp)
		buf := make([]*big.Int, 16) //改长度
		for i := 0; i < 16; i++ {
			buf[i] = new(big.Int).SetBytes(tmp[i*32 : i*32+32])
		}
		/*
			//input, err := abiObj.Pack("Response", buf)
			input, err := abiObj.Pack("Response", buf[3], buf[4], buf[5], buf[6], buf[7], buf[8], buf[9], buf[10], buf[11], buf[12], buf[13])
			must(err)
			fmt.Println("调用Response")
			//response获得Store
			_, _, vmerr := evm.Call(contractRef, contractAddr, input, statedb.GetBalance(fromAddress).Uint64(), big.NewInt(0))
			must(vmerr)
			response_count += 1
			fmt.Println("response_count:\t", response_count)*/
		response_count += 1
		fmt.Println("已接收到的response数:\t", response_count)
		Response(buf)
		//调用distribute
		if response_count == response_MAX {

			//input, err = abiObj.Pack("ListLength")
			//must(err)
			//output, _, vmerr := evm.Call(contractRef, contractAddr, input, statedb.GetBalance(fromAddress).Uint64(), big.NewInt(0))
			//must(vmerr)
			//listLength := new(big.Int).SetBytes(output)
			l := int64(response_MAX) //listLength.Int64()

			/*
				fmt.Printf("Test")
				input, err = abiObj.Pack("testAdd", pkX, pkY)
				must(err)
				//(DisH,DisCmX,DisCmY);
				out, _, vmerr := evm.Call(contractRef, contractAddr, input, statedb.GetBalance(fromAddress).Uint64(), big.NewInt(0))
				must(vmerr)
				fmt.Println("(pkX,pkY)+(HX,HY) in disc:\t", new(big.Int).SetBytes(out[:32]), new(big.Int).SetBytes(out[32:]))
				tx, ty := secp256k1.S256().Add(pkX, pkY, HX, HY)
				fmt.Println("(pkX,pkY)+(HX,HY) in go:\t", tx, ty)

				fmt.Printf("Test")
				input, err = abiObj.Pack("testMul", pkX, pkY)
				must(err)
				//(DisH,DisCmX,DisCmY);
				out, _, vmerr = evm.Call(contractRef, contractAddr, input, statedb.GetBalance(fromAddress).Uint64(), big.NewInt(0))
				must(vmerr)
				fmt.Println("101*(pkX,pkY) in disc:\t", new(big.Int).SetBytes(out[:32]), new(big.Int).SetBytes(out[32:]))
				tx, ty = secp256k1.S256().ScalarMult(pkX, pkY, big.NewInt(101).Bytes())
				fmt.Println("101*(pkX,pkY) in go:\t", tx, ty)
			*/

			/*input, err = abiObj.Pack("DistributePartOne")
			must(err)
			fmt.Println("start Distribute")
			//(DisH,DisCmX,DisCmY);
			dis, _, vmerr := evm.Call(contractRef, contractAddr, input, statedb.GetBalance(fromAddress).Uint64(), big.NewInt(0))
			must(vmerr)
			fmt.Println("get Distribution data")*/

			dis := getDistribute()

			/*
				input, err = abiObj.Pack("DistributePartTwo")
				must(err)
				//(DisTokX,DisTokY,DisDt);
				outputs2, _, vmerr := evm.Call(contractRef, contractAddr, input, statedb.GetBalance(fromAddress).Uint64(), big.NewInt(0))
				must(vmerr)*/
			/*
				input, err = abiObj.Pack("AudditionPartOne")
				must(err)
				//DisAggX,DisAggY,DeltaX,DeltaY,AComX,AComY,yX,yY,aa
				aud1, _, vmerr := evm.Call(contractRef, contractAddr, input, statedb.GetBalance(fromAddress).Uint64(), big.NewInt(0))
				must(vmerr)

				input, err = abiObj.Pack("AudditionPartTwo")
				must(err)
				//DisAggX,DisAggY,DeltaX,DeltaY,AComX,AComY,yX,yY,aa
				aud2, _, vmerr := evm.Call(contractRef, contractAddr, input, statedb.GetBalance(fromAddress).Uint64(), big.NewInt(0))
				must(vmerr)
				fmt.Println("get Auddition data")*/

			aud1 := getAudition(true)
			aud2 := getAudition(false)
			clm := getClm()

			//for i := int64(0); i < l; i++ {
			sd := new(SendData)
			sd.listLength = l
			sd.dis = &dis
			sd.aud1 = &aud1
			sd.aud2 = &aud2
			sd.clm = &clm
			sd_ch <- sd
			//}

			response_count = 0
			/*input, err = abiObj.Pack("Clear")
			evm.Call(contractRef, contractAddr, input, statedb.GetBalance(fromAddress).Uint64(), big.NewInt(0))
			must(err)*/

			clear()
		}
	}

}

func getVariables(statedb *state.StateDB, hash common.Address) {
	cb := func(key, value common.Hash) bool { //匿名函数
		fmt.Printf("key=%x,value=%x\n", key, value)
		return true
	}

	statedb.ForEachStorage(hash, cb)

}

func Print(outputs []byte, name string) {
	fmt.Printf("method=%s, output=%x\n", name, outputs)
}

type ChainContext struct{}

func (cc ChainContext) GetHeader(hash common.Hash, number uint64) *types.Header {

	return &types.Header{
		// ParentHash: common.Hash{},
		// UncleHash:  common.Hash{},
		Coinbase: fromAddress,
		//	Root:        common.Hash{},
		//	TxHash:      common.Hash{},
		//	ReceiptHash: common.Hash{},
		//	Bloom:      types.BytesToBloom([]byte("duanbing")),
		Difficulty: big.NewInt(1),
		Number:     big.NewInt(1),
		GasLimit:   1000000,
		GasUsed:    0,
		Time:       big.NewInt(time.Now().Unix()),
		Extra:      nil,
		//MixDigest:  testHash,
		//Nonce:      types.EncodeNonce(1),
	}
}
