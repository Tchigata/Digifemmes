package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	var num int
	arr := []int{}
	lowRang := 0
	highRang := 0
	for i := 1; i <= 12500; i++ {
		fmt.Fscan(os.Stdin, &num)
		arr = append(arr, num)
		med := Median(arr)
		lowRang = int(med) - 45
		highRang = int(med) + 45
		fmt.Println(lowRang, highRang)
	}
}

func Median(data []int) float64 {
	l := len(data)
	a := (l + 1) / 2
	b := l / 2
	var res float64
	sort.Ints(data)
	if l == 0 {
		return 0
	} else if l%2 == 0 {

		res = float64((data[b-1] + data[b]) / 2)
	} else {
		res = float64(data[a-1])
	}
	return res
}
