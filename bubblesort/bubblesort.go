// inspired from https://www.calhoun.io/lets-learn-algorithms-implementing-bubble-sort

package main

import "fmt"

func main() {
    var numbers []int = []int{6,5,3,4,1,2}
    fmt.Println("No bubblesort", numbers)
    bubblesort(numbers)
    fmt.Println("Sorted", numbers)
}

func bubblesort(numbers []int){
    for i:=0;i<len(numbers);i++{
	sweep(numbers,i)
    }

}

func sweep(numbers []int, step int){
    // assign n to to the lenght of the slice
    n := len(numbers)

    // index i and j start from 0,1
    i := 0
    j := 1

    // loop until j is lower than N
    for j < n-step {

	// assign numbers
	first := numbers[i]
	second := numbers[j]
	fmt.Println("considering: ", first,second)
	// change order of ith and jth element if not sorted
	if first > second {
	    fmt.Println("switching: ",first,second)
	    numbers[i] = second
	    numbers[j] = first
	}

	// increase indexes
	j++
	i++
    }

}
