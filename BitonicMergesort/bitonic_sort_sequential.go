package main

import "fmt"

func main() {
	target := []float64{10, 30, 11, 20, 4, 330, 21, 110}
	fmt.Println(bitonic_sort(true, target))
}

func bitonic_sort(up bool, x []float64) []float64 {
	if len(x) <= 1 {
		return x
	} else {
		first := bitonic_sort(true, x[:int(len(x)/2)])
		second := bitonic_sort(false, x[int(len(x)/2):])
		return bitonic_merge(up, append(first, second...))
	}

}
func bitonic_merge(up bool, x []float64) []float64 {
	if len(x) == 1 {
		return x
	} else {
		bitonic_compare(up, x)
		first := bitonic_merge(up, x[:int(len(x)/2)])
		second := bitonic_merge(up, x[int(len(x)/2):])
		return append(first, second...)

	}
}
func bitonic_compare(up bool, x []float64) {
	dist := int(len(x) / 2)
	for i := 0; i < dist; i++ {
		if x[i] > x[i+dist] == up {
			tmp := x[i]
			x[i] = x[i+dist]
			x[i+dist] = tmp
		}
	}
}
