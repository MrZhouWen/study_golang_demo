package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 23, 5, 6, 8}
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	fmt.Println(arr)
}
