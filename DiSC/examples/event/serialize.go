package main

import (
	"errors"
	"math/big"
)

// all big.Int below are 32 bytes

// reqi = (Ri, ci, pkcoi, dni)
type request struct {
	ReqRi   *big.Int // request order
	ReqCi   *big.Int // collector oder
	pubKeyX *big.Int // donee's community pubKey
	pubKeyY *big.Int
	ReqDni  *big.Int // request donation number
}

// resj = (tp1, lc, idj1, dnj, clj, pkcoi, ¯rj, reqi, dt)
type response struct {
	req *request // donee's request

	tp1  *big.Int // transpoder order
	lc   *big.Int // logic company order
	idj1 *big.Int //
	dnj  *big.Int // donator order
	clj  *big.Int // check list, limit 32 bytes
	//pubKeyX *big.Int // collector's pubKey
	//pubKeyY *big.Int // in my option, it is no use
	cm  *big.Int
	tok *big.Int
	rj  *big.Int // random used for PoA
	dt  *big.Int // datetime,limit 32 bytes
}

// sigma is signature

func Deserialize(b []byte) (*response, error) {
	if len(b) != 480 {
		return nil, errors.New("byte length not match")
	}

	req := &request{
		ReqRi:   new(big.Int).SetBytes(b[8*32 : 9*32]),
		ReqCi:   new(big.Int).SetBytes(b[9*32 : 10*32]),
		pubKeyX: new(big.Int).SetBytes(b[11*32 : 12*32]),
		pubKeyY: new(big.Int).SetBytes(b[12*32 : 13*32]),
		ReqDni:  new(big.Int).SetBytes(b[13*32 : 14*32]),
	}
	res := &response{
		tp1:  new(big.Int).SetBytes(b[0*32 : 1*32]),
		lc:   new(big.Int).SetBytes(b[1*32 : 2*32]),
		idj1: new(big.Int).SetBytes(b[2*32 : 3*32]),
		dnj:  new(big.Int).SetBytes(b[3*32 : 4*32]),
		clj:  new(big.Int).SetBytes(b[4*32 : 5*32]),
		cm:   new(big.Int).SetBytes(b[5*32 : 6*32]),
		tok:  new(big.Int).SetBytes(b[6*32 : 7*32]),
		rj:   new(big.Int).SetBytes(b[7*32 : 8*32]),
		req:  req,
		dt:   new(big.Int).SetBytes(b[14*32 : 15*32]),
	}
	return res, nil
}
