package main

import "fmt"

func fibonacci() func(int) int {
	old1 := 0;
	old2 := 1;
	return func(i int) int {
		if i == 0 {
			return 0
		}
		new1 := old2 + old1
		fmt.Printf("new1(%d) = old2(%d) + old1(%d)\n",
			new1, old2, old1)
		old1 = old2
		old2 = new1
		return new1
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
