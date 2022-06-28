package main

import (
"bytes"
"encoding/hex"
	"fmt"
	"math/big"

"github.com/ethereum/go-ethereum/accounts/abi"
)

func main(){
	const method = "aTestMethod"
	const output = "ffffffffffffffffffffffffffffffffffffffffffffffec6dc439894e200000"
	const abiJSON = `[{
  "constant":true,
  "inputs":[
  ],
  "name":"aTestMethod",
  "outputs":[
    {
      "name":"",
      "type":"int256"
    }
  ],
  "payable":false,
  "stateMutability":"view",
  "type":"function"
}]`

	var ret = big.NewInt(0)
	outputBytes, err := hex.DecodeString(output)
	if err != nil {

	}
	pricingABI, err := abi.JSON(bytes.NewReader([]byte(abiJSON)))
	if err != nil {

	}

	if err = pricingABI.Unpack(&ret, method, outputBytes); err != nil {

	}
	if ret.Sign() != -1 {
		fmt.Println("result should be negative, got: %s", ret.String())
	}
}
