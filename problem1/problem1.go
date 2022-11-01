package main

import "fmt"

func Swap(a, b *int) (int, int) {
	//your code here
	var c int = *b

	var x *int = a
	var y *int = b

	*y = *x
	*x = c
	fmt.Println(c)

	return *x, *y

}

func main() {
	a := 10
	b := 20
	fmt.Println(a, b)

	(Swap(&a, &b))
	fmt.Println(a, b)
}
