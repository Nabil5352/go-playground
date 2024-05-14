package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	var x int = 10 // declare variable, mention name and type on left side of =
	y := 20        // shorter form of previous statement
	var t, z int   //default is zero

	var m, n = 15, "hi" // declare different types
	o, p := 15, "hola"  // Short form of previous line

	fmt.Printf("x: %d y:%d t:%d z:%d\n", x, y, t, z)
	fmt.Printf("m: %d n:%s\n", m, n)
	fmt.Printf("o: %d p:%s\n", o, p)

	const value = 22
	var i int = value
	var j float32 = value
	fmt.Printf("i: %d j:%f\n", i, j)

	/*
	* Assign max value to each variable (byte/uint8, init 32, uint64)
	* Then add 1 and see the effect
	 */
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615
	fmt.Printf("max range => b: %b smallI: %d bigI: %d\n", b, smallI, bigI)
	fmt.Printf("adding 1 => b: %b smallI: %d bigI: %d\n", b+1, smallI+1, bigI+1)

	/******************
	***** Array *******
	******************/

	/*
	* Create an array of size 6.
	* Undeclared values will be zero by default
	 */
	var arr = [6]int{2, 1, 3}
	fmt.Println(arr)

	/*
	* Create an array with no adjustable size.
	* Undeclared values will be zero by default
	 */
	var unlimited_arr = [...]int{4, 5, 9}
	fmt.Println(unlimited_arr)

	/*
	* Array of size 7
	* first 3 are declared 2,1,3
	* then, 5: means set 0 until 5th value
	* set 1 as 6th value
	* rest (if any) is zero by default
	 */
	var new_arr = [7]int{2, 1, 3, 5: 1}
	fmt.Println(new_arr, len(new_arr))

	/******************
	***** Slices *******
	******************/

	// Slices are same as array except here size declaratio isn't mandatory
	var sl = []int{1, 3, 5, 7, 9}
	var sl_1 = []int{1, 4: 5, 6: 7}
	fmt.Println(sl, sl_1)

	// Multidimensional Slice
	var sl_m_1 = [][]int{{1, 2}, {3, 4}}
	fmt.Println(sl_m_1[0][0])

	// Append
	sl_1 = append(sl_1, 8, 9)
	fmt.Println(sl_1)

	// Append another slice
	var sl_2 = []int{10, 11, 5: 15}
	sl_1 = append(sl_1, sl_2...)
	fmt.Println(sl_1)
}
