package blocking

import (
	"sync"
)

var (
	threadCount = 4
)

func ChangeThreadCount(tc int) {
	threadCount = tc
}

func CountMultiples(multipleOf int64, array []int64) int64 {
	var sum int64 = 0
	var page int = len(array) / threadCount
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := 0; i < len(array); {
		var high int
		if i+page < len(array)-1 {
			high = i + page
		} else {
			high = len(array) - 1
		}
		wg.Add(1)
		go func(low, high int) {
			defer wg.Done()
			for k := low; k < high; k++ {
				if array[k]%multipleOf == 0 {
					mu.Lock()
					sum++
					mu.Unlock()
				}
			}
		}(i, high)
		i += page
	}
	wg.Wait()
	return sum
}
