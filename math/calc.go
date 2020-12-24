package math

import "github.com/google/uuid"

//Add exported
func Add(num1 int, num2 int) int {
	return num1 + num2
}

//SubtractNums exported
func SubtractNums(num1, num2 int) int {
	return num1 - num2
}

//GetUUID exported
func GetUUID() uuid.UUID {
	return uuid.New()
}
