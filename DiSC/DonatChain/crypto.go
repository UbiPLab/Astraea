package main

import (
	//"bytes"
	//"crypto/rand"
	"crypto/sha256"
	"errors"
	"math/rand"

	//"encoding/binary"
	//"encoding/hex"
	"fmt"
	//"io"
	"math/big"
	"net"
	_ "net/http/pprof"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	//"github.com/tjfoc/gmsm/sm2"
)

var (
	zero  = big.NewInt(0)
	one   = big.NewInt(1)
	two   = big.NewInt(2)
	three = big.NewInt(3)
	four  = big.NewInt(4)
	w     = 16 // at most 8 byte as a group
	l     = 32 // byteLen of random and priKey
	k     = 16 // offset
	bigk  = big.NewInt(16)
	wg    sync.WaitGroup //用于控制等待线程结束
	//bigr1 *big.Int
	//bigr2 *big.Int
)

// raw data format: plainTextLen | strLen | encryptedData

type DecryptedHeader struct {
	data         []byte
	plainTextLen int32
	strLen       int32
}

// PKCS7 standard
func Padding(msg []byte) []byte {
	fmt.Println("\t\t进行字节填充")
	//fmt.Println("origin len:", len(msg))
	var Len = 0
	if len(msg)%w == 0 {
		Len = w
	} else {
		Len = w - len(msg)%w
	}
	res := make([]byte, Len+len(msg))
	copy(res, msg)
	for i := 0; i < Len; i++ {
		res[i+len(msg)] = byte(Len)
	}
	//fmt.Println("after len:", len(res))
	//fmt.Println("填充结果：\t", res)
	return res
}

func L2B(b []byte) ([]byte, error) {
	if len(b) != 32 {
		return nil, errors.New("not 32 byte")
	}
	var (
		l    = b[0]
		r    = b[31]
		lind = 0
		rind = 0
	)
	for lind < rind {
		l = b[lind]
		r = b[rind]
		b[lind] = r
		b[rind] = l
		lind++
		rind--
	}
	return b, nil
}

// tested
// only big.Exp()
func ExpMod(a, e, p *big.Int) *big.Int { // x^y mod p
	res := big.NewInt(1)
	tmp := big.NewInt(0)
	x := new(big.Int).Set(a)
	y := new(big.Int).Set(e)
	for y.Cmp(zero) != 0 {
		if tmp.And(y, one).Cmp(one) == 0 { // ((y&1)==1)==0
			res.Mul(res, x).Mod(res, p)
			y.Sub(y, one)
		} else {
			x.Mul(x, x).Mod(x, p)
			y.Rsh(y, 1) // y>>1
		}
	}
	return res
}

// return [(x1,y1), (x2,y2)...]
func Encode(mm []byte) []*big.Int {
	fmt.Println("\t\t点的嵌入")
	msg := Padding(mm)
	//fmt.Println("origin len:", len(msg))
	res := make([]*big.Int, 2*len(msg)/w)
	big7 := big.NewInt(7)
	P := secp256k1.S256().P
	Psub1 := new(big.Int).Sub(P, one)              // P-1
	half_Psub1 := new(big.Int).Div(Psub1, two)     // (P-1)/2
	Padd1 := new(big.Int).Add(P, one)              // P+1
	quarter_Padd1 := new(big.Int).Div(Padd1, four) // (P+1)/4
	for i := 0; i < len(msg); i += w {
		m := new(big.Int).SetBytes(msg[i : i+w])
		x := new(big.Int).Mul(m, bigk) //x.Mod(x, P)// no use in theory
		for j := 0; j < k; j++ {
			xxx := ExpMod(x, three, P)        // x^3
			yy := new(big.Int).Add(xxx, big7) // y^2 = x^3 + 7 (mod p)
			Legendre := ExpMod(yy, half_Psub1, P)
			if Legendre.Cmp(one) == 0 {
				y := ExpMod(yy, quarter_Padd1, P) // P == 3 mod(4)
				res[2*i/w] = x
				res[2*i/w+1] = y
				break
			}
			x.Add(x, one) // Xj = k*X+j
		}
	}
	//fmt.Println("after len:", len(res))
	return res
}

// sgx上怎么获得安全的随机数呢
// 采用多线程对同一块数据进行脏写
func DirtyWrite1(p *[]byte, r *rand.Rand) { //传入数组指针
	defer wg.Done()
	arr := *p
	for i := 0; i < len(arr); i++ {
		arr[i] = byte(r.Int())
	}
}

func DirtyWrite2(p *[]byte, r *rand.Rand) {
	defer wg.Done()
	arr := *p
	pos := len(arr) / 4
	for i := 0; i < len(arr); i++ {
		arr[pos%len(arr)] = byte(r.Int())
		pos++
	}
}

