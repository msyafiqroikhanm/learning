package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// *first
	fmt.Println("Hello World")
	fmt.Println("=============")
	fmt.Println()

	fmt.Println("VARIABLE")
	//* Variable
	// with var
	var text = "Hello World"
	fmt.Println(text)

	// with datatype
	var textString string
	textString = "ini teks 1"
	fmt.Println(textString)

	var textString2 string = "Ini teks string 2"
	textString2 = "Ini teks string 2 telah dirubah"
	fmt.Println(textString2)

	// with :=
	variable1 := "ini variable 1" // ini deklarasi dan inisiasi
	variable1 = "ini variable 1 yang telah diubah"
	fmt.Println(variable1)

	variable2 := "Ini variable 2"
	fmt.Println(variable2)
	fmt.Println("=============")
	fmt.Println()

	//* CONSTANT
	fmt.Println("CONSTANT")
	const judul = "Ini judul"
	// judul = "judul diubah" // judul gabisa diubah karna dia constant
	fmt.Println(judul)
	fmt.Println("==================")
	fmt.Println()

	//* DATA TYPE
	fmt.Println("DATA TYPE")
	// numeric non decimal
	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644

	fmt.Printf("Bilangan Positif: %d\n", positiveNumber)
	fmt.Printf("Bilangan Negatif: %d\n", negativeNumber)

	//numeric decimal
	var decimalNumber = 2.62
	fmt.Printf("Bilangan Decimal: %f\n", decimalNumber)
	fmt.Printf("Bilangan Decimal: %.3f\n", decimalNumber)

	//boolean
	var exist bool = true
	fmt.Printf("Exist? %t \n", exist)

	//string
	var message string = "halo"
	fmt.Printf("message: %s \n", message)

	var message2 string = `Nama Saya "John Wick". 
	Salam Kenal.
	Mari Belajar "Golang".`
	fmt.Println(message2)

	var name = "John Doe"
	var age = "28"
	var sentence = `halo nama saya "` + name + `" dan berumur "` + age + `"`
	fmt.Println(sentence)

	fmt.Println("============")
	fmt.Println()

	//* CASTING
	fmt.Println("CASTING")
	var floatData float64 = float64(24)
	fmt.Printf("float data = %.2f \n", floatData)

	var intData int32 = int32(24.00)
	fmt.Printf("int data = %d \n", intData)

	var str = "halo"
	var stringData = string(str[0])
	fmt.Printf("string data = %s \n", stringData)
	fmt.Println("==============")
	fmt.Println()

	//* PACKAGE
	fmt.Println("PACKAGE")
	fmt.Println("-----------")
	fmt.Println("STRING")
	// string
	// index
	var index1 = strings.Index("ethan hunt", "ha")
	fmt.Println(index1)
	var index2 = strings.Index("ethan hunt", "n")
	fmt.Println(index2)
	//replace
	var fruit = "banana"
	var find = "a"
	var replaceWith = "o"
	var newText1 = strings.Replace(fruit, find, replaceWith, 1)
	fmt.Println(newText1)
	var newText2 = strings.Replace(fruit, find, replaceWith, 2)
	fmt.Println(newText2)
	var newText3 = strings.Replace(fruit, find, replaceWith, -1) // semua substring
	fmt.Println(newText3)
	//repeat
	var strRepeat = strings.Repeat("Syafiq", 4)
	fmt.Println(strRepeat)
	//toLower
	var strToLower = strings.ToLower("M. Syafiq Roikhan Maulana")
	fmt.Println(strToLower)
	//toUpper
	var strToUpper = strings.ToUpper("indonNesia Raya")
	fmt.Println(strToUpper)

	fmt.Println("-----------")
	fmt.Println("STRCONV")
	//atoi = konversi dari string ke int
	var strA = "1234"
	var num, err = strconv.Atoi(strA)
	if err == nil {
		fmt.Println(num)
	}
	//itoa = conversi dari integer ke string
	var numericDataType = 12345
	strItoa := strconv.Itoa(numericDataType)
	fmt.Println(strItoa)
	//parseint
	numInt64, _ := strconv.ParseInt(strA, 10, 64)
	fmt.Println(numInt64)
	strB := "1010"
	numInt8, _ := strconv.ParseInt(strB, 2, 8)
	fmt.Println(numInt8)
	//formatint = Berguna untuk konversi data numerik int64 ke string dengan basis numerik bisa ditentukan sendiri.
	var numInt = int64(24)
	strFormatInt := strconv.FormatInt(numInt, 8)
	fmt.Println(strFormatInt)

}
