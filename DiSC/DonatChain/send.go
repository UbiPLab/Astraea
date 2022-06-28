package main

import (
	"fmt"
	"math/big"
	"net"
	"time"
)

func recv(sd_ch chan *SendData) {
	//tmp1 := make([]byte, 32*9+1)
	//tmp1[0] = 100
	for {

		sd := <-sd_ch
		fmt.Println("模拟dosc接收 ########################################################")
		//var PoD [32 * 7]byte
		//var PoA [32 * 7]byte
		PoD := make([]byte, 7*32)
		PoA := make([]byte, 7*32)
		for i := int64(0); i < sd.listLength*8*32; i += 8 * 32 {
			//copy(tmp1[0:32], pubKeyX)
			//copy(tmp1[32:64], pubKeyY)
			fmt.Println("收到 Dis########################################################")

			fmt.Println("pkcoix:\t", (*sd.dis)[i+32*0:i+32*1])
			fmt.Println("pkcoix:\t", new(big.Int).SetBytes((*sd.dis)[i+32*0:i+32*1]))
			fmt.Println("pkcoiy:\t", (*sd.dis)[i+32*1:i+32*2])
			fmt.Println("pkcoiy:\t", new(big.Int).SetBytes((*sd.dis)[i+32*1:i+32*2]))
			fmt.Println("DisH:\t", (*sd.dis)[i+32*2:i+32*3])
			fmt.Println("DisH:\t", new(big.Int).SetBytes((*sd.dis)[i+32*2:i+32*3]))
			fmt.Println("DisComX:\t", (*sd.dis)[i+32*3:i+32*4])
			fmt.Println("DisComX:\t", new(big.Int).SetBytes((*sd.dis)[i+32*3:i+32*4]))
			fmt.Println("DisComY:\t", (*sd.dis)[i+32*4:i+32*5])
			fmt.Println("DisComY:\t", new(big.Int).SetBytes((*sd.dis)[i+32*4:i+32*5]))
			fmt.Println("DisTokX:\t", (*sd.dis)[i+32*5:i+32*6])
			fmt.Println("DisTokX:\t", new(big.Int).SetBytes((*sd.dis)[i+32*5:i+32*6]))
			fmt.Println("DisTokY:\t", (*sd.dis)[i+32*6:i+32*7])
			fmt.Println("DisTokY:\t", new(big.Int).SetBytes((*sd.dis)[i+32*6:i+32*7]))
			fmt.Println("DisDt:\t", (*sd.dis)[i+32*7:i+32*8])
			fmt.Println("DisDt:\t", new(big.Int).SetBytes((*sd.dis)[i+32*7:i+32*8]))

		}

		fmt.Println("进行PoD验证#######################################################")
		fmt.Println("DisAggX:\t", (*sd.aud1)[32*0:32*1])
		fmt.Println("DisAggX:\t", new(big.Int).SetBytes((*sd.aud1)[32*0:32*1]))
		fmt.Println("DisAggY:\t", (*sd.aud1)[32*1:32*2])
		fmt.Println("DisAggY:\t", new(big.Int).SetBytes((*sd.aud1)[32*1:32*2]))
		fmt.Println("ACominX:\t", (*sd.aud1)[32*2:32*3])
		fmt.Println("ACominX:\t", new(big.Int).SetBytes((*sd.aud1)[32*2:32*3]))
		fmt.Println("ACominY:\t", (*sd.aud1)[32*3:32*4])
		fmt.Println("ACominY:\t", new(big.Int).SetBytes((*sd.aud1)[32*3:32*4]))
		fmt.Println("yX:\t", (*sd.aud1)[32*4:32*5])
		fmt.Println("yX:\t", new(big.Int).SetBytes((*sd.aud1)[32*4:32*5]))
		fmt.Println("yY:\t", (*sd.aud1)[32*5:32*6])
		fmt.Println("yY:\t", new(big.Int).SetBytes((*sd.aud1)[32*5:32*6]))
		fmt.Println("deltaX:\t", (*sd.aud1)[32*6:32*7])
		fmt.Println("deltaX:\t", new(big.Int).SetBytes((*sd.aud1)[32*6:32*7]))
		fmt.Println("deltaY:\t", (*sd.aud1)[32*7:32*8])
		fmt.Println("deltaY:\t", new(big.Int).SetBytes((*sd.aud1)[32*7:32*8]))
		fmt.Println("aa:\t", (*sd.aud1)[32*8:32*9])
		fmt.Println("aa:\t", new(big.Int).SetBytes((*sd.aud1)[32*8:32*9]))
		copy(PoD, (*sd.aud1)[32*2:])

		if verify(PoD) {
			fmt.Println("PoD 验证成功")
		} else {
			fmt.Println("PoD 验证失败")
		}

		fmt.Println("进行PoA验证#######################################################")
		fmt.Println("DisAggX:\t", (*sd.aud2)[32*0:32*1])
		fmt.Println("DisAggX:\t", new(big.Int).SetBytes((*sd.aud2)[32*0:32*1]))
		fmt.Println("DisAggY:\t", (*sd.aud2)[32*1:32*2])
		fmt.Println("DisAggY:\t", new(big.Int).SetBytes((*sd.aud2)[32*1:32*2]))
		fmt.Println("ACominX:\t", (*sd.aud2)[32*2:32*3])
		fmt.Println("ACominX:\t", new(big.Int).SetBytes((*sd.aud2)[32*2:32*3]))
		fmt.Println("ACominY:\t", (*sd.aud2)[32*3:32*4])
		fmt.Println("ACominY:\t", new(big.Int).SetBytes((*sd.aud2)[32*3:32*4]))
		fmt.Println("yX:\t", (*sd.aud2)[32*4:32*5])
		fmt.Println("yX:\t", new(big.Int).SetBytes((*sd.aud2)[32*4:32*5]))
		fmt.Println("yY:\t", (*sd.aud2)[32*5:32*6])
		fmt.Println("yY:\t", new(big.Int).SetBytes((*sd.aud2)[32*5:32*6]))
		fmt.Println("deltaX:\t", (*sd.aud2)[32*6:32*7])
		fmt.Println("deltaX:\t", new(big.Int).SetBytes((*sd.aud2)[32*6:32*7]))
		fmt.Println("deltaY:\t", (*sd.aud2)[32*7:32*8])
		fmt.Println("deltaY:\t", new(big.Int).SetBytes((*sd.aud2)[32*7:32*8]))
		fmt.Println("aa:\t", (*sd.aud2)[32*8:32*9])
		fmt.Println("aa:\t", new(big.Int).SetBytes((*sd.aud2)[32*8:32*9]))
		copy(PoA, (*sd.aud2)[32*2:])

		if verify(PoA) {
			fmt.Println("PoA 验证成功")
		} else {
			fmt.Println("PoA 验证失败")
		}

		fmt.Println("\n\n\n")

		end = time.Now().UnixNano()
		fmt.Println("验证完毕，当前时间戳为:\t", end, " 纳秒")
		fmt.Println("耗时约为:\t", (end-start)/1000000, " 毫秒")
	}
}

