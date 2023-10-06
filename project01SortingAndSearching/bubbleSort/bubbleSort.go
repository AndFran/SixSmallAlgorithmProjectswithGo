package main

import (
	"fmt"
	"math/rand"
)

// makeRandomSlice make a slice containing pseudorandom numbers in [0, max).
func makeRandomSlice(numItems, max int) []int {
	randomNumbers := make([]int, 0, numItems)
	for i := 0; i < numItems; i++ {
		randomNumbers = append(randomNumbers, rand.Intn(max))
	}
	return randomNumbers
}

// Print at most numItems items.
func printSlice(slice []int, numItems int) {
	maxToShow := min(len(slice), numItems)
	fmt.Println(slice[:maxToShow])
	//for i := 0; i < maxToShow; i++ {
	//}
}

// Verify that the slice is sorted.
func checkSorted(slice []int) {
	previous := slice[0]
	for i := 1; i < len(slice); i++ {
		if previous > slice[i] {
			fmt.Println("The slice is NOT sorted!")
			return
		}
	}
	fmt.Println("The slice is sorted")
}

// Use bubble sort to sort the slice.
func bubbleSort(slice []int) {
	for {
		swap := false
		for x := 0; x < len(slice)-1; x++ {
			if slice[x] > slice[x+1] {
				slice[x], slice[x+1] = slice[x+1], slice[x]
				swap = true
			}
		}
		if !swap {
			break
		}
	}
}

func main() {
	// Get the number of items and maximum item value.
	var numItems, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	// Make and display an unsorted slice.
	slice := makeRandomSlice(numItems, max)
	printSlice(slice, 40)
	fmt.Println()

	// Sort and display the result.
	bubbleSort(slice)
	printSlice(slice, 40)

	// Verify that it's sorted.
	checkSorted(slice)
}
