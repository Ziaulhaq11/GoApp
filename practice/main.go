package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func challenge() {
	result := calculate("10", "5.5")
	fmt.Println(result)
}

func calculate(value1 string, value2 string) float64 {
	f1, err := strconv.ParseFloat(strings.TrimSpace(value1), 64)
	if err != nil {
		fmt.Println("Error Occurred")
	}
	f2, err2 := strconv.ParseFloat(strings.TrimSpace(value2), 64)
	if err2 != nil {
		fmt.Println("Error Occurred")
	}
	result := f1 + f2
	return result
}

func dateAndTime() {
	n := time.Now()
	fmt.Println("Time now is ", n) //2024-09-05 12:58:39.1531853 +0530 IST m=+0.000504801

	t := time.Date(2024, 8, 20, 23, 0, 0, 0, time.UTC)
	fmt.Println("Time given is ", t)  //2024-08-20 23:00:00 +0000 UTC
	fmt.Println(t.Format(time.ANSIC)) //Tue Aug 20 23:00:00 2024

	parsedTime, _ := time.Parse(time.ANSIC, "Tue Aug 20 23:00:00 2024")
	fmt.Printf("Type of parsedTime is %T\n", parsedTime) //Type of parsedTime is time.Time

}

func main() {
	i1, i2, i3 := 12, 45, 68

	intSum := i1 + i2 + i3
	fmt.Println("Integer sum is ", intSum)

	f1, f2, f3 := 23.5, 65.2, 76.4
	floatSum := f1 + f2 + f3
	fmt.Println("Float Sum : ", floatSum) //For Float Sum its precision is not correct as we can see 165.10000000000002. As floating values are stored in binary format like in Java etc.
	// floatSum = math.Round(floatSum)
	// fmt.Println("The sum is now fixed : ", floatSum) //But it rounds to nearest Integer
	floatSum = math.Round(floatSum*100) / 100
	fmt.Println("The sum is now fixed : ", floatSum)

	circleRadius := 15.5
	circumference := circleRadius * 2 * math.Pi
	fmt.Printf("Circumference : %T\n", circumference)   //float64
	fmt.Printf("Circumference : %.2f\n", circumference) //97.39
}

// Get Input From Console
func inputFromConsole() {
	reader := bufio.NewReader(os.Stdin) //where to read, from console
	fmt.Print("Enter text: ")
	input, _ := reader.ReadString('\n') //When enter presses then it stops I think and single quotes for Char Type
	fmt.Println("You Entered : ", input)

	fmt.Print("Enter a number: ")
	numInput, _ := reader.ReadString('\n')
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64) //64 or 32 bit
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value of number : ", aFloat)
	}
}

// Variables Basics
const aConst = 64 //These are declared outside of Functions mostly, and if initial letter is capital it would be PUBLIC

func variables() {
	//Explicit types
	var aString string = "Hello from Go"
	fmt.Println(aString)
	fmt.Printf("Variable type is %T\n", aString)

	var anInteger int = 42 //Default value 0
	fmt.Println(anInteger)
	fmt.Printf("%T\n", anInteger)

	//Implicit Typing
	var anotherStr = "This is another string"
	fmt.Println(anotherStr)
	fmt.Printf("Variable Type %T\n", anotherStr)

	//Another way of Declaration
	myString := "this is also a sting" //These can only be declared inside a functions else use VAR keyword
	fmt.Println(myString)
}
