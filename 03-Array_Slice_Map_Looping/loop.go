package main

import "fmt"

func main() {
	fmt.Println("LOOPING")
	fmt.Println("===========================")

	// * For Loop
	fmt.Println("For Loop")
	fmt.Println("----------------------------")
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}
	fmt.Println()

	//* For with only condition
	fmt.Println("For with only condition")
	fmt.Println("---------------------------")
	i := 0
	for i < 5 {
		fmt.Println("Angka", i)
		i++
	}
	fmt.Println()

	//* For Without Argument
	fmt.Println("For Without Argument")
	fmt.Println("--------------------------")
	i = 0
	for {
		fmt.Println("Angka", i)
		i++
		if i == 5 {
			break
		}
	}
	fmt.Println()

	// * For - Range Loop
	fmt.Println("For - Range Loop")
	fmt.Println("--------------------------")
	var fruits = [4]string{"apple", "grape", "banana", "melon"}
	for i, fruit := range fruits {
		fmt.Printf("element %d : %s\n", i, fruit)
	}
	fmt.Println()

	//* Using underscore (_) in For - Range
	fmt.Println("Using underscore (_) in For - Range")
	fmt.Println("-----------------------------------")
	for _, fruit := range fruits {
		fmt.Printf("nama buah : %s\n", fruit)
	}
	fmt.Println()

	// * Break & Continue
	fmt.Println("Break & Continue")
	fmt.Println("----------------")
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}
		if i > 8 {
			break
		}
		fmt.Println("Angka", i)
	}
}
