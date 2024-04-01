package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("arr=%v, %p\n", arr, &arr)
	fooArr(arr)
	fmt.Printf("after fooArr: arr=%v, %p\n", arr, &arr)

	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("slice=%v, %p\n", slice, slice)
	fooSlice(slice)
	fmt.Printf("after fooSlice: slice=%v, %p\n", slice, slice)
}

func fooArr(arr [5]int) {
	fmt.Printf("inside fooArr: arr=%v, %p\n", arr, &arr)
	arr[1] = 10
	fmt.Printf("inside fooArr after updating: arr=%v, %p\n", arr, &arr)
}

func fooSlice(slice []int) {
	fmt.Printf("inside fooSlice: slice=%v, %p\n", slice, slice)
	slice[1] = 10
	fmt.Printf("inside fooSlice after updating: slice=%v, %p\n", slice, slice)
}
