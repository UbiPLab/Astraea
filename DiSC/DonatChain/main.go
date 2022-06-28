package main

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"math/rand"
	"net"
	"time"

	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

var (
	priKey []byte
	rb1    []byte
	rb2    []byte
	bigr1  *big.Int
	bigr2  *big.Int
	bigr   *big.Int

	pubKey  []byte
	pubKeyX []byte
	pubKeyY []byte
	pkX     *big.Int
	pkY     *big.Int
	HX      *big.Int
	HY      *big.Int
	//crypto_ch   chan []byte // 将网络发来的信息通过管道传给解密协程
	response_ch chan []byte // 将解密后的结果通过管道传给evm
	sd_ch       chan *SendData
	start       int64
	end         int64
)

func ini() {
	fmt.Println("程序初始化")
	priKey, pubKey = GenKey()
	rb1 = make([]byte, 32)
	rb2 = make([]byte, 32)
	copy(rb1, priKey[:32])
	copy(rb2, priKey[32:])
	bigr1 = new(big.Int).SetBytes(rb1)
	bigr2 = new(big.Int).SetBytes(rb2)
	bigr = new(big.Int).Mul(bigr1, bigr2)
	pubKeyX = make([]byte, 32)
	pubKeyY = make([]byte, 32)
	copy(pubKeyX, pubKey[1:33])
	copy(pubKeyY, pubKey[33:])
	pkX, pkY = secp256k1.S256().Unmarshal(pubKey)
	HX, HY = secp256k1.S256().ScalarBaseMult(rb1)

	response_ch = make(chan []byte, response_MAX)
	sd_ch = make(chan *SendData, response_MAX)
	//evmInit()

	fmt.Println("初始化完毕")
}

//可能还有签名
func handleConn(conn net.Conn, sk []byte, output chan []byte) {
	fmt.Println("处理一次Cmj")
	buf := make([]byte, 5120) //创建切片缓冲区，其大小为1024
	n, err := conn.Read(buf)
	//因为填充规则是，即使已经到了16整倍数，仍然填充，所以会多一个
	if n != 130*(1+16*32/16) {
		fmt.Println("len mismatch")
		return
	}
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	//fmt.Println("buf = ", buf[:n])
	///fmt.Println("接收消息长度:\t", n)
	//fmt.Println("进行解密")
	decrypted := Decrypt(priKey, buf[:n])
	//fmt.Println("解密后长度:\t", len(decrypted))
	output <- decrypted
}

