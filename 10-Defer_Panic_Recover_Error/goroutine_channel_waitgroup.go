package main

import (
	"fmt"
	"sync"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func cetak(ch chan int) {
	fmt.Println("Ini dari go routine")
	ch <- 10 // channel mengirim value 10
}

func cetakWithClosing(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- 10
	}
	close(ch)
}

func printText(text string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
	}
	wg.Done()
}

func Goroutine() {
	fmt.Println("GOROUTINE")
	// Asynchronous GO
	fmt.Println("=================")

	SubTitle("Goroutine Implementation")
	go say("World in Goroutine")
	say("Hello Without Goroutine")

	fmt.Println("CHANNEL")
	fmt.Println("=============")
	/*
		Channel adalah penghubung antara goroutine satu dengan goroutine lainnya.
		Mekanisme channel yaitu serah terima,
		dimana ada yang mengirim dan ada yang menerima
	*/
	SubTitle("Channel Implementation")
	angka := make(chan int)
	go cetak(angka)
	nilai := <-angka // nilai menerima value dari channel
	fmt.Println("Angka CHannel", angka)
	fmt.Println("Nilai Channel integer", nilai)
	fmt.Println("In Function dari main")

	SubTitle("Closing Channel")
	/*
		Ketika kita bicara channel dua arah yaitu pengirim dan penerima maka sebenarnya ada hak lain yang dapat di lakukan.
		Pengirim mempunyai kemampuan menutup channel untuk memberi tahu penerima bahwa tidak ada data lagi.
		Penerima dapat memberikan tambahan kondisi apakah channel tersebut sudah di tutup ataupun belum.
	*/
	ch := make(chan int)
	go cetakWithClosing(ch)
	for {
		data, ok := <-ch
		if !ok {
			break
		}
		fmt.Printf("Data diterima %v \n", data)
	}

	SubTitle("Channel Buffer")
	/*
		channel secara default di lakukan secara unbufferd, artinya hanya data yang di kirim satu per satu melalui channel.
		Sedangkan dengan channel buffer kita dapat melakukan pengiriman data lebih daru satu.
		unbuffered bukan hanya pengirim saja yang di block,
		melainkan penerimanya juga di block ketika data sebelumnya belum selesai.
		dengan menggunakan channel buffer ini memungkinkan kita mengirim dan menerima banyak permintaan.
	*/
	channelBuffer := make(chan int, 3)
	// inserting value to channel
	channelBuffer <- 1
	channelBuffer <- 2
	channelBuffer <- 3

	// receiving and print
	fmt.Println(<-channelBuffer)
	fmt.Println(<-channelBuffer)
	fmt.Println(<-channelBuffer)

	SubTitle("WaitGroup")
	/*
		Jika di artikan bahasa Indonesia Wait dan Group bisa di artikan menunggu kelompok. Nah kelompok yang di maksud yaitu gorotuine.
		Sehingga WaitGroup adalah mekanisme digolang yang berfungsi untuk melakukan sinkronisasi antara goroutine
	*/
	var wg sync.WaitGroup

	wg.Add(1)
	go printText("Halo", &wg)

	wg.Add(1)
	go printText("Dunia", &wg)

	wg.Wait()
}
