package main

import "fmt"

func printHello() {
	fmt.Println("Hello")
}

func subTitle(subTitle string) {
	fmt.Println()
	fmt.Println(subTitle, len(subTitle))
	for i := 0; i < len(subTitle)+5; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}

func printAngka(angka1 int, angka2 int) {
	fmt.Println("Angka Pertama:", angka1)
	fmt.Println("Angka Kedua:", angka2)
}

func introduction(name string) string {
	return "Hello my name is " + name
}

func secondIntroduction(firstName string, lastName string) (string, string) {
	introFirstName := "Hello My First Name Is " + firstName
	introLastName := "Hello My Last Name Is " + lastName
	return introFirstName, introLastName
}

func tambahAngka(firstNumber int, lastNumber int) (jumlah int) {
	jumlah = firstNumber + lastNumber
	return
}

func tampilkanKataDanAngka() (firstWord, secondWord string, number int) {
	firstWord = "Halo"
	secondWord = "Dunia"
	number = 10
	return
}

func sum(numbers ...int) int {
	var total int = 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func yourHobbies(name string, hobbies ...string) {
	fmt.Println("Hello My Name Is", name)
	fmt.Println("My Hobbies are: ")
	for _, hobby := range hobbies {
		fmt.Println(hobby)
	}
}

func findMax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}

	return len(res), func() []int {
		return res
	}
}

// Ini Function yang makai function sebagai parameternya
func sayHelloWithFilter(name string, filter func(string) string) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

// ini Function yang nantinya akan digunakan sebagai parameter
func spamFilter(name string) string {
	if name == "Kasar" {
		return "..."
	}
	return name
}

func main() {
	fmt.Println("FUNCTION")
	fmt.Println("=================")
	// * Basic Function
	subTitle("Basic Function")
	printHello()

	// * Function with parameters
	subTitle("Function With Parameters")
	printAngka(1, 2)

	// * Function Return Value
	subTitle("Function Return Value")

	// langsung diprint
	fmt.Println(introduction("Syafiq"))

	// disimpan dulu di variable lalu diprint
	result := introduction("Roikhan")
	fmt.Println(result)

	//function as value
	secondResult := introduction
	fmt.Println(secondResult("Maulana"))

	// * Function Return Multiple Value
	subTitle("Function Return Multiple Value")
	firstName, lastName := secondIntroduction("Syafiq", "Maulana")
	fmt.Println(firstName, lastName)

	firstName1, _ := secondIntroduction("Syafiq", "Maulana")
	fmt.Println(firstName1)

	// * Function Predefined Return Value
	subTitle("Function Predefined Return Value")
	hasil := tambahAngka(4, 5)
	fmt.Println(hasil)

	kataPertama, kataKedua, angka := tampilkanKataDanAngka()
	fmt.Println(kataPertama, kataKedua, angka)

	// * Variadic Function
	subTitle("Variadic Function")
	/*
		konsep variadic function atau pembuatan fungsi dengan parameter sejenis yang tak terbatas. Maksud tak terbatas disini adalah jumlah parameter yang disisipkan ketika pemanggilan fungsi bisa berapa saja.

		Parameter variadic memiliki sifat yang mirip dengan slice. Nilai dari parameter-parameter yang disisipkan bertipe data sama, dan ditampung oleh sebuah variabel saja. Cara pengaksesan tiap datanya juga sama, dengan menggunakan index.

		Deklarasi parameter variadic sama dengan cara deklarasi variabel biasa, pembedanya adalah pada parameter jenis ini ditambahkan tanda titik tiga kali (...) tepat setelah penulisan variabel (sebelum tipe data). Nantinya semua nilai yang disisipkan sebagai parameter akan ditampung oleh variabel tersebut
	*/

	// basic variadic
	var total = sum(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	fmt.Println("total with basic number:", total)

	// variadic initialization with slice datatype
	var numbers = []int{2, 6, 7, 8, 9, 10}
	total = sum(numbers...)
	fmt.Println("total with slice datatype:", total)

	// * Function With Basic and Variadic Parameters
	subTitle("Function With Basic and Variadic Parameters")
	yourHobbies("Syafiq", "Gaming", "Sleeping", "Eating")
	fmt.Println()

	var hobbies = []string{"Sleeping", "Eating"}
	yourHobbies("Maulana", hobbies...)

	// * Closure Function
	subTitle("Closure Function")
	/*
		Definisi Closure adalah sebuah fungsi yang bisa disimpan dalam variabel.
		Dengan menerapkan konsep tersebut, kita bisa membuat fungsi didalam fungsi,
		atau bahkan membuat fungsi yang mengembalikan fungsi.

		Closure merupakan anonymous function atau fungsi tanpa nama.
		Biasa dimanfaatkan untuk membungkus suatu proses yang hanya dipakai sekali atau dipakai pada blok tertentu saja.
	*/

	// Closure saved as variable
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}

	var min, max = getMinMax(numbers)
	fmt.Println("Data Numbers:", numbers)
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
	fmt.Println()

	// Closure As Return Value
	max = 3
	numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
	var howMany, getNumbers = findMax(numbers, max)
	var theNumbers = getNumbers()
	fmt.Println("Numbers\t:", numbers)
	fmt.Printf("Find \t: %d \n\n", max)

	fmt.Println("found \t:", howMany)    // 9
	fmt.Println("value \t:", theNumbers) // [2 3 0 3 2 0 2 0 3]

	//* Function As Parameter
	subTitle("Function As Parameter")
	sayHelloWithFilter("Syafiq", spamFilter)
	sayHelloWithFilter("Kasar", spamFilter)
}
