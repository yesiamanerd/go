package main

import "fmt"

func main() {
	const value = 10
	var i int = int(value)
	var f float64 = float64(i)
	fmt.Println(i, f)
}