//这个要大改特改
func send(sd_ch chan *SendData) {
	fmt.Println("转发")
	conn, err := net.Dial("tcp", "127.0.0.1:8002")
	if err != nil {
		fmt.Println("err :", err)
		return
	}
	defer conn.Close() // 关闭连接
	//dis发给dosc
	//aud发给web，外部来添加时间戳 time.Now().Unix() 要转32
	tmp1 := make([]byte, 32*8+1)

	// 因为还有 Sotre  Ship Audditon
	// 实现一个简单的协议，通过第一字节的内容区分所发数据
	// 所有Dis型数据都是固定长度
	for {
		sd := <-sd_ch
		//dis
		tmp1[0] = 100 // 如果发的是Dis，则第一字节为100
		for i := int64(0); i < sd.listLength; i++ {
			//论文(“Distribute”, pkcoiX，pkcoiY, H(clj), cm′j, tok′j, dt)
			copy(tmp1[1:], (*sd.dis)[i+32*0:i+32*8])
			_, err = conn.Write(tmp1)
			must(err)
		}
		//PoD
		tmp1[0] = 101
		copy(tmp1[1:32*9+1], (*sd.aud1)[0:32*9])
		_, err = conn.Write(tmp1)
		must(err)
		//PoA
		tmp1[0] = 102
		copy(tmp1[1:32*9+1], (*sd.aud2)[0:32*9])
		_, err = conn.Write(tmp1)
		must(err)
		//clm
		//refund
	}
}
