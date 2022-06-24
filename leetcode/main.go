package main

import (
	"fmt"
	"math"
)

func main() {
	var nums1 = []int{1, 2, 3}
	nums2 := nums1[:2]

	fmt.Println(math.MaxInt8 >> 1)

	nums2 = append(nums2, 30)
	fmt.Printf("num1=%v, num2=%v", nums1, nums2)
	fmt.Printf("cap of num1 = %d; cap of num2 = %d \n\r", cap(nums1), cap(nums2))

	nums2 = append(nums2, 50)
	fmt.Printf("num1=%v, num2=%v", nums1, nums2)
	fmt.Printf("cap of num1 = %d; cap of num2 = %d \n\r", cap(nums1), cap(nums2))
}
