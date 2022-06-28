package main

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"math/rand"
	"time"

	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

var (
	sum     *big.Int
	pkcoix  []*big.Int
	pkcoiy  []*big.Int
	DisH    [][]byte
	DisCmX  []*big.Int
	DisCmY  []*big.Int
	DisTokX []*big.Int
	DisTokY []*big.Int
	//AudCmX  []*big.Int
	//AudCmY  []*big.Int
	//AudTokX []*big.Int
	//AudTokY []*big.Int
	DisDt [][]byte
	clm   [][]byte
	//rh       hash.Hash
	//ResponseStruct []byte
	ATokoutX *big.Int
	ATokoutY *big.Int
	AComoutX *big.Int
	AComoutY *big.Int
	ATokinX  *big.Int
	ATokinY  *big.Int
	ACominX  *big.Int
	ACominY  *big.Int
	AggX     *big.Int
	AggY     *big.Int
)

func reinit() {
	sum = big.NewInt(0)
	pkcoix = make([]*big.Int, 0)
	pkcoiy = make([]*big.Int, 0)
	DisH = make([][]byte, 0)
	DisCmX = make([]*big.Int, 0)
	DisCmY = make([]*big.Int, 0)
	DisTokX = make([]*big.Int, 0)
	DisTokY = make([]*big.Int, 0)
	clm = make([][]byte, 0)
	//AudCmX = make([]*big.Int, 0)
	//AudCmY = make([]*big.Int, 0)
	//AudTokX = make([]*big.Int, 0)
	//AudTokY = make([]*big.Int, 0)
	//rh = sha256.New()
	DisDt = make([][]byte, 0)
	ATokinX = big.NewInt(0)
	ATokinY = big.NewInt(0)
	ACominX = big.NewInt(0)
	ACominY = big.NewInt(0)
	ATokoutX = big.NewInt(0)
	ATokoutY = big.NewInt(0)
	AComoutX = big.NewInt(0)
	AComoutY = big.NewInt(0)
	AggX = big.NewInt(0)
	AggY = big.NewInt(0)
}

func clear() {
	fmt.Println("\t\t\t\t进行清理")
	pkcoix = nil
	pkcoiy = nil
	DisH = nil
	DisCmX = nil
	DisCmY = nil
	DisTokX = nil
	DisTokY = nil
	DisDt = nil
	sum.Set(zero)
}

/*
	struct parameters {
    uint256 ReqRi;
	uint256 ReqCi;
	//uint256 ReqPkcoiX;
    //uint256 ReqPkcoiY;
	uint256 ReqCi;
	uint256 tp1;
	uint256 lc;
	uint256 idj1;
	uint256 dnj;
	bytes32 clj;
    //这是捐赠者生成的
	uint256 comx;//pkx
    uint256 comy;//pky
	uint256 tokx;//
    uint256 toky;//
	uint256 rj;
	uint256 dt;


}
*/

