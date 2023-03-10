package main

import "fmt"

func main() {
	i, j := 987, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 55         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 44   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
