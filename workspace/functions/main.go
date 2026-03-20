package main

import "fmt"

func addNumbers(a int, b int) int {
	sum := a + b
	fmt.Println(sum)
	return  sum
}

func greeting(name string) {
	fmt.Println("Hey there,", name)
}

func operation(a int, b int) (int, int) {
	sum := a + b
	diff := a - b
	return sum, diff
}

// named return function -> generally used when a function has to return multiple values
func operation_2(a int, b int) (sum int, diff int) {
	// by using named return values. we declare the variables using just an equal sign, and we can only use the return keyword at the end of the function
	sum = a + b
	diff = a - b
	return
}

// Variadic functions -> can be called with 0 or more arguments
func sumNumbers(numbers ...int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}

func printDetails(student string, subjects ...string) {
	fmt.Println("Hey ", student, ", here are your subjects - ")
	for _, sub := range subjects {
		fmt.Printf("%s, ", sub)
	}
	fmt.Println("------")
}

// Blank identifier i.e _. There are used to ignore the values that are returned by a function
func f() (int, int) {
	return 42, 53
}

// Recursive function
func factorial(n int) int {
	// Base case
	if n == 1 {
		return 1
	}
	return  n * factorial(n - 1)
}

// Anonymous function -- details are in the main function

// Higher order functions
func calcArea(r float64) float64 {
	return 3.14 * r * r
}

func calcPerimeter(r float64) float64 {
	return 2 * 3.14 * r
}

func calcDiameter(r float64) float64 {
	return  2 * r
}

func printResult(radius float64, calcFunction func(r float64) float64) {
	result := calcFunction(radius)
	fmt.Println("Result: ", result)
	fmt.Println("Thank you!")
}

func getFunction (query int) func(r float64) float64 {
	query_to_func := map[int]func(r float64) float64 {
		1: calcArea,
		2: calcPerimeter,
		3: calcDiameter,
	}
	return query_to_func[query]
}

// defer statement -> displays the execution of a function, until the surrounding function returns
func printName(str string) {
	fmt.Println(str)
}

func printRollNo(rno int) {
	fmt.Println(rno)
}

func printAddress(addr string) {
	fmt.Println(addr)
}


func main() {
	/*

	*/

	addNumbers(2, 5)
	greeting("John")

	// **** function return types *****
	sum, difference := operation(20, 10)
	fmt.Println(sum, " ", difference)

	// ****** named return values
	sum_2, difference_2 := operation_2(50, 10)
	fmt.Println("Using named return values: ", sum_2, difference_2)

	// ***** variadic functions *****
	fmt.Println("--------- Variadic function calls ---------")
	fmt.Println(sumNumbers())
	fmt.Println(sumNumbers(10))
	fmt.Println(sumNumbers(10, 20))
	fmt.Println(sumNumbers(10, 20, 30, 40, 50))

	printDetails("Joe", "Physics", "Biology")

	// ***** Blank identifier ********
	fmt.Println("--------- Blank identifier ---------")
	v, _ := f()
	fmt.Println(v)

	// ******* Recursive function ***********
	fmt.Println("--------- Recursive function ---------")
	n := 5
	result := factorial(n)
	fmt.Println("Factorial of", n, "is :", result)

	fmt.Println("--------- Anonymous function ---------")
	x := func (a int, b int) int  {
		return a * b
	}
	fmt.Printf("%T \n", x) // prints the data type
	fmt.Println(x(20, 30))

	fmt.Println("----------------")

	y := func (a int, b int) int {
		return a * b
	}(40, 50)
	fmt.Printf("%T \n", y)
	fmt.Println(y)

	// Higher Order function
	fmt.Println("--------- Higher order function ---------")
	var query int
	var radius float64

	fmt.Print("Enter the radius of the circle: ")
	fmt.Scanf("%f\n", &radius)
	fmt.Printf("Enter \n 1 - Area \n 2 - Perimeter \n 3 - Diameter: ")
	fmt.Scanf("%d\n", &query)

	printResult(radius, getFunction(query))

	// defer statement
	fmt.Println("--------- defer statement ---------")
	printName("Joe")
	defer printRollNo(23) // this delays the execution of this function until the surrounding functions return
	printAddress("street-32")
}