func MSG() {
	time.Sleep(time.Duration(5) * time.Second)

	tmp1 := make([]byte, 32*16)

	dnb := make([]byte, 8)
	rrand := make([]byte, 32)
	for i := 0; i < response_MAX; i++ {
		fmt.Println("客户端发送cm")
		conn, err := net.Dial("tcp", "127.0.0.1:8001")
		if err != nil {
			fmt.Println("err :", err)
			return
		}
		rand.Read(rrand)
		copy(tmp1[32*0:32*1], rrand) //uint256 ReqRi;
		copy(tmp1[32*1:32*2], rrand) //uint256 ReqCi;
		rand.Read(rrand)
		copy(tmp1[32*2:32*3], rrand) //uint256 ReqDni;
		fmt.Println("生成社区密钥对")
		skcom, pkcom := GenKey()
		rr := new(big.Int).SetBytes(skcom[:32])
		rr.Mod(rr, secp256k1.S256().N)
		rr.Mul(rr, new(big.Int).SetBytes(skcom[32:])).Mod(rr, secp256k1.S256().N)
		pkcomx, pkcomy := secp256k1.S256().Unmarshal(pkcom)
		fmt.Println("skcom@:\t", rr)
		fmt.Println("pkcomx@:\t", pkcomx)
		fmt.Println("pkcomy@:\t", pkcomy)
		copy(tmp1[32*3:32*4], pkcomx.Bytes()) //uint256 ReqPkcoiX;
		copy(tmp1[32*4:32*5], pkcomy.Bytes()) //uint256 ReqPkcoiY;

		rand.Read(rrand)
		copy(tmp1[32*5:32*6], rrand) //uint256 tp1;
		rand.Read(rrand)
		copy(tmp1[32*6:32*7], rrand) //uint256 lc;
		rand.Read(rrand)
		copy(tmp1[32*7:32*8], rrand) //uint256 idj1;
		rand.Read(rrand)
		copy(tmp1[32*8:32*9], rrand) //bytes32 clj;
		fmt.Println("cl@:\t", rrand)
		h := sha256.New()
		h.Write(rrand)
		hash := h.Sum(nil)
		fmt.Println("hash(cl):\t", hash)

		rand.Read(dnb)
		//dn := new(big.Int).SetBytes(dnb)
		dn := big.NewInt(3)
		//dn.Mod(dn, secp256k1.S256().P)
		copy(tmp1[32*9:32*10], dn.Bytes()) //uint256 dnj;
		fmt.Println("dn@:\t", dn)
		rand.Read(rrand)
		copy(tmp1[32*10:32*11], rrand) //rj
		fmt.Println("选取的rj:\t", new(big.Int).SetBytes(rrand))

		rb := make([]byte, 32)
		rand.Read(rb)
		r := new(big.Int).SetBytes(rb)
		fmt.Println("_rj:\t", r)
		cmx1, cmy1 := secp256k1.S256().ScalarBaseMult(dn.Bytes())
		cmx2, cmy2 := secp256k1.S256().ScalarMult(HX, HY, rb)
		cmx, cmy := secp256k1.S256().Add(cmx1, cmy1, cmx2, cmy2)
		//cm := secp256k1.S256().Marshal(cmx, cmy)
		fmt.Println("cmx@:\t", cmx)
		fmt.Println("cmy@:\t", cmy)
		copy(tmp1[32*11:32*12], cmx.Bytes()) //cmx
		copy(tmp1[32*12:32*13], cmy.Bytes()) //cmy

		tokx, toky := secp256k1.S256().ScalarMult(pkX, pkY, rb)
		fmt.Println("tokx@:\t", tokx)
		fmt.Println("toky@:\t", toky)
		//tok := secp256k1.S256().Marshal(tokx, toky)
		copy(tmp1[32*13:32*14], tokx.Bytes()) //tokx
		copy(tmp1[32*14:32*15], toky.Bytes()) //toky

		dt := time.Now().Unix()
		fmt.Println("dt:\t", dt)
		dtb := big.NewInt(dt).Bytes()
		copy(tmp1[32*15:32*16], dtb) //dt
		//fmt.Println("加密前:\t", tmp1)
		//加密
		tmp2 := Encrypt(pubKey, tmp1)
		//fmt.Println("encrypt len:\t", len(tmp2))

		_, err = conn.Write(tmp2)
		must(err)
		conn.Close() // 关闭连接
		fmt.Println("发送一次CMj")
		time.Sleep(time.Duration(1) * time.Second)
	}

}

func main() {
	//start = time.Now().UnixNano()
	//fmt.Println("开始流程测试，当前时间戳为:\t", start, " 纳秒")

	//test4()
	//return
	fmt.Println("开始main()")
	ini()

	//监听
	//也许会有更好的I/O模型
	listen, err := net.Listen("tcp", "127.0.0.1:8001")
	fmt.Println("监听 8001")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	defer listen.Close() //最后程序结束后才执行listen.Close( )进行关闭
	// 必须先启动evm
	go Evm(response_ch, sd_ch)
	//go send(sd_ch)
	go recv(sd_ch)
	//go MSG() //应该单独拎出来

	//阻塞等待用户链接
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("err = ", err)
			return
		} // 一有连接就发一个协程去处理
		go handleConn(conn, priKey, response_ch)
		//recv(sd_ch)
	}

	//test3()
}