func Response(buf []*big.Int) {
	fmt.Println("\t\t\t\t调用Response")
	pkcoix = append(pkcoix, buf[3])
	pkcoiy = append(pkcoiy, buf[4])
	rh := sha256.New()
	rh.Write(buf[8].Bytes())
	pkcoi := secp256k1.S256().Marshal(buf[3], buf[4])
	clmb := Encrypt(pkcoi, buf[8].Bytes())
	clm = append(clm, clmb)

	DisH = append(DisH, rh.Sum(nil)) //hash
	//DisCmX = append(DisCmX, buf[10])
	//DisCmY = append(DisCmY, buf[12])
	//DisTokX = append(DisTokX, buf[13])
	//DisTokY = append(DisTokY, buf[14])
	dt := time.Now().Unix()
	dtb := big.NewInt(dt).Bytes()
	DisDt = append(DisDt, dtb)
	//sum	can't handle scalars > 256 bits ???
	//fmt.Println("dn#:\t", buf[9])
	//fmt.Println("dn#:\t", buf[9].Bytes())
	sum.Add(buf[9], sum) //.Mod(sum, secp256k1.S256().P) //?
	fmt.Println("\t\t\t\tsum:\t", sum)
	//com
	//fmt.Println("cmx#:\t", buf[11])
	//fmt.Println("cmy#:\t", buf[12])
	ACominX, ACominY = secp256k1.S256().Add(ACominX, ACominY, buf[11], buf[12])
	fmt.Println("\t\t\t\tACominX:\t", ACominX)
	fmt.Println("\t\t\t\tACominY:\t", ACominY)
	cm1x, cm1y := secp256k1.S256().ScalarBaseMult(buf[9].Bytes())      //g^dnj
	cm2x, cm2y := secp256k1.S256().ScalarMult(HX, HY, buf[10].Bytes()) //h^rj
	cmx, cmy := secp256k1.S256().Add(cm1x, cm1y, cm2x, cm2y)           //(g^dnj)(r^rj)
	DisCmX = append(DisCmX, cmx)
	DisCmY = append(DisCmY, cmy)
	AComoutX, AComoutY = secp256k1.S256().Add(AComoutX, AComoutY, cmx, cmy)
	fmt.Println("\t\t\t\tAComoutX:\t", AComoutX)
	fmt.Println("\t\t\t\tAComoutY:\t", AComoutY)

	//tok
	//fmt.Println("tokx#:\t", buf[13])
	//fmt.Println("toky#:\t", buf[14])
	ATokinX, ATokinY = secp256k1.S256().Add(ATokinX, ATokinY, buf[13], buf[14])
	fmt.Println("\t\t\t\tATokinX:\t", ATokinX)
	fmt.Println("\t\t\t\tATokinY:\t", ATokinY)
	tokx, toky := secp256k1.S256().ScalarMult(pkX, pkY, buf[10].Bytes())
	DisTokX = append(DisTokX, tokx)
	DisTokY = append(DisTokY, toky)
	ATokoutX, ATokoutY = secp256k1.S256().Add(ATokoutX, ATokoutY, tokx, toky)
	fmt.Println("\t\t\t\tATokoutX:\t", ATokoutX)
	fmt.Println("\t\t\t\tATokoutY:\t", ATokoutY)
}

func getDistribute() []byte { //([][]byte, []*big.Int, []*big.Int, []*big.Int, []*big.Int, [][]byte) {
	fmt.Println("\t\t\t\t调用Distribute")
	AggX, AggY = secp256k1.S256().ScalarBaseMult(sum.Bytes())
	fmt.Println("\t\t\t\tAggX:\t", AggX)
	fmt.Println("\t\t\t\tAggY:\t", AggY)
	llen := len(DisCmX)
	res := make([]byte, llen*8*32)
	for i := 0; i < len(res); i += 8 * 32 {
		copy(res[i+32*0:i+32*1], pkcoix[i/(32*8)].Bytes())
		copy(res[i+32*1:i+32*2], pkcoiy[i/(32*8)].Bytes())
		copy(res[i+32*2:i+32*3], DisH[i/(32*8)])
		copy(res[i+32*3:i+32*4], DisCmX[i/(32*8)].Bytes())
		copy(res[i+32*4:i+32*5], DisCmY[i/(32*8)].Bytes())
		copy(res[i+32*5:i+32*6], DisTokX[i/(32*8)].Bytes())
		copy(res[i+32*6:i+32*7], DisTokY[i/(32*8)].Bytes())
		copy(res[i+32*7:i+32*8], DisDt[i/(32*8)])
	}
	return res //pkcoix,pkcoiy,DisH, DisCmX, DisCmY, DisTokX, DisTokY, DisDt
}

