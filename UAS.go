package main

// Import package yang digunakan dalam program
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Menentukan kapasitas maksimum data film
const MAX = 100

// Struct untuk menyimpan data film
type Film struct {
	Judul, Genre, Status string
	Rating               float64
}

// Variabel global program
var (
	data    [MAX]Film
	jumlah  int
	scanner = bufio.NewScanner(os.Stdin)
)

// Fungsi untuk menerima input teks dari pengguna
func input(prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

// Fungsi untuk mengisi data film
func inputFilm() Film {
	var f Film

	f.Judul = input("Judul: ")
	f.Genre = input("Genre: ")

	// Input rating film
	fmt.Print("Rating: ")
	fmt.Scan(&f.Rating)

	// Membersihkan buffer input
	scanner.Scan()

	f.Status = input("Status: ")
	return f
}

// Fungsi pencarian film menggunakan Sequential Search
func cari(judul string) int {
	for i := 0; i < jumlah; i++ {
		if strings.EqualFold(data[i].Judul, judul) {
			return i
		}
	}
	return -1
}

// Fungsi untuk menampilkan satu data film
func tampil(f Film) {
	fmt.Printf("  Judul: %s | Genre: %s | Rating: %.1f | Status: %s\n", f.Judul, f.Genre, f.Rating, f.Status)
}

// Fungsi pengurutan judul secara ascending menggunakan Selection Sort
func sortAsc() {
	for i := 0; i < jumlah-1; i++ {
		min := i

		for j := i + 1; j < jumlah; j++ {
			if strings.ToLower(data[j].Judul) < strings.ToLower(data[min].Judul) {
				min = j
			}
		}

		data[i], data[min] = data[min], data[i]
	}
}

// Fungsi pengurutan rating secara descending menggunakan Insertion Sort
func sortDesc() {
	for i := 1; i < jumlah; i++{
		key := data[i]
		j := i - 1

		for j >= 0 && data[j].Rating < key.Rating{
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}

// Fungsi pencarian menggunakan Binary Search
func binarySearch(judul string) int {
	lo, hi := 0, jumlah-1

	for lo <= hi {
		mid := (lo + hi) / 2
		cmp := strings.ToLower(data[mid].Judul)

		if cmp == strings.ToLower(judul) {
			return mid
		} else if cmp < strings.ToLower(judul) {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return -1
}

// Fungsi utama program
func main() {
	for {

		// Menampilkan menu utama
		fmt.Println("\n1. Tambah Film")
		fmt.Println("2. Tampilkan Film")
		fmt.Println("3. Edit Film")
		fmt.Println("4. Hapus Film")
		fmt.Println("5. Cari Film")
		fmt.Println("6. Urutkan Film")
		fmt.Println("7. Keluar")

		var pilih int

		// Meminta pilihan menu dari pengguna
		fmt.Print("Pilih: ")
		fmt.Scan(&pilih)
		scanner.Scan()

		switch pilih {

		// Menambahkan data film baru
		case 1:
			if jumlah < MAX {
				data[jumlah] = inputFilm()
				jumlah++
				fmt.Println("Berhasil ditambahkan!")
			}

		// Menampilkan seluruh data film
		case 2:
			for i := 0; i < jumlah; i++ {
				fmt.Printf("%d.", i+1)
				tampil(data[i])
			}

		// Mengubah data film yang sudah ada
		case 3:
			if idx := cari(input("Judul dicari: ")); idx != -1 {
				data[idx] = inputFilm()
				fmt.Println("Berhasil diubah!")
			} else {
				fmt.Println("Tidak ditemukan!")
			}

		// Menghapus data film
		case 4:
			if idx := cari(input("Judul dihapus: ")); idx != -1 {
				for i := idx; i < jumlah-1; i++ {
					data[i] = data[i+1]
				}
				jumlah--
				fmt.Println("Berhasil dihapus!")
			} else {
				fmt.Println("Tidak ditemukan!")
			}

		// Menu pencarian film
		case 5:
			fmt.Println("1. Sequential Search")
			fmt.Println("2. Binary Search")
			fmt.Print("Pilih: ")

			var p int
			fmt.Scan(&p)
			scanner.Scan()

			judul := input("Judul: ")
			var idx int

			// Menggunakan Binary Search
			if p == 2 {
				sortAsc()
				idx = binarySearch(judul)
			} else {
				// Menggunakan Sequential Search
				idx = cari(judul)
			}

			if idx != -1 {
				tampil(data[idx])
			} else {
				fmt.Println("Tidak ditemukan!")
			}

		// Menu pengurutan data
		case 6:
			fmt.Print("1.Judul ASC 2.Rating DESC — Pilih: ")

			var p int
			fmt.Scan(&p)
			scanner.Scan()

			// Mengurutkan berdasarkan judul
			if p == 1 {
				sortAsc()
			} else {
				// Mengurutkan berdasarkan rating
				sortDesc()
			}

			fmt.Println("Data berhasil diurutkan!")

		// Keluar dari program
		case 7:
			return
		}
	}
}