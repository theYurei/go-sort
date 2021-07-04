package main

import (
	"fmt"
	"os"
)

func main() {

	arr := []int{32,34,11,7,47,84,88,54,777,666,66,6,137,731,1112,3222,1,1227,7331,111,9,11111111,2222222,3333333,4444444444,11111111111,32132131,1727}
	var x int = 0
	var z int = x + 1
	fmt.Println("antes: ", arr)

for true {
	if (arr[x] < arr[z]) {
		x++
		z = x + 1
		
	} else {
			var xarr int = arr[x]
			var yarr int = arr[z]
			arr[x] = yarr
			arr[z] = xarr
			x = 0
			z = x + 1
		}
		if (z == 16) {

			fmt.Println("depois: ",arr)
			os.Exit(1)
		}
	}
}