func DirtyWrite3(p *[]byte, r *rand.Rand) {
	defer wg.Done()
	arr := *p
	pos := len(arr) / 2
	for i := 0; i < len(arr); i++ {
		arr[pos%len(arr)] = byte(r.Int())
		pos++
	}
}

func DirtyWrite4(p *[]byte, r *rand.Rand) {
	defer wg.Done()
	arr := *p
	pos := 3 * (len(arr) / 4)
	for i := 0; i < len(arr); i++ {
		arr[pos%len(arr)] = byte(r.Int())
		pos++
	}
}

func GetRand(lll uint) []byte {
	res := make([]byte, lll)
	source := rand.NewSource(time.Now().UnixNano())
	//DirtyWrite1(&res, rand.New(source))
	r := rand.New((source))
	for i := 0; i < len(res); i++ {
		res[i] = byte(r.Int())
	}
	/*
		wg.Add(16) //设置这么多线程是为了触发条件竞争
		go DirtyWrite1(&res, rand.New(source))
		go DirtyWrite1(&res, rand.New(source))
		go DirtyWrite1(&res, rand.New(source))
		go DirtyWrite1(&res, rand.New(source))
		go DirtyWrite2(&res, rand.New(source))
		go DirtyWrite2(&res, rand.New(source))
		go DirtyWrite2(&res, rand.New(source))
		go DirtyWrite2(&res, rand.New(source))
		go DirtyWrite3(&res, rand.New(source))
		go DirtyWrite3(&res, rand.New(source))
		go DirtyWrite3(&res, rand.New(source))
		go DirtyWrite3(&res, rand.New(source))
		go DirtyWrite4(&res, rand.New(source))
		go DirtyWrite4(&res, rand.New(source))
		go DirtyWrite4(&res, rand.New(source))
		go DirtyWrite4(&res, rand.New(source))
		wg.Wait() //等待所有线程结束再返回
	*/
	return res
}

// Q = dG
// C = {rG, M+rQ} = {C1, C2}
// M = C2 - dC1 = M + rQ - d(rG) = M + r(dG) - d(rG) = M

// pk = h^r2
// h = g^r1
// sk = r1+r2 // cannot understand, but practise correctly
// pk = g^sk
func GenKey() (priByte []byte, pubByte []byte) {
	fmt.Println("\t\t初始化密钥对")
	// pubKey = h ^ priKey
	priKey := GetRand(2 * uint(l)) // priKey = [r1, r2]
	r1 := priKey[:l]
	//fmt.Println("r1:\t", r1)
	bigr1 := new(big.Int).SetBytes(r1)
	//fmt.Println("bigr1:\t", bigr1)
	Hx, Hy := secp256k1.S256().ScalarBaseMult(r1)
	r2 := priKey[l:]
	rand.Read(r2)
	//fmt.Println("r2:\t", r2)
	bigr2 := new(big.Int).SetBytes(r2)
	//fmt.Println("bigr2:\t", bigr2)
	bigr := new(big.Int).Mul(bigr1, bigr2)
	//P := secp256k1.S256().P
	bigr.Mod(bigr, secp256k1.S256().N)
	//fmt.Println("bigr=bigr1*bigr2 (mod phip):\t", bigr)
	//fmt.Println("priKey:\t", bigr)
	//fmt.Println("priKey:", priKey)
	//bigrb := bigr.Bytes()
	//pubX, pubY := secp256k1.S256().ScalarBaseMult(bigrb)

	//fmt.Println("pubX:", pubX, "\t", "pubY:", pubY)

	//priKey := new(big.Int).Add(bigr1, bigr2)//why add? not mul?
	//priByte := priKey.Bytes()
	pubKeyX, pubKeyY := secp256k1.S256().ScalarMult(Hx, Hy, r2) // be careful
	//pubX, pubY := secp256k1.S256().ScalarBaseMult(priByte)
	//
	//fmt.Println("pubKeyX:\t", pubKeyX, "\n", "pubKeyY:\t", pubKeyY)
	//fmt.Println("pubX:", pubX, "\t", "pubY:", pubY)
	return priKey, secp256k1.S256().Marshal(pubKeyX, pubKeyY) // []byte,big-endian
}

