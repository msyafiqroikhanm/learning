package main

import (
	"fmt"
	"math"
	"strings"
)

func subTitle(subTitle string) {
	fmt.Println()
	fmt.Println(subTitle)
	for i := 0; i < len(subTitle)+5; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}

type hitung interface {
	luas() float64
	keliling() float64
}

type lingkaran struct {
	diameter float64
}

func (l lingkaran) jariJari() float64 {
	return l.diameter / 2
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}

// * Embedded Interface
type hitung2d interface {
	luas() float64
	keliling() float64
}

type hitung3d interface {
	volume() float64
}

type perhitungan interface {
	hitung2d
	hitung3d
}

type kubus struct {
	sisi float64
}

func (k kubus) volume() float64 {
	return math.Pow(k.sisi, 3)
}

func (k kubus) luas() float64 {
	return math.Pow(k.sisi, 2) * 6
}

func (k kubus) keliling() float64 {
	return k.sisi * 12
}

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println("INTERFACE")
	fmt.Println("============")

	/*
		Interface adalah kumpulan definisi method yang tidak memiliki isi (hanya definisi saja),
		yang dibungkus dengan nama tertentu.

		Interface merupakan tipe data.
		Nilai object bertipe interface zero value-nya adalah nil.
		Interface mulai bisa digunakan jika sudah ada isinya,
		yaitu object konkret yang memiliki definisi method minimal sama dengan yang ada di interface-nya.

		Jadi interface itu kayak definisi method yang nantinya diperkirakan ada beberapa method yang sama namanya.
	*/

	subTitle("Interface Implementation")
	var bangunDatar hitung

	bangunDatar = persegi{10.0}
	fmt.Println("======= persegi")
	fmt.Println("Luas \t\t:", bangunDatar.luas())
	fmt.Println("Keliling \t:", bangunDatar.keliling())

	bangunDatar = lingkaran{14.0}
	fmt.Println("======= lingkaran")
	fmt.Println("Jari-jari \t:", bangunDatar.(lingkaran).jariJari())
	fmt.Println("Luas \t\t:", bangunDatar.luas())
	fmt.Println("Keliling \t:", bangunDatar.keliling())

	subTitle("Casting Interface")
	var bangunDatarLingkaran lingkaran = bangunDatar.(lingkaran)
	fmt.Println("Casting Jari-jari :", bangunDatarLingkaran.jariJari())

	subTitle("Embedded Interface")
	/*
		Interface bisa di-embed ke interface lain, sama seperti struct.
		Cara penerapannya juga sama, cukup dengan menuliskan nama interface yang ingin di-embed ke dalam interface tujuan.
	*/
	var bangunRuang perhitungan = kubus{4}
	fmt.Println("======= kubus")
	fmt.Println("Luas    :", bangunRuang.luas())
	fmt.Println("Keliling:", bangunRuang.keliling())
	fmt.Println("Volume  :", bangunRuang.volume())

	subTitle("Anonymous Interface")
	var secret interface{}

	secret = "ethan hunt"
	fmt.Println(secret)

	secret = []string{"apple", "mango", "banana"}
	fmt.Println(secret)

	secret = 12.4
	fmt.Println(secret)

	subTitle("Casting Variable in Anonymous Interface")
	var secret2 interface{}
	secret2 = 2
	var number = secret2.(int) * 10
	fmt.Println(secret2, "multiplied by 10 :", number)

	secret2 = []string{"apple", "mango", "banana"}
	var fruits = strings.Join(secret2.([]string), ", ")
	fmt.Println(fruits, "is my favorite fruits")

	subTitle("Casting Anonymous Interface into Pointer")
	var secreto interface{} = &person{name: "Wick", age: 23}
	var name = secreto.(*person).name
	fmt.Println(name)
}
