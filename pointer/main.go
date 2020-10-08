//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	myInt := 4
//	myIntPointer := &myInt
//	*myIntPointer = 8
//	fmt.Println(myIntPointer)
//	fmt.Println(*myIntPointer)
//
//	myFloat := 98.6
//	myFloatPoninter := &myFloat
//	fmt.Println(myFloatPoninter)
//	fmt.Println(*myFloatPoninter)
//}


//package main
//
//import "fmt"
//
//func main(){
//	myInt := 42
//	myIntPointer := &myInt
//	fmt.Println(*myIntPointer)
//}

package main

import "fmt"

func createPointer() *float64 {
	myFloat := 98.5
	return &myFloat
}

func printPointer(myBoolPointer *bool) {
	fmt.Println(*myBoolPointer)
}

func main() {
	myFloatPointer := createPointer()
	fmt.Println(*myFloatPointer)

	myBool := true
	printPointer(&myBool)
}