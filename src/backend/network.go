package backend

import (
	"fmt"
	"log"
	"math/bits"
)

type ResultRow struct {
	DetectiveAbility string // Вероятность обнаружения ошибок в процентах
	Count            uint64 // Количество обнаруженных ошибок
	ClassSize        int    // Общее количество ошибок в классе
}

const N = 7                       // Длина кодового слова
const K = 4                       // Размер информационного вектора
const InformationVector = 8       // Исходный информационный вектор (1000b)
const CodedInformationVector = 69 // Закодированный информационный вектор (1000.101b)
const GenPolynomial = 11          // Порождающий полином (1011b)

// Результат подсчета вероятностей по классам кратности
var Result = make([]ResultRow, N+1)

// Классы ошибок, сгруппированные по кратности (числу изменённых бит)
var ErrorClasses = GetErrorsByClasses(N)

// Таблица синдромов, связывающая синдром с вектором ошибки
var SyndromeTable = GetSyndromeTable(ErrorClasses[1], GenPolynomial)

// Функция для получения классов ошибок, сгруппированных по кратности
func GetErrorsByClasses(n uint64) [][]uint64 {
	ErrorClasses := make([][]uint64, n+1) // Инициализация списка классов
	sum_size := 0                         // Общий размер всех классов

	// Вычисляем количество ошибок в каждом классе кратности
	// Формула: C(n, i) = n! / (i! * (n - i)!)
	for i := uint64(1); i <= n; i++ {
		size := Factorial(n) / Factorial(n-i) / Factorial(i) // Биномиальный коэффициент
		ErrorClasses[i] = make([]uint64, 0, size)            // Выделяем память под ошибки класса
		log.Println("Кратность:", i, "количество ошибок:", size)
		sum_size += int(size)
	}
	log.Println("Суммарное количество ошибок:", sum_size)

	// Генерируем векторы ошибок и добавляем их в соответствующие классы
	for i := uint64(1); i < PowBinary(n); i++ {
		class := bits.OnesCount64(i) // Кратность определяется числом единиц в векторе
		ErrorClasses[class] = append(ErrorClasses[class], i)
	}
	return ErrorClasses
}

// Функция для построения таблицы синдромов
func GetSyndromeTable(errorVectors []uint64, GenPolynomial uint64) map[uint64]uint64 {
	errorMap := make(map[uint64]uint64, len(errorVectors)) // Карта синдромов

	// Для каждого вектора ошибки вычисляем синдром
	// Формула деления: Остаток от деления ошибки на порождающий полином
	for _, err := range errorVectors {
		_, syndrome := DivisionOperation(err, GenPolynomial) // Остаток от деления
		errorMap[syndrome] = err                             // Сохраняем соответствие синдрома и ошибки
	}
	return errorMap
}

// Наложение ошибки на кодовое слово (имитация передачи через канал с ошибками)
func ImposeError(input, err uint64) uint64 {
	inputBytes := IntToBytes(input) // Преобразуем входное число в массив битов
	eBytes := IntToBytes(err)       // Преобразуем вектор ошибки в массив битов

	// Выравниваем длины массивов ошибок и входных данных
	if len(inputBytes) > len(eBytes) {
		eBytes = append(make([]byte, len(inputBytes)-len(eBytes)), eBytes...)
	} else {
		inputBytes = append(make([]byte, len(eBytes)-len(inputBytes)), inputBytes...)
	}

	// Побитовая операция наложения ошибок
	// Формула: input XOR error
	for index, errorBit := range eBytes {
		if errorBit == 1 {
			if inputBytes[index] == 1 {
				inputBytes[index] = 0 // Инверсия бита
				continue
			}
			inputBytes[index] = 1
		}
	}

	// Преобразуем массив битов обратно в число
	input = 0
	for _, val := range inputBytes {
		input <<= 1
		input += uint64(val)
	}
	return input
}

// Расчёт вероятности обнаружения ошибок для каждого класса кратности
func Calculate() {
	for class, errorClass := range ErrorClasses {
		var detectedCounter uint64 // Счётчик обнаруженных ошибок

		// Проверяем все ошибки в классе
		for _, errorVector := range errorClass {
			// Накладываем ошибку на кодированное слово
			transferredVector := ImposeError(CodedInformationVector, errorVector)
			// Вычисляем синдром
			// Формула деления: Остаток от деления переданного вектора на порождающий полином
			_, syndrome := DivisionOperation(transferredVector, GenPolynomial)
			if syndrome != 0 {
				detectedCounter++ // Ошибка успешно обнаружена
			}
		}

		// Сохраняем результаты для текущего класса
		Result[class] = ResultRow{
			DetectiveAbility: fmt.Sprintf("%.2f", float64(detectedCounter)*100/float64(len(errorClass))),
			Count:            detectedCounter,
			ClassSize:        len(errorClass),
		}
	}
}
