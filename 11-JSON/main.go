package main

import "fmt"

func SubTitle(subTitle string) {
	fmt.Println()
	fmt.Println(subTitle)
	for i := 0; i < len(subTitle)+5; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}

func main() {
	fmt.Println("JSON DATA")
	fmt.Println("=======================")
	Json()

	fmt.Println()
	fmt.Println("WEB SERVER")
	fmt.Println("=======================")
}
