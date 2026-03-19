package main

import "fmt"

// ***** ARRAYS *****
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

	fmt.Println("----------------------")

	// ****** Loop through an array
	for i := 0; i < len(scores); i++ {
		fmt.Println(scores[i])
	}

	fmt.Println("----------------------")

	// **** Using the range keyword
	for index, element := range nums {
		fmt.Println(index, "=>", element)
	}

	fmt.Println("---------------------")

	// ***** Multidimensional array
	arr := [3][2]int{{2, 4}, {4, 16}, {8, 64}}
	fmt.Println(arr[2][1]) // -> 64

	// ******* SLICES *********
    // . continuous segment of an underlying array
	// . variable typed (elements can be added or removed)
	// . more flexible
	// . slice has three major components: pointer, length, and capacity
	/*
	  Declaring and initializing a slice:
	  <slice_name> := []<data_type>{<values>}
	*/
	/*
	 Rule:
		Capacity = length of underlying array − starting index of the slice
	*/
	/* 
		slices can be created from arrays. This is how we create a slice from an array:
		array[start_index : end_index]
		the element in th start_index is included, while the element in the end_index IS NOT included
	*/
	/*
	  array[ : 4] // the start_index is zero by default here
	  array[ : ] // since no start_index and end_index, the complete array will be sliced here for us
	*/

	fmt.Println("--------------------")

	slice := []int{10, 20, 30}
	fmt.Println(slice)

	cards := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	slice_1 := cards[1:8]
	fmt.Println(slice_1)

	sub_slice := slice_1[0:3] // creates a slice from a slice
	fmt.Println(sub_slice)

	/*
	 Another way of creating a slice is by using the make function:
	 	slice := make([]<data_type>, length, capacity)
		capacity is optional here
	*/

	fmt.Println("--------------------")

	slice_two := make([]int, 5, 8)
	fmt.Println(slice_two)
	fmt.Println(len(slice_two)) // prints length
	fmt.Println(cap(slice_two)) // prints capacity

	/*
	  When you change the value of an element in a slice, it changes that value in the underlying array itself. Since a slice is nothing but a reference to the underlying array
	*/
	fmt.Println("---------------------")

	arr_again := [5]int{10, 20, 30, 40, 50}
	slice_again := arr_again[:3]

	fmt.Println(arr_again)
	fmt.Println((slice_again))

	slice_again[1] = 9000

	fmt.Println(slice_again)

	avava := [4]int{100, 200, 300, 400}
	fmt.Println(avava)

	// The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice

	sliceavava := avava[1:3]
	fmt.Println(sliceavava)
		fmt.Println(len(sliceavava))
			fmt.Println(cap(sliceavava))

	/* Appending elements to a slice:
	   Built in function: func append(s []T, vs ...T) []T
		s []T is a slice of some data type
		vs ...T are the values of the same data type
		return []T which is a slice containing all the elements of the original slice, plus the provided values

	   i.e slice = append(slice, value1, value2..)

	*/

	arr_three := [4]int{40, 80, 120, 160}
	slice_three := arr_three[1:3]
	fmt.Println("--------- Slice append ---------")
	fmt.Println(slice_three)
	fmt.Println(len((slice_three)))
	fmt.Println(cap(slice_three))

	fmt.Println("--------- Appending 1 element ---------")
	slice_four := append(slice_three, 180)
	fmt.Println(slice_four)
	fmt.Println(len(slice_four)) // length becomes 3
	fmt.Println(cap(slice_four)) // Capacity remains 3

	fmt.Println("--------- Appending 1 element again ---------")
	slice_five := append(slice_four, 200)
	fmt.Println(len(slice_five)) // length becomes 4
	fmt.Println(cap(slice_five)) // Capacity doubles because we have exceeded the capacity of the original slice, so it becomes 6

	// ******** Appending to a slice ***********
	// slice = append(slice, anotherSlice...)

	scored_goals := [5]int{10, 20, 30, 40, 50}
	slice_of_goals := scored_goals[:2]

	scored_goals_two := [5]int{5, 15, 25, 35, 45}
	slice_of_goals_two := scored_goals_two[:2]

	new_slice := append(slice_of_goals, slice_of_goals_two...)
	fmt.Println("--------- Appending to a slice ---------")
	fmt.Println(new_slice) // [10, 20, 5, 15]

	// ******** Deleting from a slice ************
	funds := [5]int{10, 20, 30, 40, 50}
	i := 2 // delete element at index 2 i.e 30

	fmt.Println("--------- Deleting an element from a slice ---------")
	fmt.Println(funds)

	slice_funds_1 := funds[:i] // [10, 20]
	slice_funds_2 := funds[i + 1:] // [40, 50]

	final_funds_slice := append(slice_funds_1, slice_funds_2...)
	fmt.Println(final_funds_slice)

	// ********* Copying from a slice ***********
	// func copy(dst. src []Type) int
	// num := copy(destination_slice, source_slice)
	// NB: For the copy function to work, the slices must be initialized with the same data type

	source_slice := []int{10, 20, 20, 40, 50}
	destination_slice := make([]int, 3)
	num_of_elements_copied := copy(destination_slice, source_slice)
	fmt.Println(destination_slice)
	fmt.Println("Number of elements num_of_elements_copied", num_of_elements_copied)

	// ******* Looping through a slice ********
	// This is the same as looping through an array
	// For use cases where we do not need the index, we can replace the index with underscore i.e _.

	fmt.Println("--------- Looping through a slice ---------")
	for _, value := range source_slice {
		fmt.Println(value)
	}

	// ***** MAPS *****
	// . Unordered collection of key/value pairs
	// . Implemented by hash tables
	// . Provide efficient add, get and delete operations

	/* declaring and initialzing a map:
	 var <map_name> map[<key_data_type>]<value_data_type>
	 e.g var my_map map[string]int --> this line of code creates a nil map
	 The zero value of the map is nil, and the nil map does not contain any keys
	 Trying to add a value to a nil map makes the compiler throw an error
	*/

	fmt.Println("--------- Maps ---------")
	// var codes map[string]string // Creates a nil map
	// codes["en"] = "English"
	// fmt.Println(codes)

	codes_map := map[string]string{"en": "English", "fr": "French"}
	fmt.Println(codes_map)
	fmt.Println(len(codes_map))

	// using a make() function
	/* <map_name> := 
	    make(map[<key_data_type>]<value_data_type, <initial_capacity>)
	 NB: the initial_capacity is optional
	*/

	empty_codes_map := make(map[string]int)
	fmt.Println(empty_codes_map) // Prints an empty map

	// Accessing a map i.e map[key]
	fmt.Println(codes_map["en"])
	fmt.Println(codes_map["fr"])

	// getting a key -> getting the value associated with the key
	// value, found := map_name[key]
	codes_map_nums := map[string]int{"en": 1, "fr": 2, "hi": 3}
	value, found := codes_map_nums["en"]
	fmt.Println(value, found) // Prints: 1 true

	value2, found2 := codes_map_nums["hh"]
	fmt.Println(value2, found2) // Prints: 0 false

	// Add key-value pair
	codes_map["it"] = "Italian"
	fmt.Println(codes_map)

	// Update ke-value pair
	codes_map["en"] = "English Language"
	fmt.Println(codes_map)

	// delete a key-value pair
	delete(codes_map, "en")
	fmt.Println(codes_map)

	fmt.Println("--------- Loop over a map ---------")
	// Looping over a map 
	for key, value := range codes_map {
		fmt.Println(key, "=>", value)
	}

	// truncate a map i.e clearing alll elements in a map
	// method 1
	for key, _ := range codes_map {
		delete(codes_map, key)
	}
	fmt.Println(codes_map)

	// method 2 i.e simply initialize it with an empty map
	codes_map = make(map[string]string)
}