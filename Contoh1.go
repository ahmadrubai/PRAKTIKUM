package main

import "fmt"

func main() {
	var mk string = "algoritma dan pemrograman"
	var kode string
	var sks int

	fmt.Print("tuliskan kode mk dan sks : ")
	fmt.Scan(&kode, &sks)
	fmt.Print("kredit mk ", kode, "-", mk, "1 adalah ", sks, "sks")

}