func Encrypt(pubKey []byte, msgg []byte) []byte {
	fmt.Println("\t\t进行加密")
	points := Encode(msgg)
	cipherText := make([]byte, 130*len(points)/2)
	pubKeyX, pubKeyY := secp256k1.S256().Unmarshal(pubKey)
	// C = (C1, C2) = (rG, M+rQ) ; Q = dG
	for i := 0; i < len(points); i += 2 {
		//sgx上不需要作加密，所以加密使用的随机数不需要那么安全
		r := make([]byte, l)
		rand.Read(r)
		C1x, C1y := secp256k1.S256().ScalarBaseMult(r) // C1 = rG
		Mx := points[i]
		My := points[i+1]
		rQx, rQy := secp256k1.S256().ScalarMult(pubKeyX, pubKeyY, r)
		C2x, C2y := secp256k1.S256().Add(Mx, My, rQx, rQy) // C2 = M + rQ
		C1 := secp256k1.S256().Marshal(C1x, C1y)
		C2 := secp256k1.S256().Marshal(C2x, C2y)
		j := 130 * (i / 2)
		copy(cipherText[j:j+65], C1)
		copy(cipherText[j+65:j+130], C2)
	}
	//fmt.Println("密文长度:\t", len(cipherText))
	//fmt.Println("密文内容", cipherText)
	return cipherText
}

func Decrypt(priKey []byte, cipher []byte) []byte {
	fmt.Println("\t\t进行解密")
	priKey1 := priKey[:l]
	priKey2 := priKey[l:]
	plainText := make([]byte, w*len(cipher)/130)
	P := secp256k1.S256().P
	ll := 0
	for i := 0; i < len(cipher); i += 130 {
		C1x, C1y := secp256k1.S256().Unmarshal(cipher[i : i+65])
		C2x, C2y := secp256k1.S256().Unmarshal(cipher[i+65 : i+130])
		tmpx, tmpy := secp256k1.S256().ScalarMult(C1x, C1y, priKey1)
		kC1x, kC1y := secp256k1.S256().ScalarMult(tmpx, tmpy, priKey2)
		Mx, _ := secp256k1.S256().Add(C2x, C2y, kC1x, kC1y.Sub(P, kC1y)) // X = C2 - dC1
		Mbig := new(big.Int).Quo(Mx, bigk)                               // X = kM + i mod p
		tmp := Mbig.Bytes()                                              // w byte by default
		if len(tmp) < w {
			ll = len(tmp)
		} else {
			ll = w
		}
		t := w * (i / 130)
		copy(plainText[t:w+t], tmp[0:ll])
	}
	//cut
	ll = len(plainText)
	padLen := plainText[ll-1]
	res := plainText[:(ll - int(padLen))]
	//fmt.Println("解密长度:", len(res))
	//fmt.Println("解密内容:", string(res))
	return res
}

//returns 64bytes (r,s)
func Sign(msg []byte, sk []byte) (*big.Int, *big.Int) {
	fmt.Println("\t\t进行签名")
	//r!=0&&s!=0
	hasher := sha256.New()
	hasher.Write(msg)
	h := hasher.Sum(nil) //
	e := new(big.Int).SetBytes(h)
	r := big.NewInt(0)
	s := big.NewInt(0)
	for r.Cmp(zero) == 0 || s.Cmp(zero) == 0 {
		kb := make([]byte, 32)
		rand.Read(kb)
		k := new(big.Int).SetBytes(kb)
		KX, _ := secp256k1.S256().ScalarBaseMult(kb)
		r.Mod(KX, secp256k1.S256().N)
		s.Mul(r, new(big.Int).SetBytes(sk)).Mod(s, secp256k1.S256().P).Add(s, e).Mod(s, secp256k1.S256().P)
		kinv := new(big.Int).ModInverse(k, secp256k1.S256().P)
		s.Mod(s, kinv).Mod(s, secp256k1.S256().P)
	}
	return r, s
}

func Verify(msg []byte, sig []byte, pkX *big.Int, pkY *big.Int) (bool, error) {
	fmt.Println("\t\t验证签名")
	if len(sig) != 64 {
		return false, errors.New("(r,s) must be 64 bytes!")
	}
	hasher := sha256.New()
	hasher.Write(msg)
	h := hasher.Sum(nil) //
	e := new(big.Int).SetBytes(h)
	r := new(big.Int).SetBytes(sig[0:32])
	s := new(big.Int).SetBytes(sig[32:])
	sinv := new(big.Int).ModInverse(s, secp256k1.S256().P)
	u1 := new(big.Int).Mul(sinv, e)
	u1.Mod(u1, secp256k1.S256().P)
	u2 := new(big.Int).Mul(sinv, r)
	u2.Mod(u2, secp256k1.S256().P)
	tmp1X, tmp1Y := secp256k1.S256().ScalarMult(secp256k1.S256().Gx, secp256k1.S256().Gy, u1.Bytes())
	tmp2X, tmp2Y := secp256k1.S256().ScalarMult(pkX, pkY, u2.Bytes())
	PX, _ := secp256k1.S256().Add(tmp1X, tmp1Y, tmp2X, tmp2Y)
	return r.Cmp(PX) == 0, nil
}

func test0() {
	fmt.Println("测试sha256哈希")
	h := sha256.New()
	str := "this is a sha256 test"
	fmt.Println("原消息:\t", str)
	b := []byte("this is a sha256 test")
	h.Write(b)
	hash := h.Sum(nil)
	fmt.Println("哈希值:\t", hash)
}

