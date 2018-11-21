package main

import "fmt"
import "time"

func main() {
	for i := 10; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
	fmt.Println("This is a for loop in go created by ANAND PRABHAKAR")
}
