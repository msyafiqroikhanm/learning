package main

import (
	"errors"
	"fmt"
)

func SubTitle(SubTitle string) {
	fmt.Println()
	fmt.Println(SubTitle)
	for i := 0; i < len(SubTitle)+5; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	defer logging()
	fmt.Println("Run Application")
}

func endApp() {
	fmt.Println("Stopping Aplication")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("ERROR")
	}
}

func endAppWithRecover() {
	fmt.Println("Stopping Aplication")
	message := recover()
	fmt.Println("Terjadi Error :", message)
}

func runAppToRecover(error bool) {
	defer endAppWithRecover()
	if error {
		panic("ERROR")
	}
}

func pembagian(nilai uint, pembagi uint) (float64, error) {
	if pembagi == 0 {
		return 0.0, errors.New("maaf pembagi tidak boleh nol")
	} else {
		return float64(nilai / pembagi), nil
	}
}

func main() {
	SubTitle("Defer")
	/*
		Defer function adalah function yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function selesai di eksekusi.
		Defer function akan selalu dieksekusi walaupun terjadi error di function yang dieksekusi.
	*/
	runApplication()

	SubTitle("Panic")
	/*
		Panic function adalah function yang bisa kita gunakan untuk menghentikan program.
		Panic function biasanya dipanggil ketika terjadi error pada saat program kita berjalan.
		Saat panic function dipanggil, program akan terhenti,
		namun defer function tetap akan dieksekusi.
	*/
	// runApp(true)

	SubTitle("Recover")
	/*
		Recover adalah function yang bisa kita gunakan untuk menangkap data panic.
		Dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan.
	*/
	runAppToRecover(true)

	SubTitle("Error")
	/*
		Error merupakan sebuah tipe.
		Error memiliki 1 buah property berupa method Error(),
		method ini mengembalikan detail pesan error dalam string.
		Error termasuk tipe yang isinya bisa nil.
		cara menuliskan tipe data error cukup dengan menulis error (huruf kecil semua).
	*/
	// var input string
	// fmt.Print("Type some number: ")
	// fmt.Scanln(&input)

	// number, err := strconv.Atoi(input)
	// if err == nil {
	// 	fmt.Println(number, "is number")
	// } else {
	// 	fmt.Println(input, "is not number")
	// 	fmt.Println("Error :", err.Error())
	// }

	SubTitle("Custom Error")
	/*
		Selain memanfaatkan error hasil kembalian suatu fungsi internal yang tersedia,
		kita juga bisa membuat objek error sendiri dengan menggunakan fungsi errors.New()
		(harus import package errors terlebih dahulu).
	*/
	hasil, err := pembagian(8, 2)
	if err == nil {
		fmt.Println("Hasil:", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}

	hasil, err = pembagian(8, 0)
	if err == nil {
		fmt.Println("Hasil:", hasil)
	} else {
		fmt.Println("Error:", err.Error())
	}

	Package()

	Goroutine()
}
