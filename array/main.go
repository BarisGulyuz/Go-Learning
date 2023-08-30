package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Learning Arrays in Go")

	/* DEFINITION AND SET VALES

	var strArr [3]string
	strArr[0] = "First"

	fmt.Println(strArr[0])

	strArr = [3]string{"Last", "Mid", "First"}

	fmt.Println(strArr[0])

	*/

	/* DEFAULT VALUES

	var intArr [3]int

	fmt.Println(intArr[0], intArr[1], intArr[2], ?intArr[3]?)

	*/

	//arraySize()

	//multiDimensionalArays()

	slice()

	maps()
}
func arraySize() {

	//len() --> getting array length
	//cap() --> getting array capacity

	arr := [6]int{0, 1, 2, 3} // 4 item added

	fmt.Println(len(arr), cap(arr)) // 0

	arr2 := [...]int{0, 1, 2, 3}

	fmt.Println(len(arr2), cap(arr2))

	//arr2[5] = 6 cannot do that

	var arr3 []int = make([]int, 10)

	arr3[0] = 1
	arr3[2] = 2

	fmt.Println(len(arr3), cap(arr3))

	arr3 = make([]int, cap(arr3)*2)

	fmt.Println(len(arr3), cap(arr3))
}
func multiDimensionalArays() {

	arr := [2][2]int{
		{0, 0},
		{1, 0}}

	fmt.Println(arr)
	fmt.Println(arr[0])
	fmt.Println(arr[0][1])
}
func slice() {
	// arr := [6]int{0, 1, 2, 3}

	// var slice []int = arr[1:3]
	// var slice2 []int = arr[:]

	// fmt.Println(slice, slice2)
	// fmt.Println(len(slice), cap(slice))

	var arr2 []int = []int{1, 2, 3, 4}
	fmt.Println(arr2, len(arr2), cap(arr2)) //len -> 4, cap -> 4

	arr2 = append(arr2, 5)                  // +1 item
	fmt.Println(arr2, len(arr2), cap(arr2)) // len -> 5, cap -> 8

	arr2 = append(arr2, 6)                  // +1 item
	fmt.Println(arr2, len(arr2), cap(arr2)) // len -> 6, cap -> 8

	fruits := make([]string, 2)
	fmt.Println(fruits)
	fruits[0] = "apple"
	fruits[1] = "banana"
	// fruits[2] = "asdsadas" error
	fmt.Println(fruits)

	fruits = append(fruits, "grape")
	fmt.Println(fruits)
}
func maps() {
	//key-value logic
	//c# -> dictionary ?

	//WAY 1
	//key -> int value -> string
	var students map[int]string = map[int]string{
		10: "Baris",
		20: "Zeynep",
	}

	fmt.Println(students)

	baris := students[10]

	fmt.Println(baris)

	//WAY 2

	var turkishCharacterToEnglish = make(map[string]string)

	turkishCharacterToEnglish["Ç"] = "C"
	turkishCharacterToEnglish["Ğ"] = "G"
	turkishCharacterToEnglish["İ"] = "I"
	turkishCharacterToEnglish["Ş"] = "S"
	turkishCharacterToEnglish["Ö"] = "O"
	//.......

	var nameWithTurkisChacter string = "ÖKKEŞ"

	var nameWithEnglishCharacter string = nameWithTurkisChacter
	for trChar, enChar := range turkishCharacterToEnglish {
		nameWithEnglishCharacter = strings.ReplaceAll(nameWithEnglishCharacter, trChar, enChar)
	}

	fmt.Println("nameWithEnglishCharacter --> ", nameWithEnglishCharacter)

	weekdays := map[string]string{
		"pazartesi": "Monday",
		"salı":      "Tuesday",
		"çarşamba":  "Wednesday",
		"perşembe":  "Thursday",
		"cuma":      "Friday",
		"cumartesi": "Saturday",
		"pazar":     "Sunday",
	}

	turkishDay := "çarşamba"
	englishDay, found := weekdays[turkishDay]

	if found {
		fmt.Printf("Türkçe gün: %s, İngilizce gün: %s\n", turkishDay, englishDay)
	} else {
		fmt.Printf("Belirtilen gün bulunamadı: %s\n", turkishDay)
	}

}
