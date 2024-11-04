package main

import (
	"fmt"
	"golangproject/utils"
	"os"
	"strconv"
)

func main() {
	// Mengambil nilai n dari argumen jika tersedia
	var n int
	if len(os.Args) > 1 {
		nArg, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Argument harus berupa angka, menggunakan nilai default 5.")
			n = 5
		} else {
			n = nArg
		}
	} else {
		n = 5 // Default n jika tidak ada argumen
	}

	fmt.Println("Hasil MagicSum:", utils.MagicSum(n))
	fmt.Println("Hasil MagicPow:", utils.MagicPow(n))
	fmt.Println("Apakah MagicOdd:", utils.MagicOdd(n))
	fmt.Println("MagicGrade:", utils.MagicGrade(n))
	fmt.Println("MagicName:", utils.MagicName(n))
	fmt.Println("MagicTria:", utils.MagicTria(n))

	// MagicChange dengan pointer
	fmt.Println("Nilai sebelum MagicChange:", n)
	utils.MagicChange(&n)
	fmt.Println("Nilai setelah MagicChange:", n)

	// Inisialisasi MagicNumber struct dan menggunakan method Multiply
	magicNum := utils.MagicNumber{Number: n}
	fmt.Println("Nilai MagicNumber sebelum Multiply:", magicNum.Number)
	magicNum.Multiply(2) // Kali dengan 3, bisa diubah
	fmt.Println("Nilai MagicNumber setelah Multiply:", magicNum.Number)
}
