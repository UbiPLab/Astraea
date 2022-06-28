#! /bin/bash

solc --abi DiSC.sol  | tee ./test/Test_sol_Test.abi
solc --bin DiSC.sol  | tee ./test/Test_sol_Test.bin
python change.py
#go run main.go
