// package main

// import "fmt"

// func main() {

// 	var hour, min, sec, conv int
// 	fmt.Scanln(&hour, &min, &sec)

// 	conv = jamConv(hour) + menitConv(min) + sec
// 	fmt.Println("Hasil konversi:", conv, "Detik")

// }

// func jamConv(jam int) int {
// 	return jam * 3600
// }

// func menitConv(menit int) int {
// 	return menit * 60
// }


package main

import "fmt"

func jamConv(detik int) {
	var convJam, convMenit, convDetik int
	convJam = detik / 3600
	convMenit = (detik % 3600) / 60
	convDetik = (detik % 3600) % 60
	
	fmt.Println(convJam, "jam", convMenit, "Menit", convDetik, "detik")
}


func main(){
	var detik int

	fmt.Scan(&detik)
	jamConv(detik)
}
