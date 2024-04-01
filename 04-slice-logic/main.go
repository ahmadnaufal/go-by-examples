package main

import "fmt"

func main() {
	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// arr=[0 1 2 3 4 5 6 7 8 9] len=10 cap=10
	fmt.Printf("arr=%v (%p) len=%d cap=%d\n", arr, &arr, len(arr), cap(arr))

	sliced := arr[2:4]
	// sliced=[2 3] len=2 cap=8
	// capacity will be reduced by number of elements we skip by the starting point of the slice
	fmt.Printf("sliced=%v (%p) len=%d cap=%d\n", sliced, &sliced, len(sliced), cap(sliced))
	sliced = append(sliced, 10, 11)
	// sliced=[2 3 10 11] len=4 cap=8
	// slice is updated & length is also updated. but capacity remains the same
	fmt.Printf("sliced=%v (%p) len=%d cap=%d\n", sliced, &sliced, len(sliced), cap(sliced))
	// arr=[0 1 2 3 10 11 6 7 8 9] len=10 cap=10
	// interesting behavior: when we append a new element to a slice, it will actually
	// replaces the original array source where the slice is taken from
	fmt.Printf("arr=%v (%p) len=%d cap=%d\n", arr, &arr, len(arr), cap(arr))

	sliced = append(sliced, 12, 13, 14, 15, 16)
	// sliced=[2 3 10 11 12 13 14 15 16] len=9 cap=16
	// slice capacity is doubled
	fmt.Printf("sliced=%v (%p) len=%d cap=%d\n", sliced, &sliced, len(sliced), cap(sliced))
	// arr=[0 1 2 3 10 11 6 7 8 9] len=10 cap=10
	// interesting behavior: if append is done to a slice which exceeded old capacity,
	// the newly appended elements in sliced won't replace original arr
	fmt.Printf("arr=%v (%p) len=%d cap=%d\n", arr, &arr, len(arr), cap(arr))

	slicednew := arr[:0]
	fmt.Printf("slicednew=%v (%p) len=%d cap=%d\n", slicednew, &slicednew, len(slicednew), cap(slicednew))

	// slices are copied by reference, while arrays are copied by value
	arr2 := make([]int, 10)
	// arr2 = []int{}
	fmt.Printf("arr2=%v (%p) len=%d cap=%d\n", arr2, arr2, len(arr2), cap(arr2))
	arr3 := arr2
	fmt.Printf("arr3=%v (%p) len=%d cap=%d\n", arr3, arr3, len(arr3), cap(arr3))

	arrfixed := [5]int{1, 2, 3, 4, 5}
	cparrfixed := arrfixed
	fmt.Printf("arrfixed=%v (%p) cparrfixed=%v (%p)\n", arrfixed, &arrfixed, cparrfixed, &cparrfixed)

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := slice1[2:4]
	fmt.Printf("slice1=%v (%p) len=%d cap=%d\n", slice1, slice1, len(slice1), cap(slice1))
	fmt.Printf("slice2=%v (%p) len=%d cap=%d\n", slice2, slice2, len(slice2), cap(slice2))
	slice2 = append(slice2, 6)
	fmt.Printf("slice2=%v (%p) len=%d cap=%d\n", slice2, slice2, len(slice2), cap(slice2))
}
