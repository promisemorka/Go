package main

import "fmt"

func main() {

	var city string = "Dallas"
	var city_2 string = "Frisco"

	var a, b int = 5, 10

	// ******* COMPARISON OPERATORS *******
	fmt.Println(city == city_2)
	fmt.Println(city != city_2)
	fmt.Println(a < b)
	fmt.Println(a <= b)
	fmt.Println(a > b)
	fmt.Println(a >= b)

	fmt.Println("--------------------")

	// ***** ARITHMETIC OPERATORS *****
	var color,phone string = "blue", "iPhone"
	fmt.Println(color + phone)

	var temp,temp_two float64 = 79.33, 75.00
	fmt.Printf("%.2f \n", temp - temp_two)

	fmt.Println("--------------------")

	var length,breadth int = 12, 4
	fmt.Println(length * breadth)
	fmt.Println(length / breadth)
	fmt.Println(length % breadth)

	var i int = 1
	i++
	fmt.Println(i)

	i--
	fmt.Println(i)

	fmt.Println("--------------------")

	// ****** LOGICAL OPERATORS ********
	var score int = 10
	fmt.Println((score < 100) && (score < 50))
	fmt.Println((score < 0) || (score < 200))

	var num1, num2 int = 10, 20
	fmt.Println(!(num1 > num2))
	fmt.Println(!(true))
	fmt.Println(!(false))

	// ****** ASSIGNMENT OPERATOR ******
	var word1, word2 string = "foo", "bar"
    word1 += word2
    fmt.Println(word1)

	var ab, bc int = 27, 7
    ab /= bc
    fmt.Println(ab)

	var cd, ef float64 = 27.9, 7.0
    cd -= ef
    fmt.Println(cd)
    cd += ef
    fmt.Println(cd)

	var fg, gh int = 100,9
    fg /= gh
    fmt.Println(fg)
    fg %= gh
    fmt.Println(fg)

	fmt.Println("--------------------")

	// ***** BITWISE OPERATORS
	var num4, numb5 int = 100,90
    fmt.Println(num4 & numb5)
    fmt.Println(num4 | numb5)

	var num6, num7 int = 100,90
    fmt.Println((num6+num7) >> 2)

	var num8, num9 int = 100,90
    fmt.Println(!(((num8+num9) >> 2 ) == 47))

	fmt.Println("--------------------")

	// ***** IF_ELSE AND ELSE IF STATEMENTS ******
	var mood string = "happy"
	var fruit string = "grapes"
	var car string = "Benz"

	if mood == "happy" {
		fmt.Println((mood))
	}
   
	if (fruit == "apples") {
		fmt.Println("Fruit is apple")
	} else {
		fmt.Println("Fruit is not apple")
	}

	if (car == "Benz") {
		fmt.Println("I love Benz")
	} else if car == "hyundai" {
		fmt.Println("Hyundias are not Benz")
	} else {
		fmt.Println("I love to drive")
	}

	fmt.Println("--------------------")

	// ********* SWITCH STATEMENT **********
	var j int = 100
	switch j {
		case 10:
			fmt.Println("j is 10")
		case 100, 200:
			fmt.Println("j is either 100 or 200")
		default:
			fmt.Println("j is neither 0, 100 or 200")
	}

	var zz int = 10
	switch zz {
		case -5:
			fmt.Println("-5")
		case 10:
			fmt.Println("10")
			fallthrough
		case 20:
			fmt.Println("20")
			fallthrough
		default:
			fmt.Println("default")
	}

	// NB: Go uses an implicit break keyword - so when two conditions match, it breaks out of the switch statement, without falling through to the next statement that matches
	var aaa, bbb int = 10, 20
	switch {
		case aaa+bbb == 30:
			fmt.Println("Equal to 30")

		case aaa+bbb <= 30:
			fmt.Println("Less than or equal to 30")
		
		default:
			fmt.Println("Greater than 30")
	}

	fmt.Println("--------------------")

	// ****** LOOPING WITH FOR ******
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("--------------------")

	// NB: The initialization and post statements are optional in Go
	ii := 1
	for ii <= 5 {
		fmt.Println(ii * 1)
		ii += 1
	}

	fmt.Println("--------------------")

	for i := 1; i <= 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}

	fmt.Println("--------------------")
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
}