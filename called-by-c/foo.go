package main

// C is a must
import (
	"fmt"

	"C"
)

//export HelloWorld
func HelloWorld() {
	fmt.Println("Hello World from Go")
}

func main() {}
