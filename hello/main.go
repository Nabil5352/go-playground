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
}
