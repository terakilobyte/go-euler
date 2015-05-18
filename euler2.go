package main

import "fmt"

func main() {
	a := 1
	b := 2
	for b < 4000000 {
		tmp := a
		a = b
		b = tmp + b
		if b%2 == 0 {
			sum += b
		}
	}
	fmt.Println(sum)
}
