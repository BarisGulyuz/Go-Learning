package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {

	SayMyName("Baris")

	sum, err := sumPositiveInt(-1, 4)

	if err == nil {
		fmt.Println("Sum1 --->", sum)
	} else {
		fmt.Println("Error1 --->", err)
	}

	sum2, err2 := sumPositiveInt2(-3, 4)

	if err2 == nil {
		fmt.Println("Sum2 --->", sum2)
	} else {
		fmt.Println("Error2 --->", err)
	}

	sum3 := sumMultiple(1, 2, 3, 4, 5)
	fmt.Println("Sum3--->", sum3)

	defer closeRabbitMQConntection()
	openRabbitMQConntection()

	sum4 := sumMultiple(1, 2, 3, 4, 5)
	fmt.Println("Sum4--->", sum4)

	sendMessageToQeue(strconv.Itoa(sum4))

	fmt.Println("F ->>", fibonacci(10))
}

// function with parameter
func SayMyName(name string) {
	fmt.Println(name)
}

// function that return value
// retuns integer
// taking 2 integer parameter
func sumInt(num1 int, num2 int) int {
	return num1 + num2
}

// function that return multiple values
func sumPositiveInt(num1 int, num2 int) (int, error) {

	if num1 <= 0 || num2 <= 0 {
		return 0, errors.New("This method only sum positive numbers")
	}

	return num1 + num2, nil
}

// function that return multiple values
func sumPositiveInt2(num1 int, num2 int) (sum int, err error) {

	if num1 <= 0 || num2 <= 0 {
		err = errors.New("This method only sum positive numbers")
	} else {
		sum = num1 + num2
	}

	return
}

// function that dynamic parameters

func sumMultiple(numbers ...int) (sum int) {

	// for i := 0; i < len(numbers); i++ {
	// 	sum += numbers[i]
	// }

	for _, number := range numbers {
		sum += number
	}

	return
}

func openRabbitMQConntection() {
	fmt.Println("RabbitMQ Connection Opened")
}

func sendMessageToQeue(message string) {
	fmt.Println("Message Sent -->", message)
}

func closeRabbitMQConntection() {
	fmt.Println("RabbitMQ Connection Closed")
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	sum1 := fibonacci(n - 1)
	sum2 := fibonacci(n - 2)

	return sum1 + sum2
}

type predicate func(s string) bool

func Where(stringArr []string, predicate predicate) (returnValues []string) {

	for i := 0; i < len(stringArr); i++ {
		var currentValue string = stringArr[i]
		if predicate(currentValue) {
			returnValues = append(returnValues, currentValue)
		}
	}

	return
}
