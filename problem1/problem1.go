package main

import "fmt"

func Swap(a, b *int) (int, int) {
	//your code here
	var c int

	c = *a
	*a = *b
	*b = c

	return *a, *b

	// return *x, *y

	// *a,*b=*b,*a
	// return *a,*b

}

func main() {
	a := 10
	b := 20
	fmt.Println(a, b)

	(Swap(&a, &b))
	fmt.Println(a, b)
}
