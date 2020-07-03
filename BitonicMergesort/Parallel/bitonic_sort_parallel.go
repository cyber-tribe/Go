package main

import (
	"fmt"
)

func main() {
	target := []float64{10, 30, 11, 20, 4, 330, 21, 110}
	fmt.Println(bitonic_sort(true, target))
}

func bitonic_sort(up bool, x []float64) []float64 {
	if len(x) <= 1 {
		return x
	} else {
		//チャネル
		firstChan := make(chan []float64)
		secondChan := make(chan []float64)
		go func() {
			// fmt.Println(x)
			first := bitonic_sort(true, x[:int(len(x)/2)])
			firstChan <- first
		}()
		go func() {
			// fmt.Println(x)
			second := bitonic_sort(false, x[int(len(x)/2):])
			secondChan <- second
		}()
		return bitonic_merge(up, append(<-firstChan, <-secondChan...))
	}

}
func bitonic_merge(up bool, x []float64) []float64 {
	if len(x) == 1 {
		return x
	} else {
		//チャネル
		firstChan := make(chan []float64)
		secondChan := make(chan []float64)
		bitonic_compare(up, x)
		go func() {
			first := bitonic_merge(up, x[:int(len(x)/2)])
			firstChan <- first
		}()
		go func() {
			second := bitonic_merge(up, x[int(len(x)/2):])
			secondChan <- second
		}()
		return append(<-firstChan, <-secondChan...)

	}
}
func bitonic_compare(up bool, x []float64) {
	dist := len(x) / 2
	for i := 0; i < dist; i++ {
		if x[i] > x[i+dist] == up {
			tmp := x[i]
			x[i] = x[i+dist]
			x[i+dist] = tmp
		}
	}
}
