package models

import (
	"fmt"
	"math/bits"
)

type ResultRow struct {
	DetectiveAbility string
	Count            uint64
	ClassSize        int
}

const N = 7
const K = 4
const InformationVector = 2       // 0010b
const CodedInformationVector = 22 // 0010.110 b
const GenPolynomial = 11          // 1011b

var Result = make([]ResultRow, N+1)
var ErrorClasses = GetErrorsByClasses(N)
var SyndromeTable = GetSyndromeTable(ErrorClasses[1], GenPolynomial)

func GetErrorClassesVar() *[][]uint64 {
	return &ErrorClasses
}

func GetSyndromeTableVar() *map[uint64]uint64 {
	return &SyndromeTable
}

func GetGenPolynomial() int {
	return GenPolynomial
}

func PowBinary(n uint64) uint64 {
	res := uint64(1)
	for i := uint64(1); i <= n; i++ {
		res <<= 1
	}
	return res
}

func GetBinaryLength(digit uint64) uint64 {
	bitsNum := uint64(0)
	for ; digit/2 != 0; digit /= 2 {
		bitsNum++
	}
	bitsNum++
	return bitsNum
}

func IntToBytes(digit uint64) []byte {
	var res []byte
	for i := PowBinary(GetBinaryLength(digit) - 1); i > 0; i /= 2 {
		res = append(res, byte(digit/i))
		digit %= i
	}
	return res
}

func Factorial(n uint64) (Result uint64) {
	if n > 0 {
		Result = n * Factorial(n-1)
		return Result
	}
	return 1
}

func ImposeError(input, err uint64) uint64 {
	inputBytes := IntToBytes(input)
	eBytes := IntToBytes(err)

	if len(inputBytes) > len(eBytes) {
		eBytes = append(make([]byte, len(inputBytes)-len(eBytes)), eBytes...)
	} else {
		inputBytes = append(make([]byte, len(eBytes)-len(inputBytes)), inputBytes...)
	}

	for index, errorBit := range eBytes {
		if errorBit == 1 {
			if inputBytes[index] == 1 {
				inputBytes[index] = 0
				continue
			}
			inputBytes[index] = 1
		}
	}
	input = 0 //renew the value
	for _, val := range inputBytes {
		input <<= 1
		input += uint64(val)
	}
	return input //altered value
}

func DivisionOperation(numerator, denominator uint64) (uint64, uint64) {
	if numerator < denominator {
		return 0, numerator //no division can be performed
	}

	var integer uint64
	inputBytes := IntToBytes(numerator)
	bLen := GetBinaryLength(denominator)
	var cur uint64

	inputBytesIndex := uint64(0)
	for ; inputBytesIndex < bLen; inputBytesIndex++ {
		cur <<= 1
		cur += uint64(inputBytes[inputBytesIndex])
	}

	for ; inputBytesIndex <= uint64(len(inputBytes)); inputBytesIndex++ { //division starting from 7-4
		firstBitInCur := cur / PowBinary(bLen-1)
		integer <<= 1
		integer += firstBitInCur

		if firstBitInCur == 1 {
			cur ^= denominator //shift 2 in
		}
		if inputBytesIndex == uint64(len(inputBytes)) {
			break
		}

		cur <<= 1
		cur += uint64(inputBytes[inputBytesIndex]) //continue filling up the cur
	}

	return integer, cur //upon finale Result integer and binary remaining are returned
}

func GetErrorsByClasses(n uint64) [][]uint64 { //n=number of classes
	ErrorClasses := make([][]uint64, n+1)
	for i := uint64(1); i <= n; i++ {
		size := Factorial(n) / Factorial(n-i) / Factorial(i) //C i n
		ErrorClasses[i] = make([]uint64, 0, size)
	}

	for i := uint64(1); i < PowBinary(n); i++ { //append all corresponding errors to their classes
		class := bits.OnesCount64(i)
		ErrorClasses[class] = append(ErrorClasses[class], i)
	}
	return ErrorClasses
}

func GetErrorsByClassesString(ErrorClasses [][]uint64) [][]string {
	errorsView := make([][]string, len(ErrorClasses))
	for class, errorClass := range ErrorClasses {
		errorsView[class] = make([]string, len(errorClass))
		for i, err := range errorClass {
			errorsView[class][i] = fmt.Sprintf("%b", err)
		}
	}
	return errorsView
}

func GetSyndromeTable(errorVectors []uint64, GenPolynomial uint64) map[uint64]uint64 {
	errorMap := make(map[uint64]uint64, len(errorVectors))
	for _, err := range errorVectors {
		_, syndrome := DivisionOperation(err, GenPolynomial)
		errorMap[syndrome] = err
	}
	return errorMap
}

func SyndromeTableToString(SyndromeTable map[uint64]uint64) map[string]string {
	SyndromeTableStr := make(map[string]string, len(SyndromeTable))
	for syndrome, err := range SyndromeTable {
		SyndromeTableStr[fmt.Sprintf("%b", syndrome)] = fmt.Sprintf("%b", err)
	}
	return SyndromeTableStr
}

func GetSyndromeArrayStr(n, GenPolynomial uint64) map[string]string {
	errorMap := make(map[string]string, PowBinary(n))
	for i := uint64(1); i < PowBinary(n); i++ {
		_, syndrome := DivisionOperation(i, GenPolynomial)
		errorMap[fmt.Sprintf("%b", i)] = fmt.Sprintf("%b", syndrome)
	}
	return errorMap
}

func Calculate() {
	for class, errorClass := range ErrorClasses {
		var detectedCounter uint64
		for _, errorVector := range errorClass {
			transferredVector := ImposeError(CodedInformationVector, errorVector)
			_, syndrome := DivisionOperation(transferredVector, GenPolynomial)
			if syndrome != 0 {
				detectedCounter++ //the error was detected successfully
			}
		}
		Result[class] = ResultRow{
			fmt.Sprintf("%.2f", float64(detectedCounter)*100/float64(len(errorClass))),
			detectedCounter,
			len(errorClass),
		}
	}
}
