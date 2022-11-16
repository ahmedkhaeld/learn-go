package main

import "fmt"

//1.Is verbose mode activated?          00000001=1
//
//2.Is configuration loaded from disk?  00000010=2
//
//3.Is database connexion required?     00000100=3
//
//4.Is logger activated?               00001000=4
//
//5.Is debug mode activated?           00010000=5
//
//6.Is support for float activated?    00100000=6
//
//7.Is recovery mode activated?        01000000=7
//
//8.Reboot on failure?                10000000=8

// bytes read from right to left
// represent all those environment variables in one byte [8 bits] each bit is flag of each config
// instead of using 8 booleans which will take 8 bytes
//use shift left each const by iota

type MyConf uint8

// shift left:  num<<shifts
const (
	VERBOSE          MyConf = 1 << iota //no shift
	ConfigFromDisk   MyConf = 1 << 1    // shift the bit[1] one position to the left
	DatabaseRequired MyConf = 1 << 2    // shift the bit[1] two positions to the left
	LoggerActivated  MyConf = 1 << 3
	DEBUG            MyConf = 1 << 4
	FloatSupport     MyConf = 1 << 5
	RecoveryMode     MyConf = 1 << 6
	RebootOnFailure  MyConf = 1 << 7
)

func main() {
	MyComplexFunction(
		VERBOSE|ConfigFromDisk|DatabaseRequired|LoggerActivated|DEBUG|FloatSupport|RecoveryMode|RebootOnFailure,
		"test")
}

func MyComplexFunction(conf MyConf, databaseDsn string) {
	fmt.Printf("conf : %08b\n", conf)

	test := conf & RebootOnFailure
	fmt.Printf("test : %08b\n", test)
}
