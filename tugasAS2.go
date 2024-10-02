package main

import "fmt"

func main() {
	var (
		nama  string
		prodi = "S1-IF"
		kelas string
		nim   int
	)

	fmt.Println("masukan nama")
	fmt.Scan(&nama)

	fmt.Println("masukan kelas")
	fmt.Scan(&kelas)

	fmt.Println("masukan NIM")
	fmt.Scan(&nim)

	fmt.Println("perkenalkan saya adalah ", nama, "salah satu mahasiswa prodi", prodi, "dari kelas", kelas, "dengan NIM", nim)
}
