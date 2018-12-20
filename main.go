package main

import "fmt"

func main() {
	var a,b,c int
	var count = 0
	fmt.Scan(&a,&b,&c)
	for c > 0{
		c -= a
		if count % 7 == 6{
			c -= b
		}
		count++
	}
	fmt.Println(count)
}
