package indexiterator

import (
	"errors"
	"fmt"
	"testing"
)

func TestOnIterator(t *testing.T) {
	var i = []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13,
	}
	iter, err := Build(len(i), 3)
	if err != nil {
		fmt.Println(errors.Is(err, ErrInvalidParam))
		panic("param error")
	}
	
	var (
		start  int
		length int
	)
	for iter.Iter(&start, &length) {
		fmt.Println("start:", start, ", length:", length)
		fmt.Println(i[start : start+length])
	}
}
