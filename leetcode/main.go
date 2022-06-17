package main

import "fmt"

func main() {
	var nums1 []interface{}
	nums2 := []int{1, 3, 4}
	nums3 := append(nums1, nums2)
	fmt.Println(len(nums3))
}
