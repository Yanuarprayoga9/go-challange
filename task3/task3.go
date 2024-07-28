// Soal 3: Menghitung Total Harga Barang dengan Diskon
// Deskripsi:

// Buatlah sebuah program dalam bahasa Golang yang meminta pengguna untuk memasukkan harga barang dan jumlah barang yang dibeli. Jika total harga melebihi 100.000, pengguna mendapatkan diskon 10%. Program kemudian menghitung dan menampilkan total harga setelah diskon (jika ada).

// Spesifikasi:

// Program harus meminta pengguna untuk memasukkan harga barang dan jumlah barang.
// Program menghitung total harga.
// Jika total harga lebih dari 100.000, program memberikan diskon 10%.
// Program menampilkan total harga setelah diskon (jika ada).
// Contoh Input dan Output:

// yaml
// Salin kode
// Masukkan harga barang: 25000
// Masukkan jumlah barang: 5

// Total harga sebelum diskon: 125000
// Total harga setelah diskon: 112500


package main 

import "fmt"

func main (){
	var price float64 = 0;
	var x int = 0;
	fmt.Println("masukkan harga barang")
	fmt.Scan(&price)
	fmt.Println("masukkan jumlah barang")
	fmt.Scan(&x)
	
	total := price * float64(x);
	afterDiscount := total; 
	if total > 100000 {
		afterDiscount  = afterDiscount * 0.9
	}

	fmt.Printf("Total sebelum diskon: %.2f\n", total)
	fmt.Printf("Total setelah diskon: %.2f\n", afterDiscount)

	
}