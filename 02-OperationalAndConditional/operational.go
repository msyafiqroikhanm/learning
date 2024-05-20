package main

import "fmt"

func main() {
	//* CONDITIONAL WITH IF ELSE
	fmt.Println("CONDITIONAL WITH IF ELSE")
	fmt.Println("---------------------------------")
	// basic
	var mood = "happy"
	if mood == "happy" {
		fmt.Println("Hari ini aku bahagia")
	}
	fmt.Println()

	// if else
	var marketStatus = "open"
	if marketStatus == "open" {
		fmt.Println("Saya akan membeli telur dan buah")
	} else {
		fmt.Println("Minimarketnya tutup")
	}
	fmt.Println()

	// if else if and else
	var minimarketStatus = "close"
	minuteRemainingToOpen := 5
	if minimarketStatus == "open" {
		fmt.Println("Saya akan membeli telur dan buah")
	} else if minuteRemainingToOpen <= 5 {
		fmt.Println("Minimarket buka sebentar lagi, saya akan menunggunya")
	} else {
		fmt.Println("Minimarket tutup, saya pulang lagi kalau begitu")
	}
	fmt.Println()

	// nested conditional
	minimarketStatus = "open"
	var telur = "soldout"
	var buah = "soldout"
	if minimarketStatus == "open" {
		fmt.Println("Saya akan membeli telur dan buah")
		if telur == "soldout" || buah == "soldout" {
			fmt.Println("Belanjaan saya tidak lengkap")
		} else if telur == "soldout" {
			fmt.Println("Telur sudah habis")
		} else if buah == "soldout" {
			fmt.Println("Buah sudah habis")
		}
	} else {
		fmt.Println("Minimarket tutup, saya pulang lagii")
	}
	fmt.Println("===================================")
	fmt.Println()

	//* VARIABLE TEMPORARY PADA IF ELSE
	fmt.Println("VARIABLE TEMPORARY PADA IF ELSE")
	fmt.Println("---------------------------------")
	if minimarketStatus, minuteRemainingToOpen := "close", 5; minimarketStatus == "open" {
		fmt.Println("saya akan membeli telur dan buah")
	} else if minuteRemainingToOpen <= 5 {
		fmt.Println("minimarket buka sebentar lagi, saya tungguin")
	} else {
		fmt.Println("minimarket tutup, saya pulang lagi")
	}
	fmt.Println("===================================")
	fmt.Println()

	//* SWITCH CASE
	fmt.Println("SWITCH CASE")
	fmt.Println("---------------------------------")
	buttonPushed := 1
	switch buttonPushed {
	case 1:
		fmt.Println("Matikan TV")
	case 2:
		fmt.Println("Turunkan Volume TV")
	case 3:
		fmt.Println("Tingkatkan Volume Tv")
	case 4:
		fmt.Println("Matikan suara TV")
	default:
		fmt.Println("Tidak terjadi apa apa")
	}
	fmt.Println()

	switch buttonPushed {
	case 1:
		fmt.Println("Matikan TV")
	case 2, 3, 4:
		fmt.Println("Turunkan Volume TV")
	default:
		fmt.Println("Tidak terjadi apa apa")
	}
	fmt.Println()

	tvStatus := "on"
	switch {
	case buttonPushed == 1:
		fmt.Println("matikan TV!")
	case buttonPushed == 2 && tvStatus == "on":
		fmt.Println("turunkan volume TV!")
	case buttonPushed == 4:
		fmt.Println("matikan suara TV!")
	default:
		fmt.Println("Tidak terjadi apa-apa")
	}

	var point = 6

	switch {
	case point == 8:
		fmt.Println("perfect")
	case (point < 8) && (point > 3):
		fmt.Println("awesome")
		fallthrough
	case point < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}
}
