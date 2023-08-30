package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	//there is no try catch feature like c# exception handling

	// nil -> c# null

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter a number :")
	scanner.Scan()

	var number, err = strconv.ParseInt(strings.Trim(scanner.Text(), " "), 10, 64)

	if err == nil {
		fmt.Println(number)
	} else {
		log.Fatal(err)
	}
}
