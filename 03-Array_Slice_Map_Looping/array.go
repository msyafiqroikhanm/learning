package main

import "fmt"

func main() {
	fmt.Println("ARRAYY")
	fmt.Println("============================")
	// basic array
	var names [4]string
	names[0] = "John"
	names[1] = "Doe"
	names[2] = "Frank"
	names[3] = "Jack"

	fmt.Println(names[0], names[1], names[2], names[3])
	fmt.Println()

	// inisialisasi nilai awal array
	var newNames = [4]string{"John", "Doe", "Frank", "Jack"}
	fmt.Println(newNames[0], newNames[1], newNames[2], newNames[3])

	fruits := [4]string{
		"Apple",
		"Grape",
		"Banana",
		"Melon",
	}
	fmt.Println(fruits)
	fmt.Println()

	// inisialisasi array tanpa jumlah elemen
	var numbers = [...]int{2, 3, 2, 4, 3}
	fmt.Println("Data Array \t:", numbers)
	fmt.Println("Jumlah Elemen \t:", len(numbers))
	fmt.Println()

	// multidimentional array
	var numbers1 = [2][3]int{
		[3]int{3, 2, 3},
		[3]int{3, 4, 5},
	}
	var numbers2 = [2][3]int{
		{3, 2, 3},
		{3, 4, 5},
	}
	fmt.Println("Array Numbers1", numbers1)
	fmt.Println("Array Numbers2", numbers2)
}