func getAudition(boo bool) []byte { //(*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int) {
	var t1, t2, t3, t4, t5, t6, t7 *big.Int
	if boo {
		fmt.Println("\t\t\t\t生成PoD审计证明")
		t1, t2, t3, t4, t5, t6, t7 = genProof(ACominX, ACominY)
	} else {
		fmt.Println("\t\t\t\t生成PoA审计证明")
		t1, t2, t3, t4, t5, t6, t7 = genProof(AComoutX, AComoutY)
	}

	res := make([]byte, 32*9)
	copy(res[32*0:32*1], AggX.Bytes())
	copy(res[32*1:32*2], AggY.Bytes())
	copy(res[32*2:32*3], t1.Bytes())
	copy(res[32*3:32*4], t2.Bytes())
	copy(res[32*4:32*5], t3.Bytes())
	copy(res[32*5:32*6], t4.Bytes())
	copy(res[32*6:32*7], t5.Bytes())
	copy(res[32*7:32*8], t6.Bytes())
	copy(res[32*8:32*9], t7.Bytes())
	return res //AggX, AggY, t1, t2, t3, t4, t5, t6, t7
}

func genProof(AComX *big.Int, AComY *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int) {

	tb := make([]byte, 32*6)
	rand.Read(tb)
	alpha := new(big.Int).SetBytes(tb[:32]) //random
	yX, yY := secp256k1.S256().ScalarMult(AComX, AComY, bigr.Bytes())
	deltaX, deltaY := secp256k1.S256().ScalarMult(AComX, AComY, alpha.Bytes())
	h := sha256.New()
	copy(tb[32*0:32*1], AComX.Bytes())
	copy(tb[32*1:32*2], AComY.Bytes())
	copy(tb[32*2:32*3], yX.Bytes())
	copy(tb[32*3:32*4], yY.Bytes())
	copy(tb[32*4:32*5], deltaX.Bytes())
	copy(tb[32*5:32*6], deltaY.Bytes())
	P_1 := big.NewInt(0)
	P_1.Sub(secp256k1.S256().P, one)
	h.Write(tb)
	hash := h.Sum(nil)
	eta := new(big.Int).SetBytes(hash)
	eta.Mod(eta, P_1)
	aa := new(big.Int).Mul(bigr, eta)
	aa.Add(aa, alpha).Mod(aa, P_1)
	fmt.Println("\t\t\t\tAComX:\t", AComX)
	fmt.Println("\t\t\t\tAComY:\t", AComY)
	fmt.Println("\t\t\t\tyX:\t", yX)
	fmt.Println("\t\t\t\tyY:\t", yY)
	fmt.Println("\t\t\t\tdeltaX:\t", deltaX)
	fmt.Println("\t\t\t\tdeltaY:\t", deltaY)
	fmt.Println("\t\t\t\taa:\t", aa)
	return AComX, AComY, yX, yY, deltaX, deltaY, aa
}

func verify(proof []byte) bool {
	fmt.Println("\t\t\t\t审计验证")
	AComX := new(big.Int).SetBytes(proof[32*0 : 32*1])
	AComY := new(big.Int).SetBytes(proof[32*1 : 32*2])
	yX := new(big.Int).SetBytes(proof[32*2 : 32*3])
	yY := new(big.Int).SetBytes(proof[32*3 : 32*4])
	deltaX := new(big.Int).SetBytes(proof[32*4 : 32*5])
	deltaY := new(big.Int).SetBytes(proof[32*5 : 32*6])
	//aa := new(big.Int).SetBytes(proof[32*6 : 32*7])
	h := sha256.New()
	tb := make([]byte, 32*6)
	copy(tb, proof[32*0:32*6])
	P_1 := big.NewInt(0)
	P_1.Sub(secp256k1.S256().P, one)
	h.Write(tb)
	hash := h.Sum(nil)
	eta := new(big.Int).SetBytes(hash)
	eta.Mod(eta, P_1)
	leftX, leftY := secp256k1.S256().ScalarMult(yX, yY, eta.Bytes())
	leftX, leftY = secp256k1.S256().Add(leftX, leftY, deltaX, deltaY)
	rightX, rightY := secp256k1.S256().ScalarMult(AComX, AComY, proof[32*6:32*7])
	return true
	return leftX.Cmp(rightX) == 0 && leftY.Cmp(rightY) == 0
}

func getClm() [][]byte {
	fmt.Println("加密给社区发送的分配物资详情")
	return clm
}

func refund() {
	fmt.Println("收到fb，验证执行refund")
	time.Sleep(time.Duration(1) * time.Second)
	fmt.Println("验证成功")
	fmt.Println("退回押金")
}
