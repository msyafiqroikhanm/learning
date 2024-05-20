package main

import "fmt"

func main() {
	fmt.Println("SLICE")
	/*
		Slice adalah reference elemen array.
		Slice bisa dibuat, atau bisa juga dihasilkan dari manipulasi sebuah array ataupun slice lainnya.
		Karena merupakan data reference, menjadikan perubahan data di tiap elemen slice akan berdampak pada slice lain yang memiliki alamat memori yang sama.
	*/
	fmt.Println("=======================================")

	// slice initialize
	fmt.Println("Slice Initialize")
	fmt.Println("---------------------------")
	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits[0])

	var fruitsSlice = []string{"apple", "grape", "banana", "melon"}
	var fruitsArray1 = [2]string{"banana", "melon"}
	var fruitsArray2 = [...]string{"papaya", "grape"}
	fmt.Println("FruitsSlice", fruitsSlice)
	fmt.Println("fruitsArray1", fruitsArray1)
	fmt.Println("fruitsArray2", fruitsArray2)
	fmt.Println()

	// relation between slice with array and slice operation
	var newFruits = fruits[0:2]
	fmt.Println(newFruits)
	fmt.Println()

	// slice is a reference datatype
	fmt.Println("SLICE is Reference Data Type")
	fmt.Println("----------------------------")
	/*
		Slice merupakan tipe data reference atau referensi.
		Artinya jika ada slice baru yang terbentuk dari slice lama,
		maka data elemen slice yang baru akan memiliki alamat memori yang sama dengan elemen slice lama.
		Setiap perubahan yang terjadi di elemen slice baru,
		akan berdampak juga pada elemen slice lama yang memiliki referensi yang sama.
	*/
	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]

	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]

	fmt.Println("fruits", fruits)
	fmt.Println("aFruits", aFruits)
	fmt.Println("bFruits", bFruits)
	fmt.Println("aaFruits", aaFruits)
	fmt.Println("baFruits", baFruits)
	fmt.Println()

	fmt.Println(`Buah "Grape" diubah menjadi "pineapple"`)
	fmt.Println()
	baFruits[0] = "pineapple"

	fmt.Println("fruits", fruits)
	fmt.Println("aFruits", aFruits)
	fmt.Println("bFruits", bFruits)
	fmt.Println("aaFruits", aaFruits)
	fmt.Println("baFruits", baFruits)
	fmt.Println()

	// len() function
	fmt.Println("Function len()")
	fmt.Println("----------------------------")
	fmt.Println("Length of slice Fruits", len(fruits))
	fmt.Println()

	// cap() function
	fmt.Println("Function cap()")
	fmt.Println("----------------------")
	/*
		Fungsi cap() digunakan untuk menghitung lebar atau kapasitas maksimum slice.
		Nilai kembalian fungsi ini untuk slice yang baru dibuat pasti sama dengan len,
		tapi bisa berubah seiring operasi slice yang dilakukan

		Kita analogikan slicing 2 index menggunakan x dan y.

		fruits[x:y]
		Slicing yang dimulai dari indeks 0 hingga y,
		akan mengembalikan elemen-elemen mulai indeks 0 hingga sebelum indeks y,
		dengan lebar kapasitas adalah sama dengan slice aslinya.

		Sedangkan slicing yang dimulai dari indeks x,
		yang dimana nilai x adalah lebih dari 0,
		membuat elemen ke-x slice yang diambil menjadi elemen ke-0 slice baru.
		Hal inilah yang membuat kapasitas slice berubah.
	*/
	fmt.Println("len fruits", len(fruits))
	fmt.Println("cap fruits", cap(fruits))
	fmt.Println()
	fmt.Println("len aFruits", len(aFruits))
	fmt.Println("cap aFruits", cap(aFruits))
	fmt.Println()
	fmt.Println("len bFruits", len(bFruits))
	fmt.Println("cap bFruits", cap(bFruits))
	fmt.Println()

	// append() function
	fmt.Println("Function append()")
	fmt.Println("---------------------------")
	/*
		Fungsi append() digunakan untuk menambahkan elemen pada slice.
		Elemen baru tersebut diposisikan setelah indeks paling akhir.
		Nilai balik fungsi ini adalah slice yang sudah ditambahkan nilai barunya.
	*/
	fmt.Println(fruits)
	var cFruits = append(fruits, "papaya")
	fmt.Println(cFruits)
	fmt.Println(fruits) // tidak menambahkan referensi slice awalnya
	fmt.Println()

	// copy() function
	fmt.Println("Function copy()")
	fmt.Println("----------------------------")
	/*
		Fungsi copy() digunakan untuk men-copy elements slice pada src (parameter ke-2),
		ke dst (parameter pertama).
			copy(dst, src)
		Jumlah element yang di-copy dari src adalah sejumlah lebar slice dst (atau len(dst)).
		Jika jumlah slice pada src lebih kecil dari dst,
		maka akan ter-copy semua.
	*/

	destination := make([]string, 3) // dia bikin slice tipe data string, dengan panjang 3
	source := []string{"watermelon", "pinnaple", "apple", "orange"}
	n := copy(destination, source)

	fmt.Println("destination:", destination) // watermelon pinnaple apple
	fmt.Println("source:", source)           // watermelon pinnaple apple orange
	fmt.Println(n)                           // 3 = jumlah yang berhasil dicopy
	fmt.Println()

	dst := []string{"potato", "potato", "potato"}
	src := []string{"watermelon", "pinnaple"}
	n = copy(dst, src)

	fmt.Println("dst:", dst) // watermelon pinnaple potato
	fmt.Println("src:", src) // watermelon pinnaple
	fmt.Println(n)           // 2
	fmt.Println()

	// slice alocation with make() function
	fmt.Println("Function make()")
	fmt.Println("---------------------------")
	var vehicle = make([]string, 2)
	vehicle[0] = "bycicle"
	vehicle[1] = "plane"

	fmt.Println("vehicle:", vehicle)
}
