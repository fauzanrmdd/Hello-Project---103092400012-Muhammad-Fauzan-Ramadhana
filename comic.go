package main

import "fmt"

const NMAX int = 1024

type Komik struct {
	id      string
	judul   string
	penulis string
	stok    int
}

type tabKomik [NMAX]Komik

func main() {
	var data tabKomik
	var nData, tempInt int
	var idPertama, idDua, tempString string

	// Komentar
	fmt.Println("Ini untuk salah satu tugas besar nya pak")
	readData(&data, &nData)
	fmt.Println("Masukkan ID Pertama: ")
	fmt.Scan(&idPertama)
	fmt.Println("Masukkan ID Kedua: ")
	fmt.Scan(&idDua)

	printData(data, nData)

	tempString = judulKomikSeq(&data, idPertama, nData)
	tempInt = totalStok(&data, nData)
	fmt.Println("Judul komik: ", tempString)
	fmt.Println("Total stok: ", tempInt)

	urutIDMenaik(&data, nData)
	printData(data, nData)

	tempString = searchPenulisID(&data, idDua, nData)
	fmt.Println("Nama penulis: ", tempString)

	urutStokMenurun(&data, nData)
	printData(data, nData)

}

func readData(A *tabKomik, n *int) {
	var id, judul, penulis string
	var i, stok int

	i = 0
	fmt.Println("Masukkan id, judul, penulis, stok: ")
	fmt.Scan(&id, &judul, &penulis, &stok)
	for id != "none" && judul != "none" && penulis != "none" && stok != 0 {
		A[i].id = id
		A[i].judul = judul
		A[i].penulis = penulis
		A[i].stok = stok
		i++
		fmt.Scan(&id, &judul, &penulis, &stok)
	}
	*n = i
}

func printData(A tabKomik, n int) {
	fmt.Println("=======================")
	fmt.Println("		DATA KOMIK		")
	fmt.Println("=======================")
	fmt.Printf("%5s %13s %10s %5s\n", "ID", "Judul", "Penulis", "Stok")
	for i := 0; i < n; i++ {
		fmt.Printf("%5s %13s %10s %5d\n", A[i].id, A[i].judul, A[i].penulis, A[i].stok)
	}
	fmt.Println("=======================")
}

func judulKomikSeq(A *tabKomik, id string, n int) string {
	for i := 0; i < n; i++ {
		if A[i].id == id {
			return A[i].judul
		}
	}
	return "none"
}

func totalStok(A *tabKomik, n int) int {
	var total int
	total = 0
	for i := 0; i < n; i++ {
		total += A[i].stok
	}
	return total
}

func urutIDMenaik(A *tabKomik, n int) {
	var iMin int
	var temp Komik
	for pass := 1; pass < n; pass++ {
		iMin = pass - 1
		for i := pass; i < n; i++ {
			if A[i].id < A[iMin].id {
				iMin = i
			}
		}
		temp = A[pass-1]
		A[pass-1] = A[iMin]
		A[iMin] = temp
	}
}

func urutIDMenurun(A *tabKomik, n int) {
	var iMax int
	var temp Komik
	for pass := 1; pass < n; pass++ {
		iMax = pass - 1
		for i := pass; i < n; i++ {
			if A[i].id > A[iMax].id {
				iMax = i
			}
		}
		temp = A[pass-1]
		A[pass-1] = A[iMax]
		A[iMax] = temp
	}
}

func searchPenulisID(A *tabKomik, id string, n int) string {
	var right, left, mid int
	urutIDMenurun(A, n)
	left = 0
	right = n - 1
	for left <= right {
		mid = (left + right) / 2
		if A[mid].id == id {
			return A[mid].penulis
		} else if A[mid].id > id {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return "none"
}

func urutStokMenurun(A *tabKomik, n int) {
	var temp Komik
	for i := 1; i < n; i++ {
		temp = A[i]
		j := i - 1
		for j >= 0 && A[j].stok < temp.stok {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = temp
	}
}
