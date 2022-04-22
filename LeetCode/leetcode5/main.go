package main

import (
	"sort"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	sort.Ints(arr)
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))

}
