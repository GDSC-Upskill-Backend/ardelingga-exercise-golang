package utils

// import "fmt"

func Sorted(data []int) []int {
	// fmt.Println("ORIGINAL DATA")
	// fmt.Println(data)

	// TODO 1: Loop through data and sort it
	// Use method buble sort
	var n = len(data)
	for i := 0; i < n; i++ {
		for j := n - 1; j > i; j-- {
			if data[j] < data[j-1] {
				currData := data[j]
				data[j] = data[j-1]
				data[j-1] = currData
			}
		}
	}
	// fmt.Println("RESULT")
	// fmt.Println(data)

	// TODO 2: Return the sorted data
	return data
}
