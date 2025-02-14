# Исследование вероятности обнаружения ошибок в линейных кодах

Этот проект реализует вычисление вероятности обнаружения ошибок в коде, закодированном с использованием линейного кода, вероятно, из класса циклических кодов. Основная цель — исследовать классы ошибок, их синдромы и вероятность обнаружения ошибок для каждого класса.

---

## Алгоритм работы

### Этапы выполнения:
1. **Определение параметров кода:**
   - Длина кодового слова (\( N \)) = 7.
   - Размер информационного вектора (\( K \)) = 4.
   - Порождающий полином (\( \text{GenPolynomial} \)) = 11.
   - Исходный информационный вектор (\( \text{InformationVector} \)) = 8.
   - Закодированный информационный вектор (\( \text{CodedInformationVector} \)) = 69.

2. **Формирование классов ошибок:**
   - Ошибки группируются по числу изменённых битов (кратности). 
   - Используется комбинаторная формула для подсчёта количества ошибок.

3. **Генерация таблицы синдромов:**
   - Для каждого вектора ошибки вычисляется синдром с использованием деления на порождающий полином.
   - Результаты записываются как \( \text{Синдром} \to \text{Вектор ошибки} \).

4. **Наложение ошибок:**
   - Моделируется процесс передачи данных через канал с ошибками с помощью побитового наложения ошибок на закодированный вектор.

5. **Вычисление вероятности обнаружения ошибок:**
   - Проверяется, можно ли обнаружить ошибку по синдрому. Если синдром ненулевой, ошибка считается обнаруженной.
   - Для каждого класса ошибок вычисляется вероятность обнаружения.

---

## Основные функции

- **`GetErrorsByClasses(n uint64) [][]uint64`**  
  Формирует классы ошибок по кратности, используя битовые операции.

- **`GetSyndromeTable(errorVectors []uint64, GenPolynomial uint64) map[uint64]uint64`**  
  Вычисляет синдромы для каждого вектора ошибок.

- **`ImposeError(input, err uint64) uint64`**  
  Накладывает вектор ошибок на данные с помощью операции XOR.

- **`Calculate()`**  
  Основная функция, которая рассчитывает вероятность обнаружения ошибок для всех классов.

---

## Пример использования
1. Генерация классов ошибок.
2. Построение синдромной таблицы.
3. Расчёт вероятности обнаружения ошибок.

Результаты сохраняются в массиве `Result`, где указаны:
- Вероятность обнаружения ошибок.
- Общее число ошибок в классе.
- Количество обнаруженных ошибок.

---

## Запуск
1. Склонируйте репозиторий.
2. Запустите файл `main.go` командой:
   ```bash
   go run main.go
