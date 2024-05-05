package main

import "fmt"

const ROWMAX = 100
const COLMAX = 100

type arrMatriks [ROWMAX][COLMAX]int

func main() {
	var matriks arrMatriks
	var nRow, nCol int

	fmt.Println("Masukkan jumlah kolom dan baris untuk matriks:")
	fmt.Scan(&nRow, &nCol)

	baca(&matriks, nRow, nCol)
	cetak(matriks, nRow, nCol)
	minimum(matriks, nRow, nCol)

	fmt.Printf("%d matriks adalah matriks terkecil berada di baris %d kolom %d", minimum(matriks, nRow, nCol)  )


}

func baca(T *arrMatriks, nR int, nC int) {
	fmt.Println("Masukkan nilai integer untuk setiap elemen matriks:")

	for i := 0; i < nR; i++ {
		for j := 0; j < nC; j++ {
			fmt.Printf("Matriks[%d][%d]: ", i, j)
			fmt.Scan(&T[i][j])
		}
	}
}

func cetak(T arrMatriks, nR int, nC int) {
	fmt.Println("Matriks:")
	for i := 0; i < nR; i++ {
		for j := 0; j < nC; j++ {
			fmt.Print(T[i][j], " ")
		}
		fmt.Println()
	}
}

func minimum(T arrMatriks, nR, nC int) int{
	var min int
	
	min = T[0][1]

	for i := 0; i < nR; i++ {
		for j := 0; j < nC; j++ {
			if min > T[i][j] {
				min = T[i][j]
			}
		}
	}
	return min
}
