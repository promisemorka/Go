package main

import (
	"fmt"
	"reflect"
	"strconv"
)



func main()  {
// 	// var greeting string = "Hello World"
// 	// fmt.Println(greeting)

// 	var name string = "Jack"
// 	var i int = 80
// 	fmt.Printf("Hey, %v! You have scored %d/100 in Physics.", name, i)

	// const PI float64 = 3.14 // global constant

    // var radius float64 = 5.154
    // var area float64

    // area = PI * radius * radius
    // fmt.Printf("Radius: %.2f \nPI:%.1f \n", radius, PI)
    // fmt.Printf("Area of Circle is : %f", area)
    // fmt.Println("Thank You")

	// var name string
	// var is_muggle bool

	// fmt.Print("Enter your name & are you a muggle: ")
	// fmt.Scanf("%s %t", &name, &is_muggle)
	// fmt.Println(name, is_muggle)

	fmt.Println("--------------------")

	// ******** MULTIPLE INPUTS ***************

	// var a string
	// var b int
	// fmt.Print("Enter a string and a number: ")

	// count, err := fmt.Scanf("%s %d", &a, &b)

	// fmt.Println("count : ", count)
	// fmt.Println("error: ", err)
	// fmt.Println("a: ", a)
	// fmt.Println("b: ", b)

	fmt.Println("--------------------")

	// *********** FIND THE TYPE OF VARIABLE **********
	var grades int = 32
	var message string = "Hello World"
	var isCheck bool = true
	var amount float32 = 5466.54

	fmt.Printf("variable grades = %v is of type %T \n", grades, grades)
	fmt.Printf("variable message = %v is of type %T \n", message, message)
	fmt.Printf("variable isCheck = %v is of type %T \n", isCheck, isCheck)
	fmt.Printf("variable amount = %v is of type %T \n", amount, amount)

	fmt.Printf("Type: %v \n", reflect.TypeOf(1000))
	fmt.Printf("Type: %v \n", reflect.TypeOf("Promise"))
	fmt.Printf("Type: %v \n", reflect.TypeOf(46.0))
	fmt.Printf("Type: %v \n", reflect.TypeOf(true))

	fmt.Println("--------------------")

	var level int = 5
	var sentence string = "Never give up"

	fmt.Printf("variable level=%v is of type %v \n", level, reflect.TypeOf(level))
	fmt.Printf("variable sentence='%v' is of type %v \n", sentence, reflect.TypeOf(sentence))

	fmt.Println("--------------------")

	// *********** TYPE CASTING ******************
	var age int = 24
	var ageFloat float64 = float64(age)
	fmt.Printf("%.2f\n", ageFloat)

	var degree float64 = 72.55
	var degreeInteger int = int(degree)
	fmt.Printf("%v\n", degreeInteger)

	fmt.Println("--------------------")

	var temp int = 20
	var tempStr = strconv.Itoa(temp) // Converts integer to string -> returns one value - the string formed with the given integer
	fmt.Printf("%q\n", tempStr)

	fmt.Println("--------------------")

	var weight string = "115"
	i, err := strconv.Atoi(weight) // Converts string to integer -> returns two values - the integer, and error (if any)

	fmt.Printf("%v, %T \n", i, i)
	fmt.Printf("%v, %T \n", err, err)

	fmt.Println("--------------------")

	// ************ CONSTANTS ************
	/*
	 Untyped constant
	  . Constants are untyped unless they are explicitly given a type at declaration
	  . Allow for flexibility
	
	 Typed constant
      . Constants are typed when you explicitly specify the type in the declaration
	  . Flexibility that comes with untyped constant is lost
	*/

	const club = "Chelsea FC"
	const is_playing_well = false
	const league_position = 7

	fmt.Printf("%v: %T \n", club, club)
	fmt.Printf("%v: %T \n", is_playing_well, is_playing_well)
	fmt.Printf("%v: %T \n", league_position, league_position)

	fmt.Println("--------------------")


}