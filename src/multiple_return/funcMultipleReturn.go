package main

import "fmt"

func main() {
	area, perimeter := rectProps(10.8, 5.6)
	fmt.Printf("Area %f Perimeter %f", area, perimeter)

	//area, perimeter = rectProps1(10.8, 5.6)
	rectProps1(10.8, 5.6)
	fmt.Println("From rectProps1 Area %f Perimeter %f", area, perimeter)
}

func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

func rectProps1(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return //no explicit return value
}
