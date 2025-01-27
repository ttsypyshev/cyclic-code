package backend

import "fmt"

// Возвращает классы ошибок для внешнего использования
func GetErrorClassesVar() *[][]uint64 {
	return &ErrorClasses
}

// Возвращает таблицу синдромов для внешнего использования
func GetSyndromeTableVar() *map[uint64]uint64 {
	return &SyndromeTable
}

// Возвращает порождающий полином
func GetGenPolynomial() int {
	return GenPolynomial
}

// Функция для возведения в степень в двоичной системе
func PowBinary(n uint64) uint64 {
	res := uint64(1)
	for i := uint64(1); i <= n; i++ {
		res <<= 1
	}
	return res
}

// Функция для получения количества битов в числе
func GetBinaryLength(digit uint64) uint64 {
	bitsNum := uint64(0)
	for ; digit/2 != 0; digit /= 2 {
		bitsNum++
	}
	bitsNum++
	return bitsNum
}

// Функция для преобразования числа в массив байтов
func IntToBytes(digit uint64) []byte {
	var res []byte
	for i := PowBinary(GetBinaryLength(digit) - 1); i > 0; i /= 2 {
		res = append(res, byte(digit/i))
		digit %= i
	}
	return res
}

// Факториал числа с рекурсией
func Factorial(n uint64) (Result uint64) {
	if n > 0 {
		Result = n * Factorial(n-1)
		return Result
	}
	return 1
}

// Функция для представления ошибок в виде строк
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

// Функция для представления таблицы синдромов в виде строк
func SyndromeTableToString(SyndromeTable map[uint64]uint64) map[string]string {
	SyndromeTableStr := make(map[string]string, len(SyndromeTable))
	for syndrome, err := range SyndromeTable {
		SyndromeTableStr[fmt.Sprintf("%b", syndrome)] = fmt.Sprintf("%b", err)
	}
	return SyndromeTableStr
}

// Функция для получения массива синдромов в виде строк
func GetSyndromeArrayStr(n, GenPolynomial uint64) map[string]string {
	errorMap := make(map[string]string, PowBinary(n))
	for i := uint64(1); i < PowBinary(n); i++ {
		_, syndrome := DivisionOperation(i, GenPolynomial)
		errorMap[fmt.Sprintf("%b", i)] = fmt.Sprintf("%b", syndrome)
	}
	return errorMap
}

// Функция для деления двух чисел в двоичной системе
func DivisionOperation(numerator, denominator uint64) (uint64, uint64) {
	if numerator < denominator {
		// Если делитель больше числителя, деление не может быть выполнено
		return 0, numerator
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

	// Основной цикл для деления
	for ; inputBytesIndex <= uint64(len(inputBytes)); inputBytesIndex++ {
		firstBitInCur := cur / PowBinary(bLen-1)
		integer <<= 1
		integer += firstBitInCur

		// Если старший бит в cur равен 1, вычитаем делитель
		if firstBitInCur == 1 {
			cur ^= denominator // сдвиг влево
		}
		if inputBytesIndex == uint64(len(inputBytes)) {
			break
		}

		cur <<= 1
		cur += uint64(inputBytes[inputBytesIndex]) // продолжаем заполнять cur
	}

	return integer, cur // Возвращаем результат деления и остаток
}
