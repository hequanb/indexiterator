package indexiterator

import (
	"errors"
	"fmt"
)

var ErrInvalidParam = errors.New("index iterator param error")

// start = (i-1)*length
type Iterator struct {
	i         int // 已遍历次数
	count     int
	batchSize int
}

func Build(count, batchSize int) (*Iterator, error) {
	Iterator := new(Iterator)
	
	if count < 0 {
		return nil, fmt.Errorf("count must larger than 0, %w", ErrInvalidParam)
	}
	
	if batchSize <= 0 {
		return nil, fmt.Errorf("batchSize must larger or equal than 0, %w", ErrInvalidParam)
	}
	
	Iterator.count = count
	Iterator.batchSize = batchSize
	return Iterator, nil
}

func (iter *Iterator) Iter(start, length *int) bool {
	// if *start+*length >= iter.count {
	// 	return false
	// }
	*length = iter.batchSize
	// 还没结束的话，开始计算
	if iter.i > 0 {
		// 计算本次的漂移量
		*start = *start + *length
	}
	iter.i++
	
	if *start+*length > iter.count {
		// 修正剩余的长度
		*length = iter.count - *start
	}
	
	return *start < iter.count
}
