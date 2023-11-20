package counting

import (
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GenerateNumbers - random number generation
func GenerateNumbers(max int) []int {
	rand.Seed(time.Now().UnixNano())
	numbers := make([]int, max)
	for i := 0; i < max; i++ {
		numbers[i] = rand.Intn(10)
	}
	return numbers
}

// Add - sequential code to add numbers
func Add(numbers []int) int64 {
	var sum int64
	for _, n := range numbers {
		sum += int64(n)
	}
	return sum
}

//TODO: complete the concurrent version of add function.

// AddConcurrent - concurrent code to add numbers
func AddConcurrent(numbers []int) int64 {
	// Utilize all cores on machine
	numOfCores := runtime.NumCPU()
	runtime.GOMAXPROCS(numOfCores)

	var (
		wg sync.WaitGroup

		sum             int64
		sizeForEachCore = len(numbers) / numOfCores
		addNumber       = func(numbersToAdd []int) {
			defer wg.Done()

			var localSum int64

			for _, n := range numbersToAdd {
				localSum += int64(n)
			}

			atomic.AddInt64(&sum, localSum)
		}
	)

	for i := 0; i < numOfCores; i++ {
		start := i * sizeForEachCore

		wg.Add(1)
		go addNumber(numbers[start : start+sizeForEachCore])
	}

	wg.Wait()

	return sum
}
