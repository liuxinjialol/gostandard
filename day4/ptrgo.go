package main

import (
    "fmt"
)

// Multiply this function changes reply:
func Multiply(a, b int, reply *int) {
    *reply = a * b
}

func main() {
    n := 0
    reply := &n
    Multiply(10, 5, reply)
	fmt.Println("Multiply:", *reply) // Multiply: 50
	
    fmt.Println("n=", n) // Multiply: 50
    

    capitals := map[string] string {"France":"Paris", "Italy":"Rome", "Japan":"Tokyo" }
    for key := range capitals {
	    fmt.Println("Map item: Capital of", key, "is", capitals[key])
    }

}