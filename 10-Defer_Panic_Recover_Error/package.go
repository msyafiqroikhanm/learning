package main

import (
	"flag"
	"fmt"
	"math"
	"reflect"
	"sort"
	"time"
)

func Package() {
	fmt.Println("\n==========================")
	fmt.Println("PACKAGE")
	fmt.Println("==========================")
	SubTitle("Package Flag")
	var name = flag.String("name", "anonymous", "Type your name")
	var age = flag.Int64("age", 22, "Type your age")

	flag.Parse()
	fmt.Printf("name\t: %s\n", *name)
	fmt.Printf("age\t: %d\n", *age)

	SubTitle("Package Math")
	fmt.Println("Round :", math.Round(1.7))
	fmt.Println("Round :", math.Round(1.3))
	fmt.Println("Floor :", math.Floor(1.7))
	fmt.Println("Ceil :", math.Ceil(1.3))
	fmt.Println("Max :", math.Max(10, 20))
	fmt.Println("Min :", math.Min(10, 20))

	SubTitle("Package Time")
	now := time.Now()
	fmt.Println(now)

	utc := time.Date(2020, 10, 10, 10, 10, 10, 10, time.UTC)
	fmt.Println(utc)

	layout := "2006-01-02"
	parse, _ := time.Parse(layout, "1990-03-20")
	fmt.Println(parse)
	fmt.Println()

	fmt.Println("start")
	time.Sleep(time.Second * 2)
	fmt.Println("after 4 seconds")

	SubTitle("Package Sort")
	s := []string{"Xiaomi", "asus", "iPhone", "Samsung", "Realme", "motorola"}
	fmt.Println("Before Sort:", s)
	sort.Strings(s)
	fmt.Println("After Sort:", s)

	SubTitle("Package Reflect")
	/*
		Dalam bahasa pemrograman, biasanya ada fitur Reflection,
		dimana kita bisa melihat struktur kode kita pada saat aplikasi sedang berjalan.
	*/
	var number = 23
	var reflectValue = reflect.ValueOf(number)

	fmt.Println("tipe  variabel :", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variabel :", reflectValue.Int())
	}
	fmt.Println("Reflect Value", reflectValue)
}
