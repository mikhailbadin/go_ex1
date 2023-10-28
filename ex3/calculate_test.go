package main

import (
	"fmt"
	"testing"
)

/*
Реализовать функцию calculate которая принимает на вход арифметическое выражение и выводит его результат

Требования и ограничения:
- Работа идет только с целыми числами (в тч отрицательными)
- Поддеживаются следующие арифметические операции "+", "-", "*"
- Важен приоритет операций: приоритет умножения больше чем сложение или вычитание
- БОНУС: Сделать поддежку скобок "(" и ")"

Подсказки:
- смотри на Lexer и Parser
- смотри на BNF диаграмму
*/

func calculate(str string) int {
	return 1
}

// Don't edit test below

func TestCalculate(t *testing.T) {
	data := map[string]int{
		"1 + 2":           3,
		"2 - 1":           1,
		"-2 + 4":          2,
		"3 + 2 - 1":       4,
		"2 * 3":           6,
		"2 * 3 + 1":       7,
		"6 - 2 * 1":       4,
		"5 * 2 - 2 * 3":   4,
		"-1 + -1":         -2,
		"   1   +    1":   2,
		"5 * (2 - 1) * 2": 10,
		" (1 + 1) * -1 ":  -2,
		" -(1 + 1) ":      -2,
	}

	for str, expected := range data {
		testname := fmt.Sprintf(str)
		t.Run(testname, func(t *testing.T) {
			result := calculate(str)
			if expected != result {
				t.Errorf("got %d, expected: %d", result, expected)
			}
		})
	}
}
