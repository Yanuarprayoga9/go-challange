// Soal 1: Menghitung Jumlah Bilangan Ganjil dan Genap
// Deskripsi:

// Buatlah sebuah program dalam bahasa Golang yang meminta pengguna untuk memasukkan sejumlah angka, kemudian menghitung dan menampilkan jumlah bilangan ganjil dan genap dari angka-angka tersebut.

// Spesifikasi:

// Program harus meminta pengguna untuk memasukkan jumlah angka yang akan dimasukkan.
// Pengguna kemudian memasukkan angka-angka tersebut satu per satu.
// Program menghitung dan menampilkan berapa jumlah bilangan ganjil dan berapa jumlah bilangan genap dari angka-angka yang telah dimasukkan.
// Contoh Input dan Output:

// yaml
// Salin kode
// Masukkan jumlah angka: 5
// Masukkan angka ke-1: 10
// Masukkan angka ke-2: 7
// Masukkan angka ke-3: 13
// Masukkan angka ke-4: 22
// Masukkan angka ke-5: 5

// Jumlah bilangan genap: 2
// Jumlah bilangan ganjil: 3

package main

import (
	"fmt"
)

func modifya(a *string) string {
	return *a
}
func main() {
	var longNum int
	fmt.Print("masukkan panjang bilangan = ")
	fmt.Scanf("%d",&longNum)

	listNum := make([]int,longNum);

	listNum[0] = 8;
	fmt.Print(listNum)


}
