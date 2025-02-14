# Теория кодирования ошибок

## 1. Ошибка
### Определение
Ошибка в передаче данных — это изменение одного или нескольких битов в кодовом слове во время передачи через канал связи. 

Примеры:
- Одиночная ошибка: изменён только один бит.
- Множественная ошибка: изменены несколько битов.
- Блочные ошибки: изменены последовательные группы битов.

---

## 2. Синдром ошибки
### Определение
Синдром — это остаток от деления принятого слова на порождающий полином. 

Формула:
\[
S(x) = R(x) \mod G(x)
\]

- Если синдром равен нулю, ошибок нет или ошибка не может быть обнаружена.
- Если синдром ненулевой, ошибка обнаружена.

Пример:
- Получено \( R(x) = 1001101 \), порождающий полином \( G(x) = 1011 \).
- Остаток деления (синдром) \( S(x) \) определяет наличие ошибок.

---

## 3. Кратность ошибки
Количество битов, изменённых в кодовом слове.  
Пример:
- Для ошибки \( 101 \) кратность равна 2 (изменены два бита).

Кратность помогает классифицировать ошибки и оценивать устойчивость кодов.

---

## 4. Вектор ошибок
Последовательность битов, представляющая ошибки.  
Примеры:
- \( 0001 \): ошибка в последнем бите.
- \( 0101 \): ошибки во втором и четвёртом битах.

---

## 5. Классы ошибок
Ошибки группируются по кратности (числу изменённых битов).  

Пример:
- Класс кратности 1: \( 0001, 0010, 0100, 1000 \).
- Класс кратности 2: \( 0011, 0101, 0110, 1001 \).

Использование:
1. Оценка вероятности обнаружения ошибок.
2. Построение синдромных таблиц.
3. Анализ избыточности кода.

---

## 6. Линейный код
### Определение
Линейный код обладает свойством: сумма любых двух кодовых слов — это кодовое слово.  

Формула кодирования:
\[
C = M \cdot G
\]
Где:
- \( C \): Кодовое слово.
- \( M \): Информационный вектор.
- \( G \): Генераторная матрица.

Преимущества:
- Простота реализации.
- Возможность систематического кодирования.

---

## 7. Циклический код
### Определение
Циклический код — это линейный код, в котором циклический сдвиг кодового слова также является кодовым словом.

Пример:
- Кодовое слово \( C(x) = 1101 \). После сдвига \( 0111 \) также принадлежит коду.

Формулы:
- Кодирование: \( C(x) = M(x) \cdot G(x) \).
- Декодирование: остаток деления принятого слова на \( G(x) \).

---

## 8. Примеры кодов
### Код Хэмминга (\( (7,4) \))
- Обнаруживает и исправляет одиночные ошибки.
- Минимальное расстояние: \( d_{\text{min}} = 3 \).

### CRC (Cyclic Redundancy Check)
- Проверяет целостность данных.
- Пример порождающего полинома: \( G(x) = x^3 + x + 1 \).
