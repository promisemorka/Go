package main

import "fmt"

// ***** Arrays *****
//  . fixed length. It denotes the number of elements in the array
//  . capacity - This denotes the number of elements it can contain
// . In the case of arrays, length and capacity are the same
//  . elements should be of the same data type -  homogenous
//  . elements are stroed at contiguous memory locations

func main() {
	var grades [5]int
	fmt.Println(len(grades))
	fmt.Println(grades)

	var fruits [3] string
	fmt.Println(fruits)

	var scores [3]int = [3]int{70, 80, 94}
	fmt.Println(scores)

	points := [3]int{83, 66, 99}
	fmt.Println(points)

	nums := [...]int{6, 9, 12} // using elipsis, the compiler will implicitly determine the length of the array
	fmt.Println(nums)
	fmt.Println(len(nums))

	fmt.Println(nums[2])

	nums[2] = 22
	fmt.Println(nums)

	fmt.Println("=====================")

	// ****** Loop through an array
	for i := 0; i < len(scores); i++ {
		fmt.Println(scores[i])
	}

	fmt.Println("=====================")

	// **** Using the range keyword
	for index, element := range nums {
		fmt.Println(index, "=>", element)
	}

	fmt.Println("=====================")

	// ***** Multidimensional array
	arr := [3][2]int{{2, 4}, {4, 16}, {8, 64}}
	fmt.Println(arr[2][1]) // -> 64

	// ******* SLICES *********
    // . continuous segment of an underlying array
	// . variable typed (elements can be added or removed)
	// . more flexible
	// . slice has three major components: pointer, length, and capacity\
	/* 
		slices are created from arrays. This is how we create a slice
		array[start_index : end_index]
		the element in th start_index is included, while the element in the end_index IS NOT included
	*/

	fmt.Println("=====================")

	slice := []int{10, 20, 30}
	fmt.Println(slice)

	cards := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	slice_1 := cards[1:8]
	fmt.Println(slice_1)

	sub_slice := slice_1[0:3] // creates a slice from a slice
	fmt.Println(sub_slice)

	/*
	 Another way of creating a slice:
	 	slice := make([]<data_type>, length, capacity)
	*/

	fmt.Println("=====================")

	slice_two := make([]int, 5, 8)
	fmt.Println(slice_two)
	fmt.Println(len(slice_two))
	fmt.Println(cap(slice_two))
}