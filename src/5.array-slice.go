package main

import "fmt"

func main() {
	fmt.Println("Array...")
	// Fixed length array
	// var points [3] int
	points := [3]int{1, 2, 3}

	fmt.Println(points)

	// Dynamic length array
	points2 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(points2)

	// Lenghth of array
	fmt.Println("length of points2:", (len(points2)))

	// Copy array to new memory block

	copiedpoints := points

	copiedpoints[0] = 100 // points[0] doesn't change

	fmt.Println("copiedpoints ==> ", copiedpoints)

	// Clone as ref type

	cloneAsRefType := &points

	cloneAsRefType[0] = 101 // points[0] change to 101 also

	fmt.Println(points, cloneAsRefType)

	// Slice
	arraySlice := []int{1, 2, 3, 4, 5, 6}
	arraySlice2 := arraySlice
	arraySlice2[0] = 100 // arraySlice[0] change to 100 also

	fmt.Println("arraySlice ==>>", arraySlice)
	fmt.Println("arraySlice2 ==>>", arraySlice2)

	b := arraySlice[:]
	c := arraySlice[3:]
	d := arraySlice[:5]
	e := arraySlice[3:5]

	fmt.Println("=======Slice======")

	fmt.Printf("a %v %v %v\n", arraySlice, len(arraySlice), cap(arraySlice))
	fmt.Printf("b %v %v %v\n", b, len(b), cap(b))
	fmt.Printf("c %v %v %v\n", c, len(c), cap(c))
	fmt.Printf("d %v %v %v\n", d, len(d), cap(d))
	fmt.Printf("e %v %v %v\n", e, len(e), cap(e))

	// Slice with make
	sliceMake := make([]int, 0)
	fmt.Printf("sliceMake ===> %v %v %v\n", sliceMake, len(sliceMake), cap(sliceMake))

	sliceMake = append(sliceMake, 1)
	sliceMake = append(sliceMake, 1, 2, 3, 4, 5, 6)
	sliceMake = append(sliceMake, 1, 2, 3, 4, 5, 6)
	fmt.Printf("sliceMake ===> %v %v %v\n", sliceMake, len(sliceMake), cap(sliceMake))
}
