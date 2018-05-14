package main

import "fmt"
import "parent"

//func add(a , b int)

//func add(a , b int) (c int)

// declare
//var c int

//func add(a , b int) (int){
//c:=a+b
//retun c
//}

func add(a int, b int) int {
	return a + b
}

func minus(a int, b int) int {
	return a - b
}

func fn(n int) (r int) {
	if n <= 1 {
		r = n
	} else {
		r = fn(n-1) + fn(n-2)
	}
	return r
}

//func echo() { fun }

func main() {
	fmt.Println("Hello World")
	// begin with package parent
	fmt.Println(parent.Var)
	fmt.Println(add(1, 2))
	fmt.Println(minus(4, 2))

	//while
	//sum :=1
	//for sum< 1000;{sum+=sum}
	//fmt.Println(sum)

	sum := 1
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	fmt.Println(fn(5))

	fmt.Println("fibonucci result is ", fn(13))

}

//inline
//println(func(a, b int) int { return a+b} (2,3))
