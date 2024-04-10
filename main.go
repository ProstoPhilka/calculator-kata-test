package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	a, operator, b, errInput := Input()
	if errInput != nil {
		panic(errInput)
	}

	isRomeA, errDetectA := Detect(a)
	isRomeB, errDetectB := Detect(b)

	if isRomeA != isRomeB || errDetectA != nil || errDetectB != nil {
		panic("Error: Incorrect input format")
	} else {
		if isRomeA {
			//римские числа считаем

			a, _ := ToNumerals(a)
			b, _ := ToNumerals(b)

			intermediateResult, errCalculator := Calculator(a, b, operator)
			if errCalculator != nil {
				panic(errCalculator)
			}

			result, errToRoman := ToRoman(intermediateResult)

			if errToRoman != nil {
				panic("Error: Output number must be between 1 and 999 in roman system")
			}

			fmt.Println(result)
		} else {
			//считаем числа в десятичной системе

			operandA, _ := strconv.Atoi(a)
			operandB, _ := strconv.Atoi(b)

			min := 1
			max := 10

			if !(InRange(min, max, operandA) && InRange(min, max, operandB)) {
				err := "Error: Operand a or operand b is not in the range " + strconv.Itoa(min) + " to " + strconv.Itoa(max)
				panic(err)
			}

			result, errCalculator := Calculator(operandA, operandB, operator)
			if errCalculator != nil {
				panic(errCalculator)
			}

			fmt.Println(result)
		}
	}
}

func InRange(a int, b int, data int) bool {

	if data >= a && data <= b {
		return true
	}

	return false
}

func ToRoman(data int) (string, error) {

	//data приндалежит диапазону [1; 999]

	err := errors.New("Error ToRoman: Number must be between 1 and 999")
	if data < 1 || data > 999 {
		return "", err
	}

	ones := [10]string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	tens := [10]string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	hunds := [10]string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}

	h := data / 100 % 10
	t := data / 10 % 10
	o := data % 10

	return hunds[h] + tens[t] + ones[o], nil
}

func Detect(data string) (bool, error) {

	_, errToNumerals := ToNumerals(data)

	if errToNumerals != nil {
		_, errAtoi := strconv.Atoi(data)
		if errAtoi != nil {
			err := errors.New("Error Detect: Failed to determine operand")
			return false, err
		} else {
			//Операнд десятичным числом
			return false, nil
		}
	} else {
		//Операнд является римским числом от I до X
		return true, nil
	}
}

func Input() (string, string, string, error) {

	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	err := errors.New("Error: Incorrect input format")

	subStrInput := strings.Fields(input)
	if len(subStrInput) != 3 {
		return "", "", "", err
	}

	return subStrInput[0], subStrInput[1], subStrInput[2], nil
}

func ToNumerals(data string) (int, error) {

	// data принадлежит диапазону от I до X

	converter := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}

	err := errors.New("Error ToNumerals: Number must be between I and X")
	if _, ok := converter[data]; !ok {
		return 0, err
	}

	result, _ := converter[data]
	return result, nil
}

func Calculator(a int, b int, operator string) (int, error) {

	err := errors.New("Error: Operator not found")
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return Division(a, b)
	default:
		return 0, err
	}
}

func Division(a int, b int) (int, error) {

	err := errors.New("Error: Division by zero")
	if b == 0 {
		return 0, err
	}

	return a / b, nil
}
