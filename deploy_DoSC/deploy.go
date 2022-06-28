package main

import (
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

const key = "{\"address\":\"a02e74362cfdf3a36fd34c4a393200ef0ba4edb8\",\"crypto\":{\"cipher\":\"aes-128-ctr\",\"ciphertext\":\"901552e4f14f82190fb301b00b411d580fde54b710fa5e75a887524611c120f8\",\"cipherparams\":{\"iv\":\"e529de78b511be52572c5447223f09b1\"},\"kdf\":\"scrypt\",\"kdfparams\":{\"dklen\":32,\"n\":262144,\"p\":1,\"r\":8,\"salt\":\"b0d8806f8dc69ac663227488677139fa21a13ab2927014d556a5a119138d3cb6\"},\"mac\":\"e9da517c5936aca72f84dc1c592ce39e55e2eb7a5de3f974330ac4f116cedca8\"},\"id\":\"37cd8427-74fe-4c86-931f-f588782b9358\",\"version\":3}"

func main() {

	//下面是部署合约的代码

	conn, err := ethclient.Dial("http://localhost:8545")
	if err != nil {

		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	auth, err := bind.NewTransactorWithChainID(strings.NewReader(key), "123456", big.NewInt(150))
	if err != nil {

		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	// Deploy a new awesome contract for the binding demo
	// address, tx, Lottery, err := DeployLottery(auth, conn)//  new(big.Int), "Contracts in Go!!!", 0, "Go!")
	address, tx, _, err := DeployDoSC(auth, conn) //, big.NewInt(1337), "Contracts in Go!!!", 0, "Go!")

	if err != nil {

		log.Fatalf("Failed to deploy new Lottery contract: %v", err)
	}
	fmt.Printf("Contract pending deploy: 0x%x\n", address)
	fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())

	// Don't even wait, check its presence in the local pending state
	time.Sleep(250 * time.Millisecond) // Allow it to be processed by the local node :P

	//下面是使用合约的代码

	// //1
	// conn, err := ethclient.Dial("http://localhost:8545")
	// if err != nil {
	// 	log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	// }

	// // 2
	// token, err := NewDoSC(common.HexToAddress("0xe75fdf2efa6f4ff37a4393fb74d2ecad81477150"), conn)
	// if err != nil {
	// 	log.Fatalf("Failed to instantiate a Token contract: %v", err)
	// }

	// // 3
	// auth, err := bind.NewTransactorWithChainID(strings.NewReader(key), "123456", big.NewInt(150))
	// if err != nil {
	// 	log.Fatalf("Failed to create authorized transactor: %v", err)
	// }

	// alloc := make(core.GenesisAlloc)
	// alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1337000000000)}
	// //sim := backends.NewSimulatedBackend(alloc, 100000000)

	// //监听
	// fmt.Println("hello")
	// listen, err := net.Listen("tcp", "192.168.1.104:8080")
	// fmt.Println("listen 8080 ")
	// if err != nil {
	// 	fmt.Println("err = ", err)
	// 	return
	// }

	// TEST1 := "a"
	// TEST2 := "a2"
	// tx0, err := token.CompareStrings(&bind.CallOpts{Pending: true}, TEST1, TEST2)
	// if err != nil {
	// 	log.Fatalf("Failed to request token transfer: %v", err)
	// }
	// fmt.Println("tx0", tx0)

	// fmt.Println(token.Len0(nil))
	// fmt.Println(token.Len1(nil))
	// //fmt.Printf("tx sent: %s \n", tx.Hash().Hex())

	// if err != nil {
	// 	log.Fatalf("Failed to request token transfer: %v", err)
	// }

	// defer listen.Close() //最后程序结束后才执行listen.Close( )进行关闭

	// // 死循环，阻塞等待用户链接
	// for {

	// 	conn, err := listen.Accept()
	// 	if err != nil {
	// 		fmt.Println("err = ", err)
	// 		return
	// 	} // 一有连接就发一个协程去处理
	// 	buf := make([]byte, 1024) //创建切片缓冲区，其大小为1024
	// 	n, err := conn.Read(buf)
	// 	fmt.Printf("buf: %v\n", buf)

	// 	if n != 480 {
	// 		fmt.Println("hi")
	// 		//return
	// 	}

	// 	if err != nil {
	// 		fmt.Println("err = ", err)
	// 		return
	// 	}

	// 	if buf[0] == 100 {

	// 		fmt.Println("now you are send message to Distribute...")
	// 		// 接受Dis信息
	// 		//Distribute(uint256 pkcoix,uint256 pkcoiy,bytes32 hclj,uint256 cmqx,uint256 cmqy,uint256 tokqx,uint256 tokqy,bytes32 msgh,uint8 v,bytes32 r,bytes32 s)
	// 		pkcoix := new(big.Int).SetBytes(buf[1:33])
	// 		pkcoiy := new(big.Int).SetBytes(buf[33:65])
	// 		var hclj [32]byte
	// 		for i := 65; i < 97; i++ {
	// 			hclj[i-65] = buf[i]
	// 		}
	// 		fmt.Printf("hclj: %v\n", hclj)
	// 		cmqx := new(big.Int).SetBytes(buf[97:129])
	// 		cmqy := new(big.Int).SetBytes(buf[129:161])
	// 		tokqx := new(big.Int).SetBytes(buf[161:193])
	// 		tokqy := new(big.Int).SetBytes(buf[193:225])

	// 		tx, err := token.Distribute(&bind.TransactOpts{
	// 			From:   auth.From,
	// 			Signer: auth.Signer,
	// 			Value:  nil,
	// 		}, pkcoix, pkcoiy, hclj, cmqx, cmqy, tokqx, tokqy)

	// 		if err != nil {
	// 			log.Fatalf("set contract: %v \n", err)
	// 		}

	// 		fmt.Println(token.Len0(nil))
	// 		fmt.Println(token.Len1(nil))
	// 		fmt.Printf("tx sent: %s \n", tx.Hash().Hex())

	// 		if err != nil {
	// 			log.Fatalf("Failed to request token transfer: %v", err)
	// 		}

	// 		// tx2, err := token.PrintLenOfDis(&bind.TransactOpts{
	// 		// 	From:   auth.From,
	// 		// 	Signer: auth.Signer,
	// 		// 	Value:  nil,
	// 		// }, big.NewInt(1))

	// 		// fmt.Println(token.Len0(nil))
	// 		// fmt.Println(token.Len1(nil))
	// 		// fmt.Printf("tx sent: %s \n", tx2.Hash().Hex())
	// 		// if err != nil {
	// 		// 	log.Fatalf("Failed to request token transfer: %v", err)
	// 		// }

	// 	}

	// 	if buf[0] == 101 {

	// 		// 接受Store信息
	// 		//Distribute(uint256 pkcoix,uint256 pkcoiy,bytes32 hclj,uint256 cmqx,uint256 cmqy,uint256 tokqx,uint256 tokqy,bytes32 msgh,uint8 v,bytes32 r,bytes32 s)
	// 		pkcoix := new(big.Int).SetBytes(buf[1:33])
	// 		pkcoiy := new(big.Int).SetBytes(buf[33:65])
	// 		var hclj [32]byte
	// 		for i := 65; i < 397; i++ {
	// 			hclj[i] = buf[i]
	// 		}

	// 		cmqx := new(big.Int).SetBytes(buf[97:129])
	// 		cmqy := new(big.Int).SetBytes(buf[129:161])
	// 		tokqx := new(big.Int).SetBytes(buf[161:193])
	// 		tokqy := new(big.Int).SetBytes(buf[193:225])

	// 		tx, err := token.Distribute(&bind.TransactOpts{
	// 			From:   auth.From,
	// 			Signer: auth.Signer,
	// 			Value:  nil,
	// 		}, pkcoix, pkcoiy, hclj, cmqx, cmqy, tokqx, tokqy)

	// 		if err != nil {
	// 			log.Fatalf("set contract: %v \n", err)
	// 		}
	// 		fmt.Println(token.Len0(nil))
	// 		fmt.Println(token.Len1(nil))
	// 		fmt.Printf("tx sent: %s \n", tx.Hash().Hex())

	// 		if err != nil {
	// 			log.Fatalf("Failed to request token transfer: %v", err)
	// 		}

	// 	}

	// }

}
