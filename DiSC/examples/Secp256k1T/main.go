package main

import (
	"crypto/rand"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"time"
)

func main() {
	start := time.Now() // 获取当前时间
	// 被测代码
	s256k1 :=secp256k1.S256()
	Gx:= s256k1.Gx
	Gy:= s256k1.Gy
	b := make([]byte, 32)
	_, _ = rand.Reader.Read(b)
	Hx,Hy:= s256k1.ScalarMult(Gx,Gy,b)
	fmt.Println(Gx,Gy)
	fmt.Println(Hx,Hy)
	//Ecc Mul
	Tx,Ty:= s256k1.Add(Gx,Gy,Hx,Hy)
	fmt.Println(Tx,Ty)
	//Ecc Add
	message:="This is a message for test"
	var data = []byte(message)
	hash:=crypto.Keccak256(data)
	fmt.Println(common.Bytes2Hex(hash))
	cost := time.Since(start)
	fmt.Println(cost)
}
