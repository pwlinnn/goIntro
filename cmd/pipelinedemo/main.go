package main

import (
	"fmt"
	"goIntro/pipeline"
)

func main() {
	p := pipeline.InMemSort(
		pipeline.ArraySource(3, 2, 6, 7, 4))
	for {
		if num, ok := <-p; ok {
			fmt.Println(num)
		} else {
			break
		}
	}
}
