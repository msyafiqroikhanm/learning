package main

import "fmt"

func main() {
	fmt.Println("MAP")
	fmt.Println("==============================")

	//* Basic MAP
	fmt.Println("Basic MAP")
	fmt.Println("-----------------------------")
	var chicken map[string]int // tipe data map dengan key bertipe string dan value bertipe int.
	chicken = map[string]int{} // untuk ngasih inisialisasi nilai default

	chicken["january"] = 50
	chicken["february"] = 40

	fmt.Println("chicken:", chicken)
	fmt.Println("january:", chicken["january"])
	fmt.Println("may:", chicken["may"])
	fmt.Println()

	//*Map Initialization Value
	fmt.Println("Map Initialization Value")
	fmt.Println("-------------------------")
	var data map[string]int
	// data["one"] = 1
	//error because does not initialize yet

	data = map[string]int{}
	data["one"] = 1
	fmt.Println("data", data)

	var chicken1 = map[string]int{"januari": 50, "februari": 40}
	var chicken2 = map[string]int{
		"januari":  50,
		"februari": 40,
	}
	var chicken3 = map[string]int{}
	var chicken4 = make(map[string]int)
	var chicken5 = *new(map[string]int)
	fmt.Println("chicken1:", chicken1)
	fmt.Println("chicken2:", chicken2)
	fmt.Println("chicken3:", chicken3)
	fmt.Println("chicken4:", chicken4)
	fmt.Println("chicken5:", chicken5)
	fmt.Println()

	fmt.Println("Delete Map Item")
	fmt.Println("------------------------------")
	fmt.Println("len chicken:", len(chicken))
	fmt.Println("chicken:", chicken)

	delete(chicken, "january")

	fmt.Println("len chicken:", len(chicken))
	fmt.Println("chicken:", chicken)
	fmt.Println()

	// *Availability Key Detection in Map
	fmt.Println("Availability Key Detection in Map")
	fmt.Println("-----------------------------------")
	chicken["may"] = 30 //adding
	fmt.Println("chicken:", chicken)
	var value, isExist = chicken["januari"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("Item january is does not exist")
	}
}
