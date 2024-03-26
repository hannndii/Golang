package main

import "fmt"

func factailend(n int, hasil int) int{
	if n == 1{
		return hasil
	} else {
		return factailend(n-1, hasil*n)
	}
}

func main(){
	var  n, hasil int

	fmt.Scan(&n)

	hasil = 1
	factailend(n, hasil)

	fmt.Println(factailend(n, hasil))
}

// factailend(5-1, 5*1)
// factailend(4-1, 5*4)
// factailend(3-1, 20*3)
// factailend(2-1, 60*2)
// factailend(1, 120*1)