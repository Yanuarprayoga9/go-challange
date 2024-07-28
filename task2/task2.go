// Soal 2: Mencari Bilangan Prima dalam Rentang Tertentu
// Deskripsi:

// Buatlah sebuah program dalam bahasa Golang yang meminta pengguna untuk memasukkan dua buah angka, yaitu batas bawah dan batas atas. Program kemudian mencari dan menampilkan semua bilangan prima dalam rentang tersebut.

// Spesifikasi:

// Program harus meminta pengguna untuk memasukkan dua angka, yaitu batas bawah dan batas atas.
// Program kemudian mencari semua bilangan prima dalam rentang tersebut.
// Program menampilkan bilangan prima yang ditemukan.
// Contoh Input dan Output:

// yaml
// Salin kode
// Masukkan batas bawah: 10
// Masukkan batas atas: 30

// Bilangan prima dalam rentang 10-30 adalah: 11, 13, 17

package main

import (
	"fmt"
)

func main() {
	var bbawah int
	var batas int
	value := []int{}
	fmt.Println("Masukkan batas bawah")
	fmt.Scan(&bbawah)
	fmt.Println("Masukkan batas atas")
	fmt.Scan(&batas)

	if bbawah > batas {
		fmt.Println("batas bawah tidak bisa lebih tinggi daridapa batas atas")
		return
	}

	exeptNum := [3]int{2, 3, 5}
	if bbawah < 2 {
		bbawah = 2
	}
	for i := bbawah - 1; i <= batas; i++ {

		if i%2 == 1 && i%3 != 0 && i > 1 || isInArray(i, exeptNum) {
			value = append(value, i)
		}
	}
	fmt.Println(value)

}

func isInArray(num int, numblist [3]int) bool {
	for _, n := range numblist {
		if n == num {
			return true
		}
	}
	return false
}
