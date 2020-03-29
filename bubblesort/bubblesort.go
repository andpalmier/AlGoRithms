// Please refer to https://www.calhoun.io/lets-learn-algorithms-implementing-bubble-sort

/*
Bubblesort implementation with improvements:
- don't compare numbers that are in the right position
- if slice is sorted, stop the execution
*/

package main

import "fmt"


func main() {
    var numbers []int = []int{6,5,3,4,1,2}
    fmt.Println("Not sorted: ", numbers)
    bubblesort(numbers)
    fmt.Println("Sorted: ", numbers)
}

func bubblesort(numbers []int){
    for i:=0;i<len(numbers);i++{

	// i keeps track of the steps done => used to avoid comparing numbers in the right position
	// if no sweep happened, return => slice is already sorted
	if !sweep(numbers,i){
	    return
	}
    }

}

func sweep(numbers []int, step int) bool {

    // assign n to to the lenght of the slice
    n := len(numbers)
    swap := false

    // index i and j start from 0,1
    i := 0
    j := 1

    // why n-step?
    // not needed to check last number after first iteration, because it's already at the bottom of the list
    for j < n-step {

	// assign numbers
	first := numbers[i]
	second := numbers[j]

	// change order of ith and jth element if not sorted
	if first > second {
	    swap = true
	    fmt.Println("switching: ",first,second)
	    numbers[i] = second
	    numbers[j] = first
	}

	// increase indexes
	j++
	i++
    }
    return swap
}
