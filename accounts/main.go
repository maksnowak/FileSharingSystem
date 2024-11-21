package main

import "fmt"

func addOne(x int) int {
	return x + 1
}

func main() {
	fmt.Println("Hello from `accounts` microservice!")
	fmt.Println("Adding 1 to 3...")
	fmt.Println(addOne(3))
	fmt.Println("Wakatime test")

}
