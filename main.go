package main

import (
	"bufio"
	"fmt"
	helper "hello/fistFolder"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strconv"
)

const (
	NAME     string = "BARİS"
	AGE      int    = 25
	FAV_LANG string = "C#"
)

func main() {

	rootOfEqunation()

	name := helper.GetName()
	fmt.Println("Name -->", name)

	intArr := []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println((intArr))

	slice := intArr[0:3]

	fmt.Println((slice))

	for i := 0; i < len(intArr); i++ {
		fmt.Println(intArr[i])
	}

	var newName *string = &name
	*newName = "Zeynep"

	fmt.Println(name)
	fmt.Println(newName)
	fmt.Println(*newName)

	var me Person = Person{name: "Baris", surname: "Gülyüz", age: 25}

	fmt.Println(me)

	var areNamesEqual bool = (name == me.name)

	if areNamesEqual {
		fmt.Println("Names Are Same")
	} else {
		fmt.Println("Names Are Not Same")
	}

	fmt.Println(math.MaxInt)

	var maxInt int64 = 922337203685775807

	fmt.Println(maxInt)

	// fmt.Println(areNamesEqual)

	callHttp()
}

func callHttp() {
	ress, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	fmt.Println(ress.Body)
	fmt.Println(err)

	resBody, err := ioutil.ReadAll(ress.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("client: response body: %s\n", resBody)
}

func rootOfEqunation() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("a : ")

	scanner.Scan()
	a, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Print("b : ")

	scanner.Scan()
	b, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Print("c : ")

	scanner.Scan()
	c, _ := strconv.ParseFloat(scanner.Text(), 64)

	delta := math.Pow(b, 2) - 4*a*c

	fmt.Println("Delta -->", delta)

	root1 := (-b - math.Pow(delta, 0.5)) / 2 * a
	root2 := (-b + math.Pow(delta, 0.5)) / 2 * a

	fmt.Println(root1, root2)
}

type Person struct {
	name    string
	surname string
	age     int
}

/*


 */
