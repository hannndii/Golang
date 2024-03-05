package main

import (
	"fmt"
)

func cariGo(kar byte){
	var i int
	count := 0
	str := string(kar)
	length := len(str)

	for i = 0; i <= length-1; i++{
		if str[i] == 'g' && str[i+1] == 'o'{
			count++
		}
	}
	fmt.Println(count)

}

func main(){
	var kar byte
	fmt.Scanf("%c", &kar)

	cariGo(kar)
}