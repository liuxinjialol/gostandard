package main

import "fmt"
import "strings"
import "os"

func main(){

	s:="   this is my desk   "
	fmt.Println(s)

	fmt.Println(strings.TrimSpace(s))

	os.Exit(1)
}