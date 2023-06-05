package utils

// TODO 1: Import sync package
import (
	"fmt"
	"time"
)

type StatisticsResult struct {
	Mean   float64
	Median float64
	Max    int
}

// Statistics calculates thM mean, median and max of a slice of integers.
func Statistics(data []int) StatisticsResult {
	// TODO 2: Create Variables wg with type sync.WaitGroup
	// var wg sync.WaitGroup = sync.WaitGroup{}

	// TODO 3: Create Variables mean, median with type chan float64, and max with type chan int
	mean := make(chan float64)
	median := make(chan float64)
	max := make(chan int)

	// TODO 4: Create a goroutine to calculate the mean of data
	go func() {
		var n = float64(len(data))

		// Find sum
		sum := float64(0)
		for _, value := range data {
			sum += float64(value)
		}
		mean <- (sum / n)
	}()
	time.Sleep(time.Second)
	fmt.Println("MEAN : %", mean)

	// TODO 5: Create a goroutine to calculate the median of data
	go func(data []int) {
		var n = len(data)
		var dataSorted []int = Sorted(data)

		// fmt.Println("PRINT DATA SORTED")
		// fmt.Println(dataSorted)

		// TODO 5.1: Get Median if data is odd
		if n%2 == 0 {
			median1 := dataSorted[n/2]
			median2 := dataSorted[(n/2)-1]

			median <- (float64(median1) + float64(median2)) / 2

			// TODO 5.2: Get Median if data is even
		} else {
			median <- float64(dataSorted[n/2])
		}

	}(data)
	time.Sleep(time.Second)
	fmt.Println("MEDIAN : %", median)

	// TODO 6: Create a goroutine to calculate the max of data
	go func(data []int) {
		// Loop for find max value
		maxValue := 0
		for i := 0; i < len(data); i++ {
			currValue := data[i]
			if currValue > maxValue {
				maxValue = currValue
			}
		}
		// fmt.Println("PRINT MAX VALUE")
		// fmt.Println(maxValue)

		max <- maxValue

	}(data)
	time.Sleep(time.Second)
	fmt.Println("MAX : %", max)

	// TODO 7: Create a goroutine to close all channels
	go func() {
		defer close(mean)
		defer close(median)
		defer close(max)
	}()

	// TODO 8: Create a variable result with type StatisticsResult
	var result StatisticsResult

	// TODO 9: Assign the value of mean, median, max to result
	result = StatisticsResult{
		Mean:   <-mean,
		Median: <-median,
		Max:    <-max,
	}

	fmt.Println("RESULT")
	fmt.Println(result)

	// TODO 10: Return result
	return result
}
