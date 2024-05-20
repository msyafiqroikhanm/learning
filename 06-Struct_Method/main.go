package main

import (
	"fmt"
)

// * Struct Declaration
type student struct {
	name  string
	grade int
}

// * Embedded Struct
// struct 1
type person struct {
	name string
	age  int
}

// struct 2
type lecturer struct {
	grade int
	person
}

// * Nested Struct
type vehicle struct {
	bicycle struct {
		name  string
		wheel int
	}
	grade int
}

func subTitle(subTitle string) {
	fmt.Println()
	fmt.Println(subTitle)
	for i := 0; i < len(subTitle)+5; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}

func main() {
	fmt.Println("STRUCT")
	fmt.Println("=========================")
	/*
		Struct adalah kumpulan definisi variabel (atau property) dan atau fungsi (atau method),
		yang dibungkus sebagai tipe data baru dengan nama tertentu.
		Property dalam struct, tipe datanya bisa bervariasi.
		Mirip seperti map, hanya saja key-nya sudah didefinisikan di awal,
		dan tipe data tiap itemnya bisa berbeda.
	*/

	subTitle("Struct Implementation")
	var john student // struct declaration
	// struct property initialization
	john.name = "John Doe"
	john.grade = 2

	fmt.Println("name \t:", john.name)
	fmt.Println("grade \t:", john.grade)

	subTitle("Struct Literals Declaration")
	// first way
	var syafiq = student{}
	syafiq.name = "Syafiq Roikhan Maulana"
	syafiq.grade = 15

	// second way
	var roikhan = student{"Roikhan Syafiq", 13}

	// third way
	var maulana = student{name: "Maulana", grade: 21}

	fmt.Println("student 1:", syafiq)
	fmt.Println("student 2:", roikhan)
	fmt.Println("student 3:", maulana)

	subTitle("Embedded Struct")
	/*
		Embedded struct adalah mekanisme untuk menempelkan sebuah struct sebagai properti struct lain
	*/
	var lecturerone = lecturer{}
	lecturerone.grade = 2

	// dia bisa langsung initialize property yang ada di embedded struct nya
	// dua property ini dia ada di struct person soalnya
	lecturerone.name = "Syafiq Lecturer"
	lecturerone.age = 23

	fmt.Println("name \t:", lecturerone.name)
	fmt.Println("grade \t:", lecturerone.grade)
	// dan ketika diprint juga value nya sama. terbukti dengan hasil print di bawah ini
	fmt.Println("age \t:", lecturerone.age)
	fmt.Println("age-p \t:", lecturerone.person.age)
	fmt.Println()

	var lecturerPerson = person{name: "Maulana", age: 21}
	var lecturerData = lecturer{person: lecturerPerson, grade: 2}

	fmt.Println("name \t:", lecturerData.name)
	fmt.Println("age \t:", lecturerData.age)
	fmt.Println("grade \t:", lecturerData.grade)

	subTitle("Anonymous Struct")
	var johnDoe = struct {
		person
		grade int
	}{}
	johnDoe.person = person{"JohnDooe", 21}
	johnDoe.grade = 2
	fmt.Println("name  :", johnDoe.person.name)
	fmt.Println("age   :", johnDoe.person.age)
	fmt.Println("grade :", johnDoe.grade)
	fmt.Println()

	// Anonymous Struct without property initiallization
	var mr = struct {
		person
		grade int
	}{}

	// Anonymous Struct with property initialization
	var wick = struct {
		person
		grade int
	}{
		person: person{"wick", 21},
		grade:  1,
	}

	fmt.Println("mr \t:", mr)
	fmt.Println("wick \t:", wick)

	subTitle("Nested Struct")
	mybicycle := vehicle{}
	mybicycle.grade = 20
	mybicycle.bicycle.name = "Mountain Bicycle"
	mybicycle.bicycle.wheel = 2
	fmt.Println("name  :", mybicycle.bicycle.name)
	fmt.Println("age   :", mybicycle.bicycle.wheel)
	fmt.Println("wheel :", mybicycle.grade)

}