func test1() {
	fmt.Println("测试加解密.")
	priKey, pubKey := GenKey()
	msg := []byte("this is a secp256k1 test")
	fmt.Println("消息:\t", "this is a secp256k1 test")
	cipherpoint := Encrypt(pubKey, msg)
	decrypted := Decrypt(priKey, cipherpoint)
	decrypted[0] = 0
}

func test2() {
	priKey := []byte{85, 25, 66, 196, 117, 47, 108, 143, 148, 226, 45, 58, 1, 188, 171, 15, 102, 202, 206, 208, 137, 254, 94, 231, 110, 205, 249, 26, 33, 186, 128, 49, 149, 243, 208, 156, 161, 174, 175, 213, 169, 205, 42, 175, 198, 226, 64, 214, 122, 0, 136, 167, 162, 200, 111, 185, 27, 237, 118, 85, 53, 164, 25, 252}
	pubKey := []byte{4, 181, 92, 42, 120, 67, 72, 65, 196, 19, 78, 220, 230, 73, 120, 180, 68, 16, 11, 140, 99, 20, 137, 136, 113, 82, 40, 4, 105, 222, 231, 215, 14, 90, 98, 92, 119, 107, 202, 67, 63, 2, 8, 1, 213, 203, 70, 157, 28, 76, 9, 39, 99, 164, 102, 60, 110, 121, 185, 201, 240, 4, 138, 77, 12}
	msg := []byte{0xff, 0xad, 0x23, 0x89}

	cipherpoint := Encrypt(pubKey, msg)
	decrypted := Decrypt(priKey, cipherpoint)

	//监听
	listen, err := net.Listen("tcp", "localhost: 8001")
	//listen, err := net.Listen("tcp", "localhost: 20510")
	fmt.Println("listen 8001")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	defer listen.Close() //最后程序结束后才执行listen.Close( )进行关闭
	conn, err := listen.Accept()
	fmt.Println("conn")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer conn.Close()

	buf := make([]byte, 1024) //创建切片缓冲区，其大小为1024
	n, err1 := conn.Read(buf)
	if err1 != nil {
		fmt.Println("err = ", err)
		return
	}

	fmt.Println("buf = ", buf[:n])
	fmt.Println("n:\t", n)
	fmt.Println("start decrypt")
	decrypted = Decrypt(priKey, buf[:n])
	fmt.Println(string(decrypted))
}

func test3() {
	priKey, pubKey := GenKey()
	msg := []byte("khgosghwogjogfjawgwggohgowhgoghowghwoghwoghwaolhjworjfwor")
	sigr, sigs := Sign(msg, priKey)
	fmt.Println(sigr)
	fmt.Println(sigs)
	pkX, pkY := secp256k1.S256().Unmarshal(pubKey)
	sig := make([]byte, 32*2)
	copy(sig[:32], sigr.Bytes())
	copy(sig[32:], sigs.Bytes())
	fmt.Println(Verify(msg, sig, pkX, pkY))

}

func test4() {
	fmt.Println("进行加密运算的性能测试")
	fmt.Println("生成1024字节消息")
	msg := make([]byte, 1024)
	rand.Read(msg)
	fmt.Println("生成临时密钥对")
	priKey1, pubKey1 := GenKey()
	encrypted := make([]byte, 0)
	decrypted := make([]byte, 0)
	start := time.Now().UnixNano()
	end := time.Now().UnixNano()
	var a byte
	for i := 0; i < 10; i++ {
		fmt.Println("开始加密，当前时间戳为:\t", start, " 纳秒")
		//go func() {
		//log.Println(http.ListenAndServe("0.0.0.0:10000", nil))
		//}()
		encrypted = Encrypt(pubKey1, msg)
		end = time.Now().UnixNano()
		fmt.Println("加密完成，当前时间戳为:\t", end, " 纳秒")
		fmt.Println("耗时约为:\t", (end-start)/1000000, " 毫秒")
		a = encrypted[0]
	}

	fmt.Println("\n\n\n")
	for i := 0; i < 10; i++ {
		fmt.Println("开始解密，当前时间戳为:\t", start, " 纳秒")
		//go func() {
		//log.Println(http.ListenAndServe("0.0.0.0:10000", nil))
		//}()
		decrypted = Decrypt(priKey1, encrypted)
		end = time.Now().UnixNano()
		fmt.Println("解密完成，当前时间戳为:\t", end, " 纳秒")
		fmt.Println("耗时约为:\t", (end-start)/1000000, " 毫秒")
		a = decrypted[0]
	}
	fmt.Println("\n\n\n\n\n", a)
}

func main1() {
	test3()
	//test2()

}
