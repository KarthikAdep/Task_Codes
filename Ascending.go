package main

import "fmt"

func main() {

	sl := []int{6, 3, 8, 2, 7, 1}

	for i := 0; i < len(sl); i++ {
		for j := i + 1; j < len(sl); j++ {

			if sl[i] > sl[j] {
				sl[i], sl[j] = swap(sl[i], sl[j])
			}
		}
	}

	fmt.Println(sl)
}

func swap(a, b int) (int, int) {
	return b, a
}