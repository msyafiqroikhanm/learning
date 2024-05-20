package main

import "fmt"

/*
Pointer adalah reference atau alamat memori. Variabel pointer berarti variabel yang berisi alamat memori suatu nilai. Sebagai contoh sebuah variabel bertipe integer memiliki nilai 4, maka yang dimaksud pointer adalah alamat memori dimana nilai 4 disimpan, bukan nilai 4 itu sendiri.

Variabel-variabel yang memiliki reference atau alamat pointer yang sama, saling berhubungan satu sama lain dan nilainya pasti sama. Ketika ada perubahan nilai, maka akan memberikan efek kepada variabel lain (yang referensi-nya sama) yaitu nilainya ikut berubah.
*/

func subTitle(subTitle string) {
	fmt.Println()
	fmt.Println(subTitle)
	for i := 0; i < len(subTitle)+5; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}

func change(original *int, value int) {
	// fmt.Println("Value", *original)
	*original = value // value aselinya kalo *original
}

func changeWithNoPointer(original, value int) {
	original = value
}

func main() {
	fmt.Println("POINTER")
	fmt.Println("========")
	// memungkinkan perubahan variable di luar function.
	// kalau ga pakai pointer, maka tidak akan berubah.

	subTitle("Basic Pointer")
	var numberA int = 4
	var numberB *int = &numberA // int = fungsinya untuk deklarasi pointer. &numberA = fungsinya untuk mengambil nilai reference dari variable numberA
	fmt.Println("NumberA (Value) \t:", numberA)
	fmt.Println("NumberA (Address) \t:", numberA)
	fmt.Println("NumberB (Value) \t:", numberB)
	fmt.Println("NumberB (Address) \t:", numberB)

	subTitle("Change Pointer Value Effect")
	// Ketika suatu variable pointer diubah nilainya, sedangkan ada variable lain yang memiliki referensi memori yang sama, maka nilainya akan berubah
	numberA = 5
	fmt.Println("NumberA (Value) \t:", numberA)
	fmt.Println("NumberA (Address) \t:", numberA)
	fmt.Println("NumberB (Value) \t:", numberB)
	fmt.Println("NumberB (Address) \t:", numberB)

	subTitle("Doesnot change while not use pointer")
	var number = 4
	fmt.Println("Before No Pointer \t:", number)
	changeWithNoPointer(number, 10)
	fmt.Println("After No Pointer \t:", number)

	subTitle("Pointer Parameters")
	fmt.Println("Before \t:", number)

	change(&number, 10)
	fmt.Println("After \t:", number)

}